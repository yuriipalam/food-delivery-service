<script setup>

import {formatDate} from "../utils";

const props = defineProps({
  order: Object
})

function formattedProducts(products, supplierID) {
  return products
      .filter(prod => prod.product_supplier_id === supplierID)
      .map(prod => `${prod.product_quantity} ${prod.product_name}`)
      .join(", ");
}
</script>

<template>
  <div class="order">
    <div class="order-header">
      <p class="order-id">ORDER #{{ order.id }}</p>
    </div>
    <div class="order-supplier" v-for="supplier in order.suppliers">
      <div class="order-supplier-title">
        <img :src="supplier.supplier_image_url" height="50" :alt="supplier.supplier_name">
        <p class="order-supplier-name">{{ supplier.supplier_name }}</p>
      </div>
      <p class="products">
        {{ formattedProducts(order.products, supplier.supplier_id) }}
      </p>
    </div>
    <div class="order-info">
      <div class="order-info-left-side">
        <p class="order-recipient">{{ order.recipient_full_name }},</p>
        <p class="order-address">{{ order.address }}</p>
      </div>
      <div class="order-info-right-side">
        <p class="order-price">{{ order.price }} HUF</p>
        <p class="order-created-at">{{ formatDate(order.created_at) }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.order {
  border-radius: 30px;
  padding: 10px 40px;
  margin-left: 20%;
  margin-right: 20%;
  background-color: var(--very-trans-orange);
}

.order-id {
  font-size: 16px;
  font-weight: 400;
}

.order-header {
  display: flex;
  grid-gap: 10px;
}

.order-supplier {
  margin-top: 15px;
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  flex-direction: column;
}

.order-supplier-title {
  display: flex;
  align-items: center;
  grid-gap: 10px;
}

.order-supplier-name {
  font-size: 16px;
  font-weight: 500;
}

.products {
  margin: 10px 20px 40px;
  font-weight: 500;
  font-size: 14px;
}

.order-info {
  justify-content: space-between;
  display: flex;
  align-items: flex-end;
}

.order-info-left-side {
  margin-left: -20px;
  font-size: 14px;
}

.order-recipient {
  margin-top: 0;
  margin-bottom: 3px;
}

.order-address {
  margin: 0;
}

.order-info-right-side {
  margin-right: -20px;
}

.order-price {
  text-align: right;
  margin-top: 0;
  margin-bottom: 10px;
  font-weight: 600;
}

.order-created-at {
  margin-top: 0;
  margin-bottom: 0;
  font-size: 14px;
}
</style>
