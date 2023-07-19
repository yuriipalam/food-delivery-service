<script setup>
import SupplierBar from "../components/Supplier/SupplierBar.vue";
import CategoryList from "../components/Supplier/CategoryList.vue";
import SearchBar from "../components/SearchBar.vue";
import ProductCard from "../components/Supplier/ProductCard.vue";
import OrdersBlock from "../components/Supplier/OrdersBlock.vue";
import {computed, onMounted, reactive, ref} from "vue";
import {useRoute} from 'vue-router'
import {useFiltersStore} from "../store";
import {getSupplierByID, getSupplierCategoriesByID, getSupplierProductsByID} from "../api/api";

const store = useFiltersStore()

const route = useRoute()
const id = route.params.id

const supplier = ref(Object)
const categories = ref([])
const products = ref([])

const filteredProducts = computed ( () => {
  let productsArray = ref(products.value)

  if (store.searchFor !== '') {
    productsArray.value = products.value.filter((product) => product.name.toLowerCase().includes(store.searchFor.toLowerCase()))
  }
  if (store.selectedCategory === 0) {
    return productsArray.value
  }

  return productsArray.value.filter((product) => product.category_id === store.selectedCategory);
})

onMounted(async () => {
  changeOrdersHeight()
  window.addEventListener('resize', changeOrdersHeight)
  supplier.value = await getSupplierByID(id)
  categories.value = await getSupplierCategoriesByID()
  products.value = await getSupplierProductsByID()
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
    <SupplierBar :name="supplier.name" :quantity="productsLength" class="bar"></SupplierBar>
    <div class="content">
      <CategoryList :categories="categories" class="categories"></CategoryList>
      <div class="products">
        <SearchBar :class="'transparent'" :name="supplier.name"></SearchBar>
        <h2 class="category-name">{{ store.selectedCategoryName }}</h2>
        <div class="products-list">
          <ProductCard v-for="product in filteredProducts" :product="product" :key="product.id"></ProductCard>
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
  flex-grow: 1;
  display: flex;
  flex-direction: column;
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

.categories {
  height: 100%;
  position: sticky;
  top: 40px;
}

.orders {
  top: 40px;
  position: sticky;
}
</style>
