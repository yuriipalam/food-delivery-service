<script setup>
import {useRouter} from "vue-router";
import {onMounted, onUnmounted} from "vue";
import {hideFooter, showFooter} from "../utils";

const props = defineProps({
  text: String
})

const router = useRouter()

onMounted(() => {
  hideFooter()
})

onUnmounted(() => {
  showFooter()
})
</script>

<template>
  <div class="container">
    <h1>{{ props.text }}</h1>
    <router-link :to="{ name: 'Home'}">Go back Home</router-link>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin-top: 100px;
}

h1 {
  text-align: center;
  font-size: 75px;
  font-weight: 400;
}

a {
  text-align: center;
  font-size: 38px;
  text-decoration: none;
  font-weight: 600;
  color: var(--orange)
}
</style>
