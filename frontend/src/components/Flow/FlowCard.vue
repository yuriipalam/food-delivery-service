<script setup>
import {truncate} from "../../utils";

const props = defineProps({
  obj: Object
})
</script>

<template>
  <a class="flow-card" :href="props.obj.url">
    <div class="circle-mask">
      <img :src="props.obj.image_url" :alt="props.obj.name" width="100" height="100">
    </div>
    <p class="card-name">{{ props.obj.name }}</p>
    <p class="card-working-hours" v-if="typeof props.obj.time_opening !== 'undefined'">{{ props.obj.time_opening }} -
      {{ props.obj.time_closing }}</p>
    <p class="card-desc">{{ truncate(props.obj.description, 50) }}</p>
  </a>
</template>

<style scoped>
.flow-card {
  text-decoration: none;
  color: black;
  min-height: 100%;
  display: flex;
  align-items: center;
  width: 120px;
  flex-direction: column;
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 20px;
  background-color: rgba(243, 207, 121, 0.4);
  border-radius: 50px 50px 2px 2px;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  transition: opacity 0.7s ease;
}

.flow-card:hover {
  opacity: 0.75;
}

.circle-mask {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  overflow: hidden;
  margin-bottom: 15px;
  margin-top: -15px;
}

.circle-mask img {
  object-fit: cover;
}

.flow-card > p {
  text-align: center;
}

.card-name {
  margin: 0;
  padding-bottom: 8px;
}

.card-working-hours {
  opacity: 0.4;
  font-weight: 300;
  font-size: 12px;
  margin-top: 0;
}

.card-desc {
  margin: 0;
  font-size: 14px;
}
</style>
