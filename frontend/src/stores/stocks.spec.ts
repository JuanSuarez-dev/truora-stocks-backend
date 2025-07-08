import { setActivePinia, createPinia } from 'pinia'
import { useStockStore } from './stocks'
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'

describe('Stock Store', () => {
  let mock: MockAdapter

  beforeEach(() => {
    setActivePinia(createPinia())
    mock = new MockAdapter(axios)
  })
  afterEach(() => {
    mock.restore()
  })

  it('fetchAll() fills list with API data', async () => {
    const fakeItems = [
      {
        ticker: 'A',
        company: 'A Co',
        brokerage: 'X',
        action: 'up',
        rating_from: 'R1',
        rating_to: 'R2',
        target_from: '$1',
        target_to: '$2',
        time: '2025-01-01T00:00:00Z'
      }
    ]
    // intercept any GET ending in /api/stocks
    mock.onGet(/\/api\/stocks$/).reply(200, fakeItems)

    const store = useStockStore()
    await store.fetchAll()

    expect(store.list).toHaveLength(1)
    expect(store.list[0].ticker).toBe('A')
  })

  it('fetchBest() sets recommended correctly', async () => {
    const fakePick = { ticker: 'B', upside: '50.0%' }
    // intercept any GET ending in /api/stocks/best
    mock.onGet(/\/api\/stocks\/best$/).reply(200, fakePick)

    const store = useStockStore()
    await store.fetchBest()

    expect(store.recommended.ticker).toBe('B')
    expect(store.recommended.upside).toBe('50.0%')
  })
})
