<template>
  <div class="p-6">
    <h1 class="text-3xl mb-4">Stock Recommendations</h1>

    <div v-if="stocks.recommended.ticker" class="mb-4">
      <span class="font-semibold">Best pick:</span>
      {{ stocks.recommended.ticker }} ({{ stocks.recommended.upside }})
    </div>

    <button
      @click="loadAll"
      class="px-4 py-2 bg-blue-500 text-white rounded mb-4"
      :disabled="stocks.loading"
    >
      {{ stocks.loading ? 'Loading…' : 'Load Stocks' }}
    </button>

    <table class="w-full table-auto border">
      <thead class="bg-gray-100">
        <tr>
          <th class="border px-2">Ticker</th>
          <th class="border px-2">Company</th>
          <th class="border px-2">Rating</th>
          <th class="border px-2">Target</th>
          <th class="border px-2">Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="s in stocks.list" :key="s.ticker">
          <td class="border px-2">{{ s.ticker }}</td>
          <td class="border px-2">{{ s.company }}</td>
          <td class="border px-2">{{ s.rating_from }} → {{ s.rating_to }}</td>
          <td class="border px-2">{{ s.target_from }} → {{ s.target_to }}</td>
          <td class="border px-2">
          {{
            // strip after 3 fractional digits, then parse
            new Date(s.time.replace(/\.\d{3}\d+Z$/, (m) => m.slice(0, 4) + "Z"))
              .toLocaleString()
          }}
        </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useStockStore } from './stores/stocks'

const stocks = useStockStore()

async function loadAll() {
  await stocks.fetchAll()
  await stocks.fetchBest()
}

onMounted(loadAll)
</script>
