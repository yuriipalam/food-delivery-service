<script setup>
import {useFiltersStore} from "../store";

const props = defineProps({
  name: String,
  class: String
})

const store = useFiltersStore()
</script>

<template>
  <div class="search-bar">
    <svg xmlns="http://www.w3.org/2000/svg" width="26" height="22" viewBox="0 0 26 22" fill="none">
      <circle cx="9.04237" cy="8.57557" r="7" transform="rotate(8.9509 9.04237 8.57557)" stroke="#272D2F"/>
      <line x1="14.8744" y1="13.0957" x2="23.2185" y2="21.0212" stroke="#272D2F"/>
    </svg>
    <input type="text" v-model="store.searchFor" :class="props.class" :placeholder="typeof props.name === 'undefined' ? '': 'Search in ' + props.name">
  </div>
</template>

<style scoped>
.search-bar {
  position: relative;
  display: flex;
}

.search-bar svg {
  position: absolute;
  top: 50%;
  left: 30px;
  transform: translateY(-50%);
  pointer-events: none;
  z-index: 1;
}

input {
  display: flex;
  border-radius: 24px;
  border: 1px solid #494d4f;
  opacity: 1;
  width: 100%;
  padding-left: 70px;
  padding-top: 17px;
  padding-bottom: 17px;
  font-size: 16px;
}

input:focus {
  outline: none;
}

input.pinky {
  background-color: #e1beb4;
}
input.pinky:focus  {
  background-color: #eec8be;
}

input.transparent {
  background-color: rgba(255, 255, 255, 0);
}
input.transparent:focus {
  background-color: rgba(239, 238, 238, 0.4);
}
</style>
