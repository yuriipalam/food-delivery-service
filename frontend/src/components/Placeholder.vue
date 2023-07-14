<script setup>
// If you are using PurgeCSS, make sure to whitelist the carousel CSS classes
import 'vue3-carousel/dist/carousel.css'
import { Carousel, Slide, Pagination, Navigation } from 'vue3-carousel'

const items = [
  'slide 1',
  'slide 2',
  'slide 3',
  'slide 4',
  'slide 5',
]
</script>

<template>
  <carousel :items-to-show="5" :autoplay="5000">
    <slide v-for="slide in items" :key="slide">
      {{ slide }}
    </slide>

    <template #addons>
      <navigation />
      <pagination />
    </template>
  </carousel>
</template>

<style scoped>
ol {
  padding-left: 0;
}
</style>