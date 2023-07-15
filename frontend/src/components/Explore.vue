<script setup></script>

<template>
  <div class="explore-block">
    <h2>Explore<br>categories & suppliers</h2>
    <div class="search-input">
      <svg xmlns="http://www.w3.org/2000/svg" width="26" height="22" viewBox="0 0 26 22" fill="none">
        <circle cx="9.04237" cy="8.57557" r="7" transform="rotate(8.9509 9.04237 8.57557)" stroke="#272D2F"/>
        <line x1="14.8744" y1="13.0957" x2="23.2185" y2="21.0212" stroke="#272D2F"/>
      </svg>
      <input type="text">
    </div>
  </div>
</template>

<style scoped>
.explore-block {
  display: flex;
  flex-direction: column;
  padding-left: 60px;
  padding-top: 110px;
  padding-bottom: 60px;
  background-color: var(--pinky);
  border-radius: 40px;
  background-image: url("../assets/img/explore-background.png");
  background-position: right bottom;
  background-size: contain;
  background-repeat: no-repeat;
}

h2 {
  margin-top: 0;
  font-size: 36px;
  font-weight: 400;
}

.search-input {
  position: relative;
  display: flex;
  max-width: 600px;
}

.search-input svg {
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
  background-color: #e1beb4;
  padding-left: 60px;
  padding-top: 17px;
  padding-bottom: 17px;
  font-size: 16px;
}

input:focus {
  background-color: #eec8be;
  outline: none;
}
</style>
