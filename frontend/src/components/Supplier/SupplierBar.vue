<script setup>


import {truncate} from "../../utils";

const props = defineProps({
  name: String,
  desc: String,
  quantity: String
})
</script>

<template>
  <div class="bar">
    <h1>{{ props.name }}</h1>
    <p class="desc" v-if="props.desc">{{ truncate(props.desc, 100) }}</p>
    <p class="quantity">{{ props.quantity }} available products</p>
  </div>
</template>

<style scoped>
.bar {
  text-align: left;
  border-radius: 30px;
  background: rgba(243, 207, 121, 0.3);
  box-shadow: 4px 4px 7px 5px rgba(0, 0, 0, 0.1);
  padding-top: 40px;
  padding-bottom: 40px;
  padding-left: 65px;
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
</style>
