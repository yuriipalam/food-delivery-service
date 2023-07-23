<script setup>
const props = defineProps({
      class: String,
    }
)
</script>

<template>
  <button :class="props.class"><slot></slot></button>
</template>

<style scoped>
  button {
    border: none;
    padding: 10px 30px;
    display: flex;
    border-radius: 19px;
    transition: opacity 0.7s ease;
    font-size: 16px;
    font-weight: 600;
  }
  button:hover {
    cursor: pointer;
    opacity: 75%;
  }

  .header-button {
    background: var(--orange);
    color: white;
  }

  .solid-button {
    background: var(--orange);
    color: var(--blackish);
  }

  .transparent-button {
    background: none;
    border: 1px solid var(--orange)!important;
    color: var(--milky);
  }
</style>
