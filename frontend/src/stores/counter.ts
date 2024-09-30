import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useTableStore
    = defineStore('table', () => {
      const country = ref<string>('')

      return { country }
    })
