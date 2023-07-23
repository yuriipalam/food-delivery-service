<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getSuppliers} from "../api/api";
import {scrollToExploreBlock, setMainHeight} from "../utils";
import {ResponseError} from "../api/errors";
import router from "../router";

const suppliers = ref([])

onMounted(async () => {
  try {
    suppliers.value = await getSuppliers()
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
})
</script>

<template>
  <div class="container">
    <Header/>
    <main>
      <Explore class="explore"/>
      <Flow :items="suppliers" :name="'Suppliers'"/>
    </main>
  </div>
</template>

<style scoped>

</style>
