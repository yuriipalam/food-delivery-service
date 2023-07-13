<script setup>
import {onBeforeMount, onUnmounted} from "vue";
import PrimaryButton from "../components/PrimaryButton.vue";

onBeforeMount(() => {
  document.body.className = 'home';
  document.documentElement.className = 'home';
  document.querySelector('#body-container').classList.add('home')
  document.querySelector('#app').classList.add('home')
})

onUnmounted(() => {
  document.body.classList.remove('home')
  document.documentElement.classList.remove('home')
  document.querySelector('#body-container').classList.remove('home')
  document.querySelector('#app').classList.remove('home')
})
</script>

<template>
  <header>
    <h1>The best food delivery in Budapest</h1>
    <p>Order now with 10% discount</p>
    <div class="buttons">
      <PrimaryButton :class="'solid-button'">Order now</PrimaryButton>
      <PrimaryButton :class="'transparent-button'">Sign up</PrimaryButton>
    </div>
  </header>
</template>

<style>
html.home {
  height: 100%;
}

#app.home {
  height: inherit;
}
#body-container.home {
  height: 90%;
}

body.home {
  max-width: 100%!important;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: bottom;
  background-image: url('../assets/svg/home-background.svg');
}
</style>

<style scoped>
header {
  height: inherit;
  display: flex;
  flex-direction: column;
  max-width: 500px;
  justify-content: center;
}

h1 {
  color: white;
  font-size: 40px;
  font-weight: 400;
  text-transform: uppercase;
  margin-top: 0;
  margin-bottom: 20px;
}

p {
  color: white;
  font-size: 24px;
  font-weight: 400;
  margin-top: 0;
  margin-bottom: 70px;
}

.buttons {
  display: flex;
}

.solid-button {
  margin-right: 40px;
}
</style>
