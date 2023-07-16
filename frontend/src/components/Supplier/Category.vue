<script setup>
const props = defineProps({
  isActive: Boolean
})
</script>

<template>
  <p :class="{active : props.isActive}"><slot></slot></p>
</template>

<style scoped>
p {
  font-size: 22px;
  font-weight: 400;
  text-transform: uppercase;
  margin-bottom: 12px;
  margin-top: 12px;
  transition: opacity 0.7s ease;
}

p:hover {
  cursor: pointer;
  opacity: 0.7;
}

p.active {
  font-weight: 600;
  pointer-events: none;
  opacity: 1;
}
</style>
