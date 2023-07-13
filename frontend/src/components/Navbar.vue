<script setup>
import NavLink from "./NavLink.vue";
import PrimaryButton from "./PrimaryButton.vue";
</script>

<template>
  <nav>
    <div class="container">
      <NavLink :route-name="'Home'"><strong><b>FoodJet</b></strong></NavLink>
      <div class="header-content">
        <ul class="navbar-nav">
          <li>
            <NavLink :route-name="'Suppliers'">Suppliers</NavLink>
          </li>
          <li>
            <NavLink :route-name="'Categories'">Categories</NavLink>
          </li>
          <li>
            <NavLink :route-name="'ContactUs'">Contact us</NavLink>
          </li>
        </ul>
        <PrimaryButton :class="'header-button'">Sign up</PrimaryButton>
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav {
  padding-top: 20px;
  padding-bottom: 20px;
}

nav > .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-content {
  display: flex;
  align-items: center;
}

ul {
  margin: 0 0;
  padding-left: 0;
  list-style: none;
  display: flex;
}

ul li {
  margin-right: 15px;
}
</style>