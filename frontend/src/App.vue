<script setup>
import "@/assets/css/main.css"
import "vue-router";
import Navbar from "./components/Navigation/Navbar.vue";
import {useRoute} from "vue-router";

const route = useRoute()

function isHomeRoute() {
  return route.name === 'Home'
}
</script>

<template>
  <Navbar :is-light="isHomeRoute()"></Navbar>
  <router-view></router-view>
</template>
