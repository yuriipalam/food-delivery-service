<script setup>
const props = defineProps({
      class: String,
    }
)
</script>

<template>
  <button :class="props.class">
    <slot></slot>
  </button>
</template>

<style scoped>
button {
  border: none;
  padding: 10px 30px;
  display: flex;
  border-radius: 19px;
  transition: background 0.4s ease-in-out,
  color 0.4s ease-in-out,
  outline 0.4s ease-in-out;
  font-size: 16px;
  font-weight: 600;
}

button:hover {
  cursor: pointer;
}

.header-button {
  background: var(--orange);
  color: white;
}

.header-button:hover {
  background: white;
  color: var(--orange);
}

.solid-button {
  background: var(--orange);
  color: var(--blackish);
}

.solid-button:hover {
  color: var(--orange);
  background: var(--blackish);
  outline: var(--orange) 1px solid;
}

.transparent-button {
  background: none;
  border: 1px solid var(--orange);
  color: var(--milky);
}

.transparent-button:hover {
  background: var(--orange);
  color: var(--blackish);
}
</style>
