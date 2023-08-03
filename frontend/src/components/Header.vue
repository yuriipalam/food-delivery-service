<script setup>
import {onMounted, onUnmounted} from "vue";
import PrimaryButton from "./PrimaryButton.vue";
import {useRouter} from "vue-router";

const router = useRouter()

let resizeTimout = 0

onMounted(() => {
  document.body.className = 'home';
  document.documentElement.className = 'home';
  if (!isSafariOnIphone()) {
    calculateMargins()
    window.addEventListener('resize', calculateMargins)
    resizeTimout = 300
  } else {
    const header = document.querySelector('header')
    header.style.marginTop = 80 + 'px'
    header.style.marginBottom = 120 + 'px'
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', calculateMargins)
  document.body.classList.remove('home')
  document.documentElement.classList.remove('home')
})

function isSafariOnIphone() {
  // Check for Safari and exclude Chrome and Android
  const safariTest = /^((?!chrome|android).)*safari/i;

  // Check if device is an iPhone
  const iPhoneTest = /iPhone/i;

  return safariTest.test(navigator.userAgent) && iPhoneTest.test(navigator.userAgent);
}

function calculateMargins() {
  const header = document.querySelector('header')
  const nav = document.querySelector('nav')
  const containerHeader = document.querySelector('.container header')
  const space = ((window.innerHeight - containerHeader.offsetHeight) / 2) - nav.offsetHeight

  setTimeout(() => {
    header.style.marginTop = space + 'px'
    header.style.marginBottom = space + nav.offsetHeight + 'px'
  }, resizeTimout)
}

function orderClick() {
  const scrollTo = document.querySelector('.explore').offsetTop - 40

  window.scrollTo({
    top: scrollTo,
    behavior: 'smooth'
  })
}
</script>

<template>
  <header>
    <h1>The best food delivery in Budapest</h1>
    <p>Order now with 10% discount</p>
    <div class="buttons">
      <PrimaryButton @click="orderClick()" :class="'solid-button'">Order now</PrimaryButton>
      <PrimaryButton @click="router.push({'name': 'SignUp'})" :class="'transparent-button'">Sign up</PrimaryButton>
    </div>
  </header>
</template>

<style>
body.home {
  height: 100vh;
  max-width: 100% !important;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: bottom;
  background-image: url('../assets/svg/home-background.svg');
}
</style>

<style scoped>
header {
  display: flex;
  flex-direction: column;
  max-width: 500px;
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

@media screen and (max-width: 780px) {
  header {
    max-width: 450px;
    margin-right: 30px;
  }

  p {
    margin-bottom: 40px;
  }

  .buttons {
    flex-direction: column-reverse;
    align-items: flex-start;
  }

  .transparent-button {
    margin-bottom: 15px;
  }
}

@media screen and (max-height: 500px) {
  header {
    max-width: 400px;
  }

  .buttons {
    flex-direction: row;
  }
}

@media screen and (max-height: 440px) {
  .transparent-button {
    background-color: var(--blackish);
  }
}
</style>
