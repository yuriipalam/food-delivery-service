<script setup>
import "@/assets/css/main.css"
import "vue-router";
import Navbar from "./components/Navbar.vue";
import {useRoute} from "vue-router";

const route = useRoute()

function isHomeRoute() {
  return route.name === 'Home'
}
</script>

<template>
  <Navbar :is-light="isHomeRoute()"></Navbar>
  <div class="container" id="body-container">
    <router-view></router-view>
  </div>
</template>
