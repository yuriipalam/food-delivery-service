<script setup>
import {useFiltersStore} from "../../store";
import CategoryHorizontal from "./CategoryHorizontal.vue";

defineProps({
  categories: Array
})

const useFilters = useFiltersStore()

function selectCategory(id, name) {
  useFilters.selectCategory(id, name)
}
</script>

<template>
  <div class="category-list-horizontal">
    <p class="categories-title-horizontal">Categories</p>
    <div class="list-horizontal">
      <CategoryHorizontal @click="selectCategory(0, 'All products')" :is-active="useFilters.selectedCategory === 0">
        All products
      </CategoryHorizontal>
      <CategoryHorizontal @click="selectCategory(category.category_id, category.category_name)"
                          :is-active="useFilters.selectedCategory === category.category_id" v-for="category in categories"
                          :key="category.category_id">{{ category.category_name }}
      </CategoryHorizontal>
    </div>
  </div>
</template>

<style scoped>
.category-list-horizontal {
  overflow-x: auto;
  display: flex;
  align-items: center;
}

.categories-title-horizontal {
  font-weight: 700;
  font-size: 20px;
  color: var(--blackish);
  opacity: 0.7;
  margin-right: 25px;
  margin-top: 0;
  grid-gap: 10px;
  margin-bottom: 0;
}

.list-horizontal {
  display: flex;
  grid-gap: 25px;
  white-space: nowrap;
  flex-grow: 1;
  overflow-x: auto;
  justify-content: space-between;
  padding-top: 15px;
  padding-bottom: 15px;
}

@media screen and (max-width: 480px) {
  .category-list-horizontal {
    overflow: initial;
    justify-content: center;
    flex-flow: wrap;
  }

  .categories-title-horizontal {
    margin-right: 0;
    margin-bottom: 15px;
  }

  .list-horizontal {
    overflow-x: auto;
  }
}
</style>
