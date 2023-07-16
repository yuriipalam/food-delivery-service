<script setup>

import Order from "./Order.vue";
</script>

<template>
  <div class="orders-block">
    <p class="title">Your order</p>
    <div class="orders-list">
      <Order></Order>
      <Order></Order>
      <Order></Order>
      <Order></Order>
      <Order></Order>
    </div>
    <button type="button">
      <span>Go to checkout</span>
      <span>3880 HUF</span>
    </button>
  </div>
</template>

<style scoped>
.orders-block {
  border-radius: 31px;
  background: rgba(232, 230, 230, 0.4);
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  padding: 30px 20px 20px;
  min-height: 400px;
}

.title {
  font-size: 32px;
  font-weight: 500;
  text-align: center;
  margin-bottom: 25px;
  margin-top: 0;
}

.orders-list {
  overflow: auto;
}

button {
  margin-top: auto;
  font-size: 20px;
  display: flex;
  justify-content: space-between;
  font-weight: 400;
  padding: 13px 25px;
  border: none;
  background-color: rgba(254, 114, 76, 0.4);
  border-radius: 24px;
  transition: opacity 0.7s ease;
}
button span:first-child {
  text-align: left;
}
button span:last-child {
  text-align: right;
}
button:hover {
 opacity: 0.8;
  cursor: pointer;
}
</style>
