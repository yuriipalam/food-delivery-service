<script setup>
import {getCustomer, getOrders} from "../api/api";
import {computed, onMounted, ref} from "vue";
import {useAuthStore} from "../store";
import {useRoute, useRouter} from "vue-router";
import PrimaryButton from "../components/PrimaryButton.vue";
import OrderCard from "../components/OrderCard.vue";
import CustomerInfo from "../components/CustomerInfo.vue";
import {REFRESH_TOKEN_EXPIRED} from "../api/errors";

const useAuth = useAuthStore()

const router = useRouter()
const route = useRoute()

const customer = ref(Object)
const orders = ref([])

const isProfileRoute = computed(() => {
  return route.name === 'Profile'
})

const isProfileOrdersRoute = computed(() => {
  return route.name === 'ProfileOrders'
})

// const isProfileSettingsRoute = computed(() => {
//   return route.name === 'ProfileSettings'
// })

onMounted(async () => {
  customer.value = await getCustomer().catch(err => {
    if (err.message === REFRESH_TOKEN_EXPIRED) {
      router.push({name: 'SignIn'})
    }
  })
  orders.value = await getOrders()
})
</script>

<template>
  <div class="container-small">
    <main>
      <div class="profile">
        <h1>User profile</h1>
        <div class="control-panel">
          <PrimaryButton @click="router.push({name: 'Profile'})" :class="{'active': isProfileRoute}"
                         class="header-button">Info
          </PrimaryButton>
          <PrimaryButton @click="router.push({name: 'ProfileOrders'})" :class="{'active': isProfileOrdersRoute}"
                         class="header-button">Orders
          </PrimaryButton>
          <PrimaryButton @click="useAuth.signOut(); router.push({name: 'Home'})" class="header-button">Sign out
          </PrimaryButton>
        </div>
        <div v-if="isProfileRoute" class="info-section">
          <CustomerInfo class="info" :customer="customer" :orders="orders"/>
        </div>
        <div v-if="isProfileOrdersRoute" class="orders-section">
          <OrderCard class="order" :order="order" v-for="order in orders"/>
          <p v-if="orders.length === 0" class="not-found">No orders so far =(</p>
        </div>
        <!--      <div v-if="isProfileSettingsRoute" class="settings">-->
        <!--      </div>-->
      </div>
    </main>
  </div>
</template>

<style scoped>
.profile {
  margin-top: 20px;
  border-radius: 30px;
  background: rgba(232, 230, 230, 0.4);
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  justify-self: center;

  padding: 60px 80px 40px;
  flex-grow: 1;
  min-height: 30vh;
}

h1 {
  font-size: 48px;
  font-weight: 500;
  color: var(--blackish);
  text-align: center;
}

.control-panel {
  display: flex;
  justify-content: center;
  grid-gap: 20px;
  margin-bottom: 30px;
}

.info-section, .orders-section {
  height: 50vh;
  overflow: auto;
}

.active {
  color: var(--orange);
  background: none;
  border: 1px solid var(--orange)
}

ul {
  text-align: center;
  margin-top: 40px;
  list-style: none;
  font-size: 20px;
  padding-left: 0;

}

li {
  margin-bottom: 10px;
}

.order {
  margin-bottom: 30px;
}

.not-found {
  font-size: 28px;
  margin-top: 25px;
  text-align: center;
}
</style>
