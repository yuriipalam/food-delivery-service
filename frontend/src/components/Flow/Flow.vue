<script setup>
import Explore from "../Explore.vue";
import {computed, nextTick, onBeforeMount, onBeforeUpdate, onMounted, onUnmounted, watch} from "vue";
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

// onBeforeUpdate(() => {
//   const neededHeight =  window.getComputedStyle(document.querySelector('.flow-card')).getPropertyValue('height')
//   const allFlowCards = document.querySelectorAll('.flow-card')
//   allFlowCards.forEach(flowCard => {
//     flowCard.style.height = neededHeight + 'px'
//   })
// })
</script>

<template>
  <div class="items-list-container">
    <h3>{{ props.name }}</h3>
    <div class="items-flow">
      <FlowCard v-for="item in filteredItems" :obj="item" class="flow-card"></FlowCard>
      <span v-if="filteredItems.length === 0" class="not-found">Not found</span>
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

.not-found {
  font-size: 28px;
  margin-top: 25px;
}

.flow-card {
  flex: 0 1 calc(20% - 50px);
}
</style>
