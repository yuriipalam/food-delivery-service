<script setup>
import CarouselCard from "./CarouselCard.vue";
import {Carousel, Navigation, Pagination, Slide} from "vue3-carousel";

const props = defineProps({
  name: String
})

</script>

<template>
  <h3>{{ props.name }}</h3>
  <div style="max-width: 85%; margin: 0 auto;">
    <carousel :items-to-show="5">
      <slide v-for="slide in 10" :key="slide">
        <CarouselCard></CarouselCard>
      </slide>

      <template #addons>
        <navigation/>
        <pagination/>
      </template>
    </carousel>
  </div>

  <!--  <div class="carousel">-->
  <!--    <h3>{{ props.name }}</h3>-->
  <!--    <div class="carousel-items">-->
  <!--      <svg class="left-arrow" xmlns="http://www.w3.org/2000/svg" width="18" height="33" viewBox="0 0 18 33" fill="none">-->
  <!--        <g opacity="0.7">-->
  <!--          <path-->
  <!--              d="M16.7278 4.37967C17.4476 3.5377 17.3864 2.23742 16.5911 1.4754C15.7958 0.713382 14.5676 0.778196 13.8478 1.62017L16.7278 4.37967ZM13.8478 1.62017L0.27218 17.5005L3.15219 20.26L16.7278 4.37967L13.8478 1.62017Z"-->
  <!--              fill="#272D2F"/>-->
  <!--          <line x1="2" y1="-2" x2="17.1522" y2="-2" transform="matrix(0.711991 0.702189 -0.660592 0.750745 2 19)"-->
  <!--                stroke="#272D2F" stroke-width="4" stroke-linecap="round"/>-->
  <!--        </g>-->
  <!--      </svg>-->
  <!--      <div class="cards">-->
  <!--        <CarouselCard></CarouselCard>-->
  <!--        <CarouselCard></CarouselCard>-->
  <!--        <CarouselCard></CarouselCard>-->
  <!--        <CarouselCard></CarouselCard>-->
  <!--        <CarouselCard></CarouselCard>-->
  <!--      </div>-->
  <!--      <svg class="right-arrow" xmlns="http://www.w3.org/2000/svg" width="18" height="33" viewBox="0 0 18 33"-->
  <!--           fill="none">-->
  <!--        <g opacity="0.7">-->
  <!--          <path-->
  <!--              d="M1.27218 4.37967C0.552409 3.5377 0.613629 2.23742 1.40892 1.4754C2.20422 0.713382 3.43242 0.778196 4.15219 1.62017L1.27218 4.37967ZM4.15219 1.62017L17.7278 17.5005L14.8478 20.26L1.27218 4.37967L4.15219 1.62017Z"-->
  <!--              fill="#272D2F"/>-->
  <!--          <line x1="2" y1="-2" x2="17.1522" y2="-2" transform="matrix(-0.711991 0.702189 0.660592 0.750745 16 19)"-->
  <!--                stroke="#272D2F" stroke-width="4" stroke-linecap="round"/>-->
  <!--        </g>-->
  <!--      </svg>-->
  <!--    </div>-->
  <!--    <ul class="carousel-points">-->
  <!--      <li class="point active"></li>-->
  <!--      <li class="point"></li>-->
  <!--      <li class="point"></li>-->
  <!--    </ul>-->
  <!--  </div>-->
</template>
<style>
button.carousel__prev {
  margin-left: -50px;
}

button.carousel__next {
  margin-right: -50px;
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

/*
.carousel-items {
  display: flex;
  flex-grow: 1;
  align-items: center;
}

.cards {
  display: flex;
  justify-content: space-between;
  flex-grow: 1;
}

.left-arrow {
  padding-right: 55px;
}

.right-arrow {
  padding-left: 55px;
}

.carousel-points {
  display: flex;
  justify-content: center;
}

ul {
  padding: 0;
}
li {
  list-style: none;
  margin-right: 15px;
}
li:last-child {
  margin-right: 0;
}

.point {
  border-radius: 50%;
  height: 10px;
  width: 10px;
  opacity: 0.5;
  background-color: var(--blackish);
}

.point.active {
  opacity: 1;
}*/
</style>