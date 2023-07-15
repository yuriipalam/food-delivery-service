<script setup>
import SearchBar from "./SearchBar.vue";
</script>

<template>
  <div class="explore-block">
    <h2>Explore<br>categories & suppliers</h2>
    <SearchBar :class="'pinky'"></SearchBar>
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
</style>
