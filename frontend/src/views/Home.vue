<script setup>
import Header from "../components/Header.vue";
import Explore from "../components/Explore.vue";
import Carousel from "../components/Carousel/Carousel.vue";
</script>

<template>
    <div class="container">
      <Header></Header>
      <main>
        <Explore class="explore"></Explore>
        <Carousel :name="'Top Suppliers'" class="carousel-suppliers"></Carousel>
        <Carousel :name="'Top Categories'" class="carousel-categories"></Carousel>
      </main>
    </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}

.carousel-suppliers {
  margin-bottom: 55px;
}

.carousel-categories {
  padding-bottom: 55px;
}
</style>
