<script setup>
import "@/assets/css/main.css"

const props = defineProps({
      "type": String,
      "class": String,
    }
)
</script>

<template>
  <button :type="props.type" :class="props.class"><slot></slot></button>
</template>

<style scoped>
  button {
    border: none;
    padding: 13px 50px;
    display: flex;
    border-radius: 19px;
    transition: opacity 0.7s ease;
    font-size: 20px;
    font-weight: 600;
  }
  button:hover {
    cursor: pointer;
    opacity: 85%;
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
  }
</style>