<script setup>
import "@/assets/css/main.css"
import "vue-router";
import Navbar from "./components/Navigation/Navbar.vue";
import {useRoute} from "vue-router";

const route = useRoute()

function isLight() {
  return route.name === 'Home' || route.name === 'Suppliers' || route.name === 'Categories' || route.name === 'Category'
}
</script>

<template>
  <Navbar :is-light="isLight()"></Navbar>
  <router-view></router-view>
</template>
