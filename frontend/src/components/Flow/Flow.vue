<script setup>
import Explore from "../Explore.vue";
import ItemsFlow from "../Flow/ItemsFlow.vue";
import {computed, onMounted, ref} from "vue";
import {useFiltersStore} from "../../store";

const props = defineProps({
  items: Array,
  name: String
})

const store = useFiltersStore()

const filteredItems = computed(() => {
  if (store.searchFor === '') {
    return props.items
  }
  return props.items.filter((item) => item.name.toLowerCase().includes(store.searchFor.toLowerCase()))
})
</script>

<template>
  <Explore class="explore"></Explore>
  <div class="items-list-container">
    <ItemsFlow :items="filteredItems" :name="props.name"></ItemsFlow>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}

.items-list-container {
  max-width: 1100px;
  margin: 0 auto;
}
</style>
