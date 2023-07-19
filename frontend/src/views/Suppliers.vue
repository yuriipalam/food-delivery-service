<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onBeforeMount, onMounted, onUnmounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";

const suppliers = ref([])

const fetchSuppliers = async () => {
  return fetch(`http://localhost:8080/suppliers`).then((response) => response.json())
}

onMounted(async () => {
  suppliers.value = await fetchSuppliers()

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
      <Flow :items="suppliers" :name="'Suppliers'"></Flow>
    </main>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}
</style>
