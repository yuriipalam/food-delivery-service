<script setup>
import FlowCard from "./FlowCard.vue";

const props = defineProps({
  name: String,
  items: Array
})
</script>

<template>
  <h3>{{ name }}</h3>
  <div class="items-flow">
    <FlowCard v-for="item in items" :obj="item" class="flow-card"></FlowCard>
  </div>
</template>

<style scoped>
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