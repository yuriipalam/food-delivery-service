<script setup>

import OrderField from "./OrderField.vue";
import GoButton from "../GoButton.vue";
import {useCartStore} from "../../store";

const useCart = useCartStore()
</script>

<template>
  <div class="orders-block">
    <p class="title">Your order</p>
    <div class="orders-list">
      <OrderField :product="product.prod" v-for="product in useCart.products"/>
    </div>
    <GoButton>
      <span>Go to checkout</span>
      <span>{{ useCart.products.map(prod => prod.prod.price * prod.quantity).reduce((total, current) => total + current, 0) }} HUF</span>
    </GoButton>
  </div>
</template>

<style scoped>
.orders-block {
  border-radius: 30px;
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
  display: flex;
  justify-content: space-between;
}
button span:first-child {
  text-align: left;
}
button span:last-child {
  text-align: right;
}

</style>
