<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, onUnmounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getCategories} from "../api/api";
import {useFiltersStore} from "../store";

const store = useFiltersStore()

const categories = ref([])

function saveScrollPosition() {
  sessionStorage.setItem('scrollPosition', window.scrollY.toString())
}

onMounted(async () => {
  categories.value = await getCategories()

  await nextTick()

  const main = document.querySelector('main')
  main.style.minHeight = main.offsetHeight + 'px'

  const scrollTo = document.querySelector('.explore').offsetTop - 40

  window.scrollTo({
    top: scrollTo,
    behavior: 'smooth'
  })

  document.querySelectorAll('.flow-card').forEach(card => {
    card.addEventListener('click', saveScrollPosition)
  })
})

onUnmounted(() => {
  store.reset()
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Explore class="explore"></Explore>
      <Flow :items="categories" :name="'Categories'"></Flow>
    </main>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}
</style>
