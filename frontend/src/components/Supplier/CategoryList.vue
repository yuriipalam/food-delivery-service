<script setup>
import Category from "./Category.vue";
import {useProductsFiltersStore} from "../../store";

defineProps({
  categories: Array
})

const store = useProductsFiltersStore()

function selectCategory(id, name) {
  store.selectCategory(id, name)
}
</script>

<template>
  <div class="category-list">
    <p class="categories-title">Categories:</p>
    <div class="list">
      <Category @click="selectCategory(category.id, category.name)" :is-active="store.selectedCategory === category.id" v-for="category in categories" :key="category.id">{{ category.name }}</Category>
      <Category @click="selectCategory(0, 'All products')" :is-active="store.selectedCategory === 0">All products</Category>
    </div>
  </div>
</template>

<style scoped>
.category-list {
  display: inline-block;
}

.category-list > p {
  font-weight: 700;
  font-size: 20px;
  color: var(--blackish);
  opacity: 0.7;
  margin-bottom: 35px;
}

.categories-title {
  margin-top: 0;
}

</style>
