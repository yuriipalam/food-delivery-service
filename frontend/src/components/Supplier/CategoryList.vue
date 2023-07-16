<script setup>
import Category from "./Category.vue";

const categories = [
  'McMenu',
  'Desserts',
  'Drinks',
  'McCafe'
]
</script>

<template>
  <div class="category-list">
    <p class="categories-title">Categories:</p>
    <div class="list" v-for="category in categories">
      <hr>
      <Category>{{ category }}</Category>
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

.category-list hr {
  height: 1px;
  border: 0;
  background-color: rgba(39, 45, 47, 0.3);
}
</style>
