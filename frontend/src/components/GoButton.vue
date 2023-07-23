<script setup>

</script>

<template>
  <button type="button">
    <slot></slot>
  </button>
</template>

<style scoped>
button {
  font-size: 20px;
  font-weight: 400;
  padding: 13px 25px;
  border: none;
  background-color: var(--trans-orange);
  border-radius: 24px;
  transition: opacity 0.7s ease;
  display: flex;
  justify-content: space-between;
}

button {
  text-align: center;
}

button:hover {
  opacity: 0.8;
  cursor: pointer;
}
</style>
