<script setup>

</script>

<template>
  <div class="card">
    <div class="circle-mask">
      <img src="../../assets/img/mcdonalds-icon.png" alt="" width="100" height="100">
    </div>
    <p>McDonald's</p>
  </div>
</template>

<style scoped>
.card {
  display: flex;
  align-items: center;
  width: 120px;
  flex-direction: column;
  padding-left: 5px;
  padding-right: 5px;
  background-color: rgba(243, 207, 121, 0.4);
  border-radius: 50px 50px 2px 2px;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
}

.circle-mask {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  overflow: hidden;
  margin-bottom: 15px;
  margin-top: -15px;
}

.circle-mask img {
  object-fit: cover;
}

.card p {
  margin: 0;
  padding-bottom: 15px;
}
</style>
