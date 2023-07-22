<script setup>
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import Carousel from "../components/Carousel/Carousel.vue";
import {useFiltersStore} from "../store";
import {computed, nextTick, onMounted, ref} from "vue";
import Flow from "../components/Flow/Flow.vue";
import {getCategories, getSuppliers} from "../api/api";

const store = useFiltersStore()

const suppliers = ref([])
const categories = ref([])

// filtering suppliers & categories depending on searchbar
const categoriesAndSuppliersFiltered = computed(() => {
  const categoriesAndSuppliers = suppliers.value.concat(categories.value)
  return categoriesAndSuppliers.filter((item) => item.name.toLowerCase().includes(store.searchFor.toLowerCase()))
})


onMounted(async () => {
  suppliers.value = await getSuppliers()
  categories.value = await getCategories()

  await nextTick()

  const main = document.querySelector('main')
  main.style.minHeight = main.offsetHeight + 'px'
})
</script>

<template>
  <div class="container">
    <Header/>
    <Explore class="explore"/>
    <main>
      <Carousel v-if="store.searchFor === ''" :name="'Top Suppliers'" :objects="suppliers"
                class="carousel-suppliers"/>
      <Carousel v-if="store.searchFor === ''" :name="'Top Categories'" :objects="categories"
                class="carousel-categories"/>
      <Flow v-if="store.searchFor !== ''" :items="categoriesAndSuppliersFiltered"
            :name="'Results for suppliers & categories'"/>
    </main>
  </div>
</template>

<style scoped>
.carousel-suppliers {
  margin-bottom: 55px;
}

.carousel-categories {
  padding-bottom: 55px;
}
</style>
