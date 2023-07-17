<script setup>
const props = defineProps({
  name: String,
  quantity: String
})
</script>

<template>
  <div class="bar">
    <h1>{{ props.name }}</h1>
    <p>{{ props.quantity }} available products</p>
  </div>
</template>

<style scoped>
.bar {
  text-align: left;
  border-radius: 30px;
  background: rgba(243, 207, 121, 0.3);
  box-shadow: 4px 4px 7px 5px rgba(0, 0, 0, 0.1);
  padding-top: 40px;
  padding-bottom: 40px;
  padding-left: 65px;
  flex-direction: column;
}

p, h1 {
  margin: 0 0;
}

h1 {
  font-size: 48px;
  font-weight: 500;
  color: var(--blackish)
}

p {
  color: black;
  font-weight: 500;
  font-size: 16px;
}
</style>
