<script setup>
import {ref} from "vue";

const count = ref(0);
</script>

<template>
  <button @click="count++">Add +1</button>
  <p>{{ count }}</p>
</template>
