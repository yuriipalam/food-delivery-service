<script setup>
import {useCartStore} from "../../store";

const props = defineProps({
  product: Object
})

const useCart = useCartStore()
</script>

<template>
  <div class="order">
    <div class="order-info">
      <img class="order-img" :src="props.product.image_url" :alt="props.product.name">
      <div class="order-detail">
        <p class="order-name">{{ props.product.name }}</p>
        <p class="order-price">{{ useCart.getProductTotalPrice(props.product.id) }} huf</p>
      </div>
    </div>
    <div class="order-controls">
      <button class="minus-btn" @click="useCart.reduceQuantity(props.product.id)"><span class="minus-sign"></span>
      </button>
      <span class="quantity">{{ useCart.getQuantity(props.product.id) }}</span>
      <button class="plus-btn" @click="useCart.increaseQuantity(props.product.id)"><span class="minus-sign"></span><span
          class="minus-sign rotated"></span></button>
    </div>
  </div>
</template>

<style scoped>
.order {
  display: flex;
  padding-bottom: 10px;
  padding-top: 10px;
  justify-content: space-between;
  align-items: center;
}

.order-info {
  display: flex;
  align-items: center;
}

.order-img {
  height: 60px;
  margin-right: 15px;
}

.order-name {
  font-weight: 400;
  font-size: 20px;
  color: black;
  margin-bottom: 8px;
  margin-top: 0;
}

.order-price {
  margin-top: 0;
  margin-bottom: 0;
  color: var(--orange);
  font-size: 16px;
  font-weight: 500;
}

.order-controls {
  display: flex;
  align-items: center;
  margin-left: 20px;
  margin-right: 10px;
}

.minus-sign {
  display: block;
  width: 10px;
  height: 2px;
  position: relative;
  background-color: var(--blackish);
}

.rotated {
  transform: rotate(90deg) translateX(-1.5px);
  position: relative;
}

.minus-btn, .plus-btn {
  height: 30px;
  width: 30px;
  font-size: 32px;
  color: var(--blackish);
  font-weight: 400;
  line-height: 11px;
  border: none;
  background-color: var(--trans-orange);
  border-radius: 50%;
  padding: 10px 10px;
  transition: opacity 0.7s ease;
}

.minus-btn:hover, .plus-btn:hover {
  opacity: 0.75;
  cursor: pointer;
}

.quantity {
  margin: 0 8px;
  background-color: var(--trans-orange);
  padding: 8px 12px;
  border-radius: 12px;
}

@media screen and (max-width: 480px) {
  .order-name {
    font-size: 16px;
  }

  .order-price {
    font-size: 14px;
  }

  .order-controls {
    margin-left: 10px;
  }

  .order-img {
    margin-right: 5px;
  }

  .quantity {
    font-size: 15px;
  }
}
</style>
