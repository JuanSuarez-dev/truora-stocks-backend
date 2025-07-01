import { defineStore } from 'pinia'
import axios from 'axios'

// Modelo de un ítem de stock
export interface StockItem {
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: string
  target_to: string
  time: string
}

// Modelo de la recomendación
export interface BestPick {
  ticker: string
  upside: string
}

export const useStockStore = defineStore('stocks', {
  state: () => ({
    list: [] as StockItem[],
    recommended: {} as BestPick,
    loading: false,
  }),
  actions: {
    // Trae todos los stocks
    async fetchAll() {
      this.loading = true
      try {
        const res = await axios.get<StockItem[]>('/api/stocks')
        this.list = res.data
      } finally {
        this.loading = false
      }
    },
    // Trae la mejor recomendación
    async fetchBest() {
      const res = await axios.get<BestPick>('/api/stocks/best')
      this.recommended = res.data
    },
  },
})
