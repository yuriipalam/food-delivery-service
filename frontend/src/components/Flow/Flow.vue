<script setup>
import Explore from "../Explore.vue";
import {computed, onUnmounted} from "vue";
import {useFiltersStore} from "../../store";
import FlowCard from "./FlowCard.vue";

const props = defineProps({
  items: Array,
  name: String,
})

const store = useFiltersStore()

const filteredItems = computed(() => {
  if (store.searchFor === '') {
    return props.items
  }
  return props.items.filter((item) => item.name.toLowerCase().includes(store.searchFor.toLowerCase()))
})

onUnmounted(async () => {
  await store.reset()
})
</script>

<template>
  <div class="items-list-container">
    <h3>{{ props.name }}</h3>
    <div class="items-flow">
      <FlowCard v-for="item in filteredItems" :obj="item" class="flow-card"></FlowCard>
    </div>
  </div>
</template>

<style scoped>
.items-list-container {
  max-width: 1100px;
  margin: 0 auto;
}

h3 {
  text-align: center;
  padding-bottom: 10px;
  color: var(--blackish);
  font-size: 24px;
  font-weight: 400;
}

.items-flow {
  display: flex;
  flex-flow: row wrap;
  grid-gap: 60px;
  align-self: center;
  justify-content: center;
}

.flow-card {
  flex: 0 1 calc(20% - 50px);
}
</style>
