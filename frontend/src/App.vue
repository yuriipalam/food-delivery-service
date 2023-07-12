<script setup>
import ButtonComponent from "./components/PrimaryButton.vue";
import "vue-router";
import MenuLink from "./components/MenuLink.vue";
</script>

<template>
<!--  <button @click="count++">Add +1</button>-->
<!--  <p>{{ count }}</p>-->
  <router-link :to="{ name: 'Home' }">blabla</router-link>
  <MenuLink></MenuLink>
  <router-view></router-view>
<!--  <ButtonComponent :type="'btn'" :class="'header-button'">Damn</ButtonComponent>-->
<!--  <ButtonComponent :type="'btn'" :class="'solid-button'">Damn</ButtonComponent>-->
<!--  <ButtonComponent :type="'btn'" :class="'transparent-button'">Damn</ButtonComponent>-->
</template>
