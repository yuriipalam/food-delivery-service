<script setup>
import OrderField from "./OrderField.vue";
import GoButton from "../GoButton.vue";
import {useCartStore} from "../../store";
import {useRouter} from "vue-router";

const useCart = useCartStore()

const router = useRouter()
</script>

<template>
  <div class="orders-block">
    <p class="orders-title">Your order</p>
    <div class="orders-list">
      <OrderField :product="product.product" v-for="product in useCart.products"/>
    </div>
    <GoButton @click="router.push({'name': 'Cart'})">
      <span>Go to checkout</span>
      <span>{{ useCart.getTotalPrice() }} HUF</span>
    </GoButton>
  </div>
</template>

<style scoped>
.orders-block {
  border-radius: 30px;
  background: rgba(232, 230, 230, 0.4);
  box-shadow: var(--container-shadow);
  display: flex;
  flex-direction: column;
  padding: 30px 20px 20px;
  min-height: 400px;
}

.orders-title {
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
}
</style>
