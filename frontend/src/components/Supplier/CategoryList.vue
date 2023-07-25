<script setup>
import Category from "./Category.vue";
import {useFiltersStore} from "../../store";

defineProps({
  categories: Array
})

const useFilters = useFiltersStore()

function selectCategory(id, name) {
  useFilters.selectCategory(id, name)
}
</script>

<template>
  <div class="category-list">
    <p class="categories-title">Categories:</p>
    <div class="list">
      <Category @click="selectCategory(category.id, category.name)" :is-active="useFilters.selectedCategory === category.id" v-for="category in categories" :key="category.id">{{ category.name }}</Category>
      <Category @click="selectCategory(0, 'All products')" :is-active="useFilters.selectedCategory === 0">All products</Category>
    </div>
  </div>
</template>

<style scoped>
.category-list {
  display: inline-block;
}

.categories-title {
  font-weight: 700;
  font-size: 20px;
  color: var(--blackish);
  opacity: 0.7;
  margin-bottom: 35px;
  margin-top: 0;
}
</style>
