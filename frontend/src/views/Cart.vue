<script setup>

import OrderField from "../components/Supplier/OrderField.vue";
import GoButton from "../components/GoButton.vue";
import {useCartStore} from "../store";

const useCart = useCartStore()
</script>

<template>
  <div class="container-small">
    <div class="cart">
      <h1>Your cart</h1>
      <div class="orders-list">
        <OrderField :product="product.prod" v-for="product in useCart.products"/>
      </div>
      <GoButton>Confirm order</GoButton>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-size: 42px;
  font-weight: 400;
  text-align: center;
  margin-bottom: 40px;
  margin-top: 0;
}

.orders-list {
  overflow: auto;
  height: 40vh;
}

.cart {
  margin-top: 70px;
  border-radius: 30px;
  background: rgba(232, 230, 230, 0.4);
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  justify-self: center;
  padding: 60px 80px;
  flex-grow: 1;
  min-height: 35vh;
}

button {
  margin-top: auto;
}
</style>
