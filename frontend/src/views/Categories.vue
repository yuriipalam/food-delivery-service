<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, onUnmounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getCategories} from "../api/api";
import {scrollToExploreBlock, setMainHeight} from "../utils";
import {ResponseError} from "../api/errors";
import router from "../router";

const categories = ref([])

// saving scroll position to session storage
function saveScrollPosition() {
  sessionStorage.setItem('scrollPosition', window.scrollY.toString())
}

onMounted(async () => {
  try {
    categories.value = await getCategories()
  } catch (err) {
    switch (err.message) {
      case ResponseError.notFound:
        await router.push({name: '404'})
        return
      default:
        await router.push({name: '500'})
        return
    }
  }

  await nextTick()

  setMainHeight()
  scrollToExploreBlock()

  // when user clicks we save scroll position that on the next page his scroll position
  // will remain the same by fetching it from session storage
  document.querySelectorAll('.flow-card').forEach(card => {
    card.addEventListener('click', saveScrollPosition)
  })
})

onUnmounted(() => {
  document.querySelectorAll('.flow-card').forEach(card => {
    card.removeEventListener('click', saveScrollPosition)
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
