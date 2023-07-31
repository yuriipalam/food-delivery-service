<script setup>


import {truncate} from "../../utils";

const props = defineProps({
  name: String,
  desc: String,
  quantity: String,
  imageURL: String
})
</script>

<template>
  <div class="bar">
    <img class="bar-img" :src="props.imageURL" :alt="props.name">
    <div class="bar-content">
      <h1>{{ props.name }}</h1>
      <p class="desc" v-if="props.desc">{{ truncate(props.desc, 100) }}</p>
      <p class="quantity">{{ props.quantity }} available products</p>
    </div>
  </div>
</template>

<style scoped>
.bar {
  display: flex;
  align-items: center;
  grid-gap: 25px;
  border-radius: 30px;
  background: rgba(243, 207, 121, 0.3);
  box-shadow: 4px 4px 7px 5px rgba(0, 0, 0, 0.1);
  padding-top: 40px;
  padding-bottom: 40px;
  padding-left: 65px;
}

.bar-img {
  padding-top: 5px;
  height: fit-content;
}

.bar-content {
  text-align: left;
  display: flex;
  flex-direction: column;
}

h1 {
  font-size: 48px;
  font-weight: 500;
  color: var(--blackish);
  margin: 0 0;
}

.desc {
  color: black;
  font-weight: 500;
  font-size: 16px;
  margin-top: 5px;
  margin-bottom: 10px;
  max-width: 500px;
}

.quantity {
  margin: 0 0;
  font-size: 14px;
  font-weight: 400;
}

@media screen and (max-width: 780px) {
  .bar {
    justify-content: center;
    flex-flow: wrap;
    grid-gap: 30px;
    padding-left: 30px;
    padding-right: 30px;
  }

  .bar-content h1, p {
    text-align: center;
  }
}

@media screen and (max-width: 480px) {
  .bar {
    padding-left: 15px;
    padding-right: 15px;
  }

  .bar-content h1 {
    word-break: break-all;
  }
}
</style>
