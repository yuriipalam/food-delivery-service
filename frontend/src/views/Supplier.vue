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
      <div class="products">
        <SearchBar :class="'transparent'" :placeholder="'Search in McDonald\'s'"></SearchBar>
        <h2 class="category-name">McMenu</h2>
        <div class="products-list">
          <ProductCard :name="'BigMac Bacon McMenu'" :desc="'lorem lorem lorem'" :price="5142"></ProductCard>
          <ProductCard :name="'BigMac Bacon McMenu'" :desc="'lorem lorem lorem'" :price="5142"></ProductCard>
          <ProductCard :name="'BigMac Bacon McMenu'" :desc="'lorem lorem lorem'" :price="5142"></ProductCard>
          <ProductCard :name="'BigMac Bacon McMenu'" :desc="'lorem lorem lorem'" :price="5142"></ProductCard>
        </div>
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
  grid-gap: 30px;
  margin-top: 40px;
  display: flex;
  justify-content: space-between;
}

.products {
  display: flex;
  flex-direction: column;
  max-width: 685px;
}

.products-list {
  grid-gap: 25px;
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
}

.category-name {
  color: var(--blackish);
  font-size: 22px;
  font-weight: 700;
  padding-left: 15px;
}

.orders {
  height: auto;
  max-height: 600px;
}
</style>
