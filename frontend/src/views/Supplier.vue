<script setup>
import SupplierBar from "../components/Supplier/SupplierBar.vue";
import CategoryList from "../components/Supplier/CategoryList.vue";
import SearchBar from "../components/SearchBar.vue";
import ProductCard from "../components/Supplier/ProductCard.vue";
import OrdersBlock from "../components/Supplier/OrdersBlock.vue";
import {computed, onMounted, ref} from "vue";
import {useRoute} from 'vue-router'
import {useCartStore, useFiltersStore} from "../store";
import {getSupplierByID, getSupplierCategoriesByID, getSupplierProductsByID} from "../api/api";

const useFilters = useFiltersStore()
const useCart = useCartStore()

const route = useRoute()
const id = route.params.id

const supplier = ref(Object)
const categories = ref([])
const products = ref([])

const filteredProducts = computed ( () => {
  let productsArray = ref(products.value)

  if (useFilters.searchFor !== '') {
    productsArray.value = products.value.filter((product) => product.name.toLowerCase().includes(useFilters.searchFor.toLowerCase()))
  }
  if (useFilters.selectedCategory === 0) {
    return productsArray.value
  }

  return productsArray.value.filter((product) => product.category_id === useFilters.selectedCategory);
})

onMounted(async () => {
  changeOrdersHeight()
  window.addEventListener('resize', changeOrdersHeight)
  supplier.value = await getSupplierByID(id)
  categories.value = await getSupplierCategoriesByID(id)
  products.value = await getSupplierProductsByID(id)
})

function changeOrdersHeight() {
  const nav = document.querySelector('nav')
  const bar = document.querySelector('.bar')

  const ordersBlock = document.querySelector('.orders-block.orders')
  ordersBlock.style.height = (window.innerHeight - getElmHeight(nav) - getElmHeight(bar)) + 'px'
}

function getElmHeight(node) {
  const list = [
    'margin-top',
    'margin-bottom',
    'border-top',
    'border-bottom',
    'padding-top',
    'padding-bottom',
    'height'
  ]

  const style = window.getComputedStyle(node)
  return list
      .map(k => parseInt(style.getPropertyValue(k), 10))
      .reduce((prev, cur) => prev + cur)
}

const productsLength = computed(() => {
  if (typeof products.value === 'undefined') {
    return "0"
  }
  return products.value.length.toString()
})
</script>

<template>
  <div class="container">
    <SupplierBar :name="supplier.name" :desc="supplier.description" :quantity="productsLength" class="bar"></SupplierBar>
    <div class="content">
      <CategoryList :categories="categories" class="categories"></CategoryList>
      <div class="products">
        <SearchBar :class="'transparent'" :name="supplier.name"></SearchBar>
        <h2 class="category-name">{{ useFilters.selectedCategoryName }}</h2>
        <div class="products-list">
          <ProductCard v-for="product in filteredProducts" :product="product" :key="product.id"></ProductCard>
          <span class="no-products-found" v-if="filteredProducts.length === 0">No products found</span>
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
  width: 685px;
  display: flex;
  flex-direction: column;
}

.products-list {
  grid-gap: 25px;
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
  justify-content: center;
}

.no-products-found {
  font-size: 28px;
  margin-top: 25px;
}

.category-name {
  color: var(--blackish);
  font-size: 22px;
  font-weight: 700;
  padding-left: 15px;
}

.categories {
  width: 170px;
  height: 100%;
  position: sticky;
  top: 40px;
}

.orders {
  flex-grow: 1;
  top: 40px;
  position: sticky;
}
</style>
