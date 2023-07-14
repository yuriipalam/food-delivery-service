<script setup>
import Header from "./Header.vue";
import Explore from "./Explore.vue";
import Carousel from "../../components/Carousel.vue";
</script>

<template>
    <div class="container">
      <Header></Header>
      <main>
        <Explore class="explore"></Explore>
        <Carousel :name="'Suppliers'" class="carousel-suppliers"></Carousel>
        <Carousel :name="'Categories'" class="carousel-categories"></Carousel>
      </main>
    </div>
</template>

<style scoped>
.explore {
  margin-bottom: 55px;
}

.carousel-suppliers {
  margin-bottom: 20px;
}

.carousel-categories {
  margin-bottom: 20px;
}
</style>
