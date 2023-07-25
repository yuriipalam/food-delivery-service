<script setup>
import SearchBar from "./SearchBar.vue";
</script>

<template>
  <div class="explore-block">
    <h2>Explore<br>categories & suppliers</h2>
    <SearchBar class="search-bar" :class="'pinky'"/>
  </div>
</template>

<style scoped>
.explore-block {
  display: flex;
  flex-direction: column;
  padding-left: 60px;
  padding-top: 110px;
  padding-bottom: 60px;
  background-color: var(--pinky);
  border-radius: 40px;
  background-image: url("../assets/img/explore-background.png");
  background-position: right bottom;
  background-size: contain;
  background-repeat: no-repeat;
}

h2 {
  margin-top: 0;
  font-size: 36px;
  font-weight: 400;
}

.search-bar {
  max-width: 600px;
}

@media screen and (max-width: 1275px) {
  .search-bar {
    max-width: 460px;
  }
}

@media screen and (max-width: 1160px) {
  .explore-block {
    background-image: url("../assets/img/explore-background-small-devices.png");
    padding-top: 95px;
  }

  .search-bar {
    max-width: 600px;
  }
}

@media screen and (max-width: 1028px) {
  .search-bar {
    max-width: 400px;
  }
}

@media screen and (max-width: 875px) {
  .explore-block {
    background-position-x: 200px;
  }
}

@media screen and (max-width: 780px) {
  .explore-block {
    background-position-x: 150px;
    padding-left: 35px;
  }
}

@media screen and (max-width: 660px) {
  .explore-block {
    background-position-x: 200px;
    background-image: none;
    padding-right: 30px;
    padding-left: 30px;
    padding-top: 70px;
  }

  .search-bar {
    max-width: 100%;
  }
}
</style>
