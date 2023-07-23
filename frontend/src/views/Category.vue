<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, onUnmounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getSuppliersByCategoryID} from "../api/api";
import {useRoute} from "vue-router";
import {ResponseError} from "../api/errors";
import router from "../router";

const route = useRoute()
const id = route.params.id

const suppliers = ref([])
const categoryName = ref('')

onMounted(async () => {
  try {
    suppliers.value = await getSuppliersByCategoryID(id)
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

  categoryName.value = suppliers.value[0].categories.find(category => category.category_id === parseInt(id)).category_name

  await nextTick()

  // fetching scroll position from sessionStorage which was set on Categories view
  if (sessionStorage.getItem('scrollPosition')) {
    window.scrollTo(0, parseInt(sessionStorage.getItem('scrollPosition')))
  }
})

onUnmounted(() => {
  sessionStorage.removeItem('scrollPosition')
})
</script>

<template>
  <div class="container">
    <Header/>
    <main>
      <Explore class="explore"/>
      <Flow :items="suppliers" :name="'Suppliers in ' + categoryName"/>
    </main>
  </div>
</template>

<style scoped>

</style>
