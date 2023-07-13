<script setup>
const props = defineProps({
  routeName: String,
  isLight: Boolean,
})
</script>

<template>
<router-link :to="{ name: props.routeName}" :class="{ light: isLight }"><slot></slot></router-link>
</template>

<style scoped>
  a {
    color: var(--blackish);
    font-size: 16px;
    font-weight: 400;
    text-decoration: none;
    transition: opacity 0.7s ease;
  }
  a:hover {
    opacity: 85%;
  }

  .light {
    color: white!important;;
  }

  .router-link-active {
    pointer-events: none;
    font-weight: 500;
  }
  .router-link-active:hover {
    opacity: 100%;
  }
</style>