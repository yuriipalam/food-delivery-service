<script setup>

const props = defineProps({
  obj: Object
})
</script>

<template>
  <a class="card" :href="obj.url">
      <img :src="props.obj.image_url" :alt="props.obj.name" width="100" height="100">
    <p>{{ props.obj.name }}</p>
  </a>
</template>

<style scoped>
.card {
  min-height: 100%;
  display: flex;
  align-items: center;
  width: 120px;
  flex-direction: column;
  padding-left: 5px;
  padding-right: 5px;
  background-color: var(--trans-yellowish);
  border-radius: 50px 50px 2px 2px;
  box-shadow: var(--card-shadow);
  text-decoration: none;
  transition: transform 0.4s ease-in-out;
  color: black;
}

.card:hover {
  transform: translateY(5px);
}

img {
  width: 100px;
  height: 100px;
  overflow: hidden;
  margin-bottom: 15px;
  margin-top: -15px;
}

.card p {
  margin: 0;
  padding-bottom: 15px;
}
</style>
