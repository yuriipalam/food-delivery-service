<script setup>
import SupplierBar from "../components/Supplier/SupplierBar.vue";
import CategoryList from "../components/Supplier/CategoryList.vue";
import SearchBar from "../components/SearchBar.vue";
import ProductCard from "../components/Supplier/ProductCard.vue";
import OrdersBlock from "../components/Supplier/OrdersBlock.vue";
import {onMounted} from "vue";
onMounted(() => {
  changeOrdersWidth()
  window.addEventListener('resize', changeOrdersWidth)
})

function changeOrdersWidth() {
  const nav = document.querySelector('nav')
  const bar = document.querySelector('.bar')
  const csBar = window.getComputedStyle(bar);
  const heightNav = nav.offsetHeight
  const heightBar = parseFloat(csBar.height) + parseFloat(csBar.marginTop) + parseFloat(csBar.marginBottom);


  const ordersBlock = document.querySelector('.orders-block.orders')
  console.log(heightBar)
  ordersBlock.style.height = (window.innerHeight - heightNav - heightBar) + 'px'
}
</script>

<template>
  <div class="container">
    <SupplierBar class="bar"></SupplierBar>
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
.bar {
  margin-top: 15px;
  margin-bottom: 40px;
}

.content {
  grid-gap: 30px;
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
  //height: auto;
  //max-height: 600px;
}
</style>
