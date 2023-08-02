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
  margin-top: 2vh;
  margin-bottom: 10vh;
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

@media screen and (max-width: 600px) {
  h1 {
    font-size: 62px;
  }

  a {
    font-size: 36px;
  }
}

@media screen and (max-width: 480px) {
  h1 {
    font-size: 58px;
  }

  a {
    font-size: 34px;
  }
}
</style>
