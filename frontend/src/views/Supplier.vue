<script setup>

import SupplierBar from "../components/Supplier/SupplierBar.vue";
import CategoryList from "../components/Supplier/CategoryList.vue";
import SearchBar from "../components/SearchBar.vue";
import ProductCard from "../components/Supplier/ProductCard.vue";
import OrdersBlock from "../components/Supplier/OrdersBlock.vue";
</script>

<template>
  <div class="container">
    <SupplierBar></SupplierBar>
    <div class="content">
      <CategoryList class="categories"></CategoryList>
      <div class="products-list">
        <SearchBar :class="'transparent'" :placeholder="'Search in McDonald\'s'"></SearchBar>
        <h2 class="category-name">McMenu</h2>
        <ProductCard :name="'BigMac Bacon McMenu'" :desc="'lorem lorem lorem'" :price="5142"></ProductCard>
      </div>
      <OrdersBlock class="orders"></OrdersBlock>
    </div>
  </div>
</template>

<style scoped>
.container {
  margin-top: 15px;
}

.content {
  margin-top: 40px;
  display: flex;
  justify-content: space-between;
}

.products-list {
  display: flex;
  flex-direction: column;
}

.category-name {
  color: var(--blackish);
  font-size: 22px;
  font-weight: 700;
  padding-left: 15px;
}

.orders {
  flex-shrink: 1;
}
</style>
