<script setup>
import Flow from "../components/Flow/Flow.vue";
import {nextTick, onMounted, onUnmounted, ref} from "vue";
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import {getSuppliersByCategoryID} from "../api/api";
import {useFiltersStore} from "../store";
import {useRoute} from "vue-router";

const route = useRoute()
const id = route.params.id

const store = useFiltersStore()

const suppliers = ref([])
const categoryName = ref('')

onMounted(async () => {
  suppliers.value = await getSuppliersByCategoryID(id)
  categoryName.value = suppliers.value[0].categories.find(category => category.category_id === parseInt(id)).category_name

  await nextTick()

  if (sessionStorage.getItem('scrollPosition')) {
  window.scrollTo(0, parseInt(sessionStorage.getItem('scrollPosition')))
  }
})

onUnmounted(() => {
  store.reset()
  sessionStorage.removeItem('scrollPosition')
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Explore class="explore"></Explore>
      <Flow :items="suppliers" :name="'Suppliers in ' + categoryName"></Flow>
    </main>
  </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}
</style>
