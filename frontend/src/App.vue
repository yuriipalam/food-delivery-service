<script setup>
import "@/assets/css/main.css"
import "vue-router";

import MenuLink from "./components/MenuLink.vue";
import PrimaryButton from "./components/PrimaryButton.vue";
</script>

<template>
  <MenuLink :route-name="'Home'">Home</MenuLink>
  <MenuLink :route-name="'Suppliers'">Suppliers</MenuLink>
  <main class="container">
    <router-view></router-view>
  </main>
</template>
