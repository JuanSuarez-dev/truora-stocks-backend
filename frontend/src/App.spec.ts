import { mount, flushPromises } from '@vue/test-utils'
import App from './App.vue'
import { createPinia, setActivePinia } from 'pinia'
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'

describe('App.vue', () => {
  let mock: MockAdapter

  beforeEach(() => {
    setActivePinia(createPinia())
    mock = new MockAdapter(axios)
  })
  afterEach(() => {
    mock.restore()
  })

  it('renders table and best pick', async () => {
    // mock the stocks list
    mock.onGet(/\/api\/stocks$/).reply(200, [
      {
        ticker: 'X',
        company: 'X Co',
        brokerage: 'Y',
        action: 'up',
        rating_from: 'A',
        rating_to: 'B',
        target_from: '$10',
        target_to: '$15',
        time: '2025-01-01T00:00:00Z'
      }
    ])
    // mock the best pick
    mock.onGet(/\/api\/stocks\/best$/).reply(200, { ticker: 'X', upside: '50.0%' })

    const wrapper = mount(App, {
      global: {
        plugins: [createPinia()]
      }
    })

    // wait for both fetchAll() and fetchBest() to resolve
    await flushPromises()

    // check best pick rendered
    expect(wrapper.text()).toContain('Best pick: X (50.0%)')
    // check table row contains ticker X
    const firstCell = wrapper.find('tbody tr td')
    expect(firstCell.exists()).toBe(true)
    expect(firstCell.text()).toBe('X')
  })
})
