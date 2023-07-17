import { defineStore } from 'pinia'
import {ref} from "vue";

export const useProductsFiltersStore = defineStore("products", () => {
    const selectedCategory = ref(0)
    const selectedCategoryName = ref('')
    const searchFor = ref('')

    async function selectCategory(id, name) {
        selectedCategory.value = id
        selectedCategoryName.value = name
    }

    return { selectedCategory, searchFor, selectCategory, selectedCategoryName }
})
