<script setup>
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import Carousel from "../components/Carousel/Carousel.vue";
import {onBeforeUnmount, onMounted, ref} from "vue";

let arr = [
  {
    "id": 1,
    "categories": [
      {
        "category_id": 1,
        "category_name": "Fastfood"
      },
      {
        "category_id": 2,
        "category_name": "Category 1"
      }
    ],
    "name": "ABC Electronics",
    "image_url": "http://localhost:8080/images/suppliers/mcdonalds.png",
    "description": "ABC Electronics supplier description",
    "time_opening": "09:00",
    "time_closing": "18:00"
  },
  {
    "id": 2,
    "categories": [
      {
        "category_id": 2,
        "category_name": "Category 1"
      }
    ],
    "name": "XYZ Clothing",
    "image_url": "http://localhost:8080/images/suppliers/mcdonalds.png",
    "description": "XYZ Clothing supplier description",
    "time_opening": "10:00",
    "time_closing": "20:00"
  },
  {
    "id": 3,
    "categories": [
      {
        "category_id": 3,
        "category_name": "Category 2"
      }
    ],
    "name": "Bookworms Bookstore",
    "image_url": "http://localhost:8080/images/suppliers/mcdonalds.png",
    "description": "Bookworms Bookstore supplier description",
    "time_opening": "08:30",
    "time_closing": "17:30"
  }
]

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
  categories.value = await  fetchCategories()
})

</script>

<template>
    <div class="container">
      <Header></Header>
      <main>
        <Explore class="explore"></Explore>
        <Carousel :name="'Top Suppliers'" :objects="suppliers" class="carousel-suppliers"></Carousel>
        <Carousel :name="'Top Categories'" :objects="categories" class="carousel-categories"></Carousel>
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
