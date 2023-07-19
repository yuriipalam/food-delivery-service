<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";

const categories = ref([])

const fetchCategories = async () => {
  return fetch(`http://localhost:8080/categories`).then((response) => response.json())
}

onMounted(async () => {
  categories.value = await fetchCategories()

  await nextTick()

  const main = document.querySelector('main')
  main.style.minHeight = main.offsetHeight + 'px'
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Explore class="explore"></Explore>
      <Flow :items="categories" :name="'Categories'"></Flow>
    </main>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}
</style>
