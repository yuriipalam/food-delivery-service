<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getSuppliers} from "../api/api";

const suppliers = ref([])

onMounted(async () => {
  suppliers.value = await getSuppliers()

  await nextTick()

  const main = document.querySelector('main')
  main.style.minHeight = main.offsetHeight + 'px'

  const scrollTo = document.querySelector('.explore').offsetTop - 40

  window.scrollTo({
    top: scrollTo,
    behavior: 'smooth'
  })
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
