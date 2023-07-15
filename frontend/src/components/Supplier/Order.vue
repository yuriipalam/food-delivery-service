<script setup>

</script>

<template>
<div class="order">
  <div class="order-info">
    <p class="order-name">Maestro Honey McMenu</p>
    <p class="order-price">3880 huf</p>
  </div>
  <div class="order-controls">
    <button class="minus-btn"><span class="minus-sign"></span></button>
    <span class="quantity">2</span>
    <button class="plus-btn"><span class="minus-sign"></span><span class="minus-sign rotated"></span></button>
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
  flex-direction: column;
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
}

.minus-sign {
  display: block;
  width: 10px;
  height: 2px;
  position: absolute;
  background-color: var(--blackish);
  transform: translateY(-1px);
}

.rotated {
  transform: rotate(90deg) translateX(-1px);
  position: absolute;
}

.minus-btn, .plus-btn {
  height: 30px;
  width: 30px;
  font-size: 32px;
  color: var(--blackish);
  font-weight: 400;
  line-height: 11px;
  border: none;
  background-color: rgba(254, 114, 76, 0.4);
  border-radius: 50%;
  padding: 10px 10px;
}

.quantity {
  margin: 0 8px;
  background-color: rgba(254, 114, 76, 0.4);
  padding: 8px 12px;
  border-radius: 12px;
}
</style>
