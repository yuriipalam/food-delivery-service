<script setup>
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import Carousel from "../components/Carousel/Carousel.vue";
import {onMounted, ref} from "vue";

const suppliers = ref([])
const categories = ref([])

const fetchSuppliers = async () => {
  return fetch('http://localhost:8080/suppliers').then((response) => response.json())
}

const fetchCategories = async () => {
  return fetch('http://localhost:8080/categories').then((response) => response.json())
}

onMounted(async () => {
  suppliers.value = await fetchSuppliers()
  categories.value = await fetchCategories()
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Explore class="explore"></Explore>
      <Carousel :name="'Top Suppliers'" :objects="suppliers" :url-path="'suppliers'"
                class="carousel-suppliers"></Carousel>
      <Carousel :name="'Top Categories'" :objects="categories" :url-path="'categories'"
                class="carousel-categories"></Carousel>
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
