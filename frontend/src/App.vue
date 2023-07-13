<script setup>
import "@/assets/css/main.css"
import "vue-router";
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <Navbar></Navbar>
  <main class="container">
    <router-view></router-view>
  </main>
</template>
