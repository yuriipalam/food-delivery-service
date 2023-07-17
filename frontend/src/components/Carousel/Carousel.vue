<script setup>
import CarouselCard from "./CarouselCard.vue";
import {Carousel, Navigation, Pagination, Slide} from "vue3-carousel";
import 'vue3-carousel/dist/carousel.css'

const props = defineProps({
  name: String,
  objects: Array,
  urlPath: String,
})
</script>

<template>
  <div>
    <h3>{{ props.name }}</h3>
    <div style="max-width: 85%; margin: 0 auto;">
      <carousel :items-to-show="5">
        <slide v-for="obj in objects" :key="obj">
          <CarouselCard :obj="obj" :url-path="props.urlPath"></CarouselCard>
        </slide>
        <template #addons>
          <navigation/>
          <pagination/>
        </template>
      </carousel>
    </div>
  </div>
</template>

<style>
button.carousel__prev {
  margin-left: -50px;
}

button.carousel__next {
  margin-right: -50px;
}

.carousel__pagination-button:hover::after, .carousel__pagination-button--active::after {
  opacity: 0.5;
}

.carousel__pagination-button::after {
  height: 10px!important;
  width: 10px!important;
  border-radius: 50%!important;
  opacity: 0.5;
}
</style>

<style scoped>
ol {
  padding-left: 0;
}

h3 {
  text-align: center;
  padding-bottom: 10px;
  color: var(--blackish);
  font-size: 24px;
  font-weight: 400;
}
</style>
