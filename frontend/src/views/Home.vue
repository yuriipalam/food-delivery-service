<script setup>
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import Carousel from "../components/Carousel/Carousel.vue";
import {useFiltersStore} from "../store";
import {computed, nextTick, onMounted, ref} from "vue";
import Flow from "../components/Flow/Flow.vue";

const store = useFiltersStore()

const suppliers = ref([])
const categories = ref([])
const categoriesAndSuppliersFiltered = computed(() => {
  const categoriesAndSuppliers = suppliers.value.concat(categories.value)
  return categoriesAndSuppliers.filter((item) => item.name.toLowerCase().includes(store.searchFor.toLowerCase()))
})

const fetchSuppliers = async () => {
  return fetch('http://localhost:8080/suppliers').then((response) => response.json())
}

const fetchCategories = async () => {
  return fetch('http://localhost:8080/categories').then((response) => response.json())
}

onMounted(async () => {
  suppliers.value = await fetchSuppliers()

  categories.value = await fetchCategories()

  await nextTick()

  const main = document.querySelector('main')
  main.style.minHeight = main.offsetHeight + 'px'
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <Explore class="explore"></Explore>
    <main>
      <Carousel v-if="store.searchFor === ''" :name="'Top Suppliers'" :objects="suppliers" :url-path="'suppliers'"
                class="carousel-suppliers"></Carousel>
      <Carousel v-if="store.searchFor === ''" :name="'Top Categories'" :objects="categories" :url-path="'categories'"
                class="carousel-categories"></Carousel>
      <Flow v-if="store.searchFor !== ''" :items="categoriesAndSuppliersFiltered"
            :name="'Results for suppliers & categories'"></Flow>
    </main>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}

.carousel-suppliers {
  margin-bottom: 55px;
}

.carousel-categories {
  padding-bottom: 55px;
}
</style>
