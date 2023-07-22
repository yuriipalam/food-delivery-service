<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getCategories} from "../api/api";

const categories = ref([])

// saving scroll position to session storage
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

  // when user clicks we save scroll position that on the next page his scroll position
  // will remain the same by fetching it from session storage
  document.querySelectorAll('.flow-card').forEach(card => {
    card.addEventListener('click', saveScrollPosition)
  })
})
</script>

<template>
  <div class="container">
    <Header/>
    <main>
      <Explore class="explore"/>
      <Flow :items="categories" :name="'Categories'"/>
    </main>
  </div>
</template>

<style scoped>

</style>
