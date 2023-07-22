<script setup>
import {formatDate} from "../utils";

const props = defineProps({
  customer: Object,
  orders: []
})

function moneySpent() {
  const prices = props.orders.map(order => parseInt(order.price))
  let sum = 0
  for (let price of prices) {
    sum += price
  }
  return sum
}
</script>

<template>
  <div class="customer-info">
    <p class="customer-full-name">{{ props.customer.first_name }} {{ props.customer.last_name }}</p>
    <p class="customer-contact">
      <svg class="customer-icon" xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24">
        <path fill="currentColor"
              d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 4l-8 5l-8-5V6l8 5l8-5v2z"/>
      </svg>

      {{ props.customer.email }}
    </p>
    <p class="customer-contact">
      <svg class="customer-icon" xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24">
        <path fill="currentColor"
              d="M6.62 10.79c1.44 2.83 3.76 5.14 6.59 6.59l2.2-2.2c.27-.27.67-.36 1.02-.24c1.12.37 2.33.57 3.57.57c.55 0 1 .45 1 1V20c0 .55-.45 1-1 1c-9.39 0-17-7.61-17-17c0-.55.45-1 1-1h3.5c.55 0 1 .45 1 1c0 1.25.2 2.45.57 3.57c.11.35.03.74-.25 1.02l-2.2 2.2z"/>
      </svg>
      {{ props.customer.phone }}
    </p>

    <div class="customer-content">
      <div class="customer-left-side">
        <p class="customer-side-text">Registered at</p>
        <p class="customer-side-text">{{ formatDate(props.customer.created_at) }}</p>
      </div>
      <div class="customer-right-side">
        <p class="customer-side-text">{{ props.orders.length }} orders</p>
        <p class="customer-side-text customer-money-spent">{{ moneySpent() }} HUF</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.customer-info {
  border-radius: 30px;
  padding: 10px 40px;
  margin-left: 20%;
  margin-right: 20%;
  background-color: rgba(254, 114, 76, 0.15);
}

.customer-full-name {
  font-size: 24px;
  font-weight: 500;
  text-align: center;
  margin-bottom: 4px;
}

.customer-contact {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  margin-top: 15px;
  margin-bottom: 15px;
  grid-gap: 7px;
}

.customer-content {
  margin-top: 50px;
  display: flex;
  justify-content: space-between;
  margin-bottom: 5px;
}

.customer-side-text {
  margin: 2px;
}

.customer-created-at {
  margin: 0;
}

.customer-icon {
  color: var(--orange);
}

.customer-left-side {
  margin-left: -15px;
}

.customer-right-side {
  display: flex;
  flex-direction: column;
  margin-right: -15px;
}

.customer-money-spent {
  text-align: right;
  margin-top: 0;
  margin-bottom: 0;
  font-weight: 600;
}

.customer-amount-of-orders {
  margin-top: 0;
  margin-bottom: 0;
}
</style>
