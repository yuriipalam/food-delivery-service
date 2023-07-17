<script setup>
import Flow from "../components/Flow/Flow.vue";
import {onBeforeMount, onMounted, ref} from "vue";
import Header from "../components/Header.vue";

const categories = ref([])

const fetchCategories = async () => {
  return fetch(`http://localhost:8080/categories`).then((response) => response.json())
}

onMounted(async () => {
  categories.value = await fetchCategories()
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Flow :items="categories" :name="'Categories'"></Flow>
    </main>
  </div>
</template>

<style scoped>

</style>
