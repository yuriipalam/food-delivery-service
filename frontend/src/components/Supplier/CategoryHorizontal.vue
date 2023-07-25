<script setup>
const props = defineProps({
  isActive: Boolean
})
</script>

<template>
  <div>
    <p :class="{active : props.isActive}">
      <slot></slot>
    </p>
  </div>
</template>

<style scoped>
p {
  font-size: 20px;
  font-weight: 400;
  text-transform: uppercase;
  margin-bottom: 0;
  margin-top: 0;
  transition: opacity 0.7s ease;
  margin-right: 0;
  color: black;
  opacity: 1;
}

p:hover {
  cursor: pointer;
  opacity: 0.7;
}

p.active {
  font-weight: 600;
  pointer-events: none;
  opacity: 1;
  margin-right: 0;
}
</style>
