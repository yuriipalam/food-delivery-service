<script setup></script>

<template>
  <div class="modal">
    <slot></slot>
  </div>
  <div id="overlay"></div>
</template>

<style scoped>
.modal {
  display: none;
  padding: 35px 40px;
  opacity: 0;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 20px;
  z-index: 99;
  background-color: white;
  transition: opacity 400ms;
}

.active {
  display: block;
  opacity: 1;
  pointer-events: all;
}

#overlay {
  position: fixed;
  z-index: 3;
  opacity: 0;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  pointer-events: none;
  transition: 400ms ease-in-out;
}

#overlay.active {
  display: block;
  opacity: 1;
  pointer-events: all;
}
</style>
