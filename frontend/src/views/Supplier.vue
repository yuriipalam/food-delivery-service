<script setup>
import SupplierBar from "../components/Supplier/SupplierBar.vue";
import CategoryList from "../components/Supplier/CategoryList.vue";
import SearchBar from "../components/SearchBar.vue";
import ProductCard from "../components/Supplier/ProductCard.vue";
import OrdersBlock from "../components/Supplier/OrdersBlock.vue";
import {computed, onMounted, onUnmounted, ref} from "vue";
import {useRoute, useRouter} from 'vue-router'
import {useFiltersStore} from "../store";
import {getCategoriesBySupplierID, getProductsBySupplierID, getSupplierByID} from "../api/api";
import {getElmHeight} from "../utils";
import {ResponseError} from "../api/errors";
import CategoryListHorizontal from "../components/Supplier/CategoryListHorizontal.vue";

const useFilters = useFiltersStore()

const router = useRouter()

const route = useRoute()
const id = route.params.id

const supplier = ref(Object)
const categories = ref([])
const products = ref([])

// filtering our products
const filteredProducts = computed(() => {
  let productsArray = ref(products.value)

  if (useFilters.searchFor !== '') {
    productsArray.value = products.value.filter((product) => product.name.toLowerCase().includes(useFilters.searchFor.toLowerCase()))
  }
  if (useFilters.selectedCategory === 0) {
    return productsArray.value
  }

  return productsArray.value.filter((product) => product.category_id === useFilters.selectedCategory);
})


const productsLength = computed(() => {
  if (typeof products.value === 'undefined') {
    return "0"
  }
  return products.value.length.toString()
})

// setting ideal height for orders-block
function changeOrdersHeight() {
  const nav = document.querySelector('nav')

  const bar = document.querySelector('.bar')
  const ordersBlock = document.querySelector('.orders-block.orders')
  ordersBlock.style.height = (window.innerHeight - getElmHeight(nav) - getElmHeight(bar)) + 'px'
}

onMounted(async () => {
  changeOrdersHeight()
  window.addEventListener('resize', changeOrdersHeight)

  try {
    supplier.value = await getSupplierByID(id)
    categories.value = await getCategoriesBySupplierID(id)
    products.value = await getProductsBySupplierID(id)
  } catch (err) {
    switch (err.message) {
      case ResponseError.notFound:
        await router.push({name: '404'})
        return
      default:
        await router.push({name: '500'})
        return
    }
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', changeOrdersHeight)
  useFilters.reset()
})
</script>

<template>
  <div class="container">
    <main>
      <SupplierBar :name="supplier.name" :desc="supplier.description" :quantity="productsLength"
                   :imageURL="supplier.image_url"
                   class="bar"></SupplierBar>
      <div class="content">
        <div class="content-categories">
          <CategoryListHorizontal :categories="categories" class="categories-horizontal"/>
          <CategoryList :categories="categories" class="categories"/>
        </div>
        <div class="content-products-and-orders">
          <div class="products">
            <SearchBar :class="'transparent'" :name="supplier.name"/>
            <h2 class="category-name">{{ useFilters.selectedCategoryName }}</h2>
            <div class="products-list">
              <ProductCard v-for="product in filteredProducts" :product="product" :key="product.id"/>
              <span class="no-products-found" v-if="filteredProducts.length === 0">No products found</span>
            </div>
          </div>
          <OrdersBlock class="orders"/>
        </div>

      </div>
    </main>
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
  flex-direction: column;
}

.products {
  width: 685px;
  min-width: 685px;
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
  min-width: 170px;
  height: 100%;
  position: sticky;
  top: 40px;
}

.categories-horizontal {
  flex-basis: 100%;
  flex-grow: 1;
}

.content-products-and-orders {
  display: flex;
  flex-direction: row;
  grid-gap: 30px;
}

.orders {
  flex-grow: 1;
  top: 40px;
  position: sticky;

}

@media screen and (min-width: 1421px) {
  .content {
    flex-direction: row;
  }

  .categories {
    display: block;
  }

  .categories-horizontal {
    display: none;
  }
}

@media screen and (max-width: 1420px) {
.content {
  flex-direction: column;
}

  .categories {
    display: none;
  }

  .categories-horizontal {
    display: flex;
  }

}

@media screen and (max-width: 1230px) {
  .products {
    width: 380px;
    min-width: 380px;
  }
  .product-card {
    width: 100%;
  }
}

@media screen and (max-width: 940px) {
  .content-products-and-orders {
    flex-direction: column-reverse;
  }

  .orders {
    position: initial;
  }

  .products {
    width: 100%;
    min-width: 100%;
  }

  .product-card {
    width: 43%;
  }
}

@media screen and (max-width: 780px) {
  .products-list {
    grid-gap: 20px;
  }
}

@media screen and (max-width: 635px) {
  .product-card {
    width: 100%
  }
  .products-list {
    grid-gap: 30px;
  }
}
</style>
