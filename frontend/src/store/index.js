import { defineStore } from 'pinia'
import {ref} from "vue";

export const useFiltersStore = defineStore("productsFilter", () => {
    const selectedCategory = ref(0)
    const selectedCategoryName = ref('')
    const searchFor = ref('')

    async function selectCategory(id, name) {
        selectedCategory.value = id
        selectedCategoryName.value = name
    }

    async function reset() {
        selectedCategory.value = 0
        selectedCategoryName.value = ''
        searchFor.value = ''
    }

    return { selectedCategory, searchFor, selectCategory, selectedCategoryName, reset }
})
