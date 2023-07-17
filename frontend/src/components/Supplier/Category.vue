<script setup>
const props = defineProps({
  isActive: Boolean
})
</script>

<template>
  <div>
    <hr>
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

hr {
  height: 1px;
  border: 0;
  background-color: rgba(39, 45, 47, 0.3);
}
</style>
