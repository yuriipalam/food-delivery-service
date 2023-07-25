<script setup>
import {useCartStore} from "../../store";
import {truncate} from "../../utils";

const props = defineProps({
  product: Object
})

const useCart = useCartStore()
</script>

<template>
  <div class="product-card">
    <div class="product-left-side">
      <p class="product-name">{{ props.product.name }}</p>
      <p class="product-desc">{{ truncate(props.product.description, 100) }}</p>
      <p class="product-price">{{ props.product.price }} HUF</p>
    </div>
    <div class="product-right-side">
      <button type="button" @click="useCart.addProduct(props.product)">+</button>
      <img :src="props.product.image_url" :alt="props.product.name">
    </div>
  </div>
</template>

<style scoped>
.product-card {
  width: 300px;
  display: flex;
  justify-content: space-between;
  background: rgba(243, 207, 121, 0.3);
  padding: 25px 15px 15px;
  border-radius: 31px;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.15);
}

.product-name {
  font-size: 16px;
  font-weight: 400;
  color: black;
  margin-top: 0;
  margin-bottom: 10px;
}

.product-desc {
  font-weight: 400;
  font-size: 12px;
  opacity: 0.5;
}

.product-left-side {
  display: flex;
  flex-direction: column;
}

.product-price {
  margin-top: auto;
  font-weight: 600;
  font-size: 12px;
  margin-bottom: 0;
}

.product-right-side {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

button {
  margin-top: -10px;
  margin-bottom: 25px;
  align-self: flex-end;
  max-width: 41px;
  font-size: 36px;
  color: var(--blackish);
  font-weight: 400;
  line-height: 11px;
  border: none;
  background-color: var(--orange);
  border-radius: 0 31px 0 10px;
  padding: 10px 10px 8px;
  transition: opacity 0.7s ease;
}

button:hover {
  cursor: pointer;
  opacity: 0.75;
}

img {
  height: 100px;
  margin-bottom: -5px;
}

@media screen and (max-width: 1230px) {

}
</style>
