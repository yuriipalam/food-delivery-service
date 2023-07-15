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
}

p .active {
  font-weight: 600;
}
</style>
