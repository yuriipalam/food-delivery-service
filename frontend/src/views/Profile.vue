<script setup>
import {getCustomer, getOrders} from "../api/api";
import {computed, onMounted, ref} from "vue";
import {useAuthStore} from "../store";
import {useRoute, useRouter} from "vue-router";
import PrimaryButton from "../components/PrimaryButton.vue";
import OrderCard from "../components/OrderCard.vue";
import CustomerInfo from "../components/CustomerInfo.vue";
import {ResponseError} from "../api/errors";

const useAuth = useAuthStore()

const router = useRouter()
const route = useRoute()

const customer = ref({})
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
  try {
    customer.value = await getCustomer()
    orders.value = await getOrders()
  } catch (err) {
    switch (err.message) {
      case ResponseError.notFound:
        await router.push({name: '404'})
        return
      case ResponseError.sessionExpired:
        await router.push({name: 'SignIn'})
        return
      default:
        await router.push({name: '500'})
        return
    }
  }
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
          <PrimaryButton @click="useAuth.signOut(); router.push({name: 'Home'})" class="sign-out-button">Sign out
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
  background: var(--trans-milky);
  box-shadow: var(--container-shadow);
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

.info {
  margin-left: 20%;
  margin-right: 20%;
}

button.active {
  color: var(--orange);
  background: none;
  border: 1px solid var(--orange);
}

.order {
  margin-bottom: 30px;
}

.not-found {
  font-size: 28px;
  margin-top: 25px;
  text-align: center;
}

.sign-out-button {
  background: #ffc5b4;
  color: #444444;
}

@media screen and (max-width: 1024px) {
  .profile {
    padding-left: 30px;
    padding-right: 30px;
  }
}

@media screen and (max-width: 768px) {
  h1 {
    margin-top: 0;
  }

  .profile {
    padding-left: 10px;
    padding-right: 10px;
  }

  .info {
    margin-left: 40px;
    margin-right: 40px;
  }
}

@media screen and (max-width: 480px) {
  .control-panel {
    flex-flow: wrap;
    grid-gap: 10px;
  }

  .info {
    margin-left: 10px;
    margin-right: 10px;
  }
}
</style>
