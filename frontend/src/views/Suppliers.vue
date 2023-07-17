<script setup>
import Flow from "../components/Flow/Flow.vue";
import {onBeforeMount, onMounted, ref} from "vue";
import Header from "../components/Header.vue";

const suppliers = ref([])

const fetchSuppliers = async () => {
  return fetch(`http://localhost:8080/suppliers`).then((response) => response.json())
}

onMounted(async () => {
  suppliers.value = await fetchSuppliers()
})
</script>

<template>
  <div class="container">
    <Header></Header>
    <main>
      <Flow :items="suppliers" :name="'Suppliers'"></Flow>
    </main>
  </div>
</template>

<style scoped>

</style>
