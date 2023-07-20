<script setup>
import {getCustomer} from "../api/api";
import {computed, onMounted, ref} from "vue";
import {useAuthStore} from "../store";
import {useRoute, useRouter} from "vue-router";
import PrimaryButton from "../components/PrimaryButton.vue";

const router = useRouter()
const useAuth = useAuthStore()
const route = useRoute()

const customer = ref(Object)

const isProfileRoute = computed(() => {
  return route.name === 'Profile'
})

const isProfileOrdersRoute = computed(() => {
  return route.name === 'ProfileOrders'
})

const isProfileSettingsRoute = computed(() => {
  return route.name === 'ProfileSettings'
})

onMounted(async () => {
  customer.value = await getCustomer()
})
</script>

<template>
  <main>
    <div class="container-small">
      <h1>User profile</h1>
      <div class="control-panel">
        <PrimaryButton @click="router.push({name: 'Profile'})" :class="{'active': isProfileRoute}"
                       class="header-button">Info
        </PrimaryButton>
        <PrimaryButton @click="router.push({name: 'ProfileOrders'})" :class="{'active': isProfileOrdersRoute}"
                       class="header-button">Orders
        </PrimaryButton>
        <PrimaryButton @click="useAuth.signOut(); router.push({name: 'Home'})" class="solid-button">Sign out
        </PrimaryButton>
      </div>
      <div v-if="isProfileRoute" class="info">
        <ul>
          <li>Account ID: {{ customer.id }}</li>
          <li>{{ customer.first_name }} {{ customer.last_name }}</li>
          <li>{{ customer.email }}</li>
          <li>{{ customer.phone }}</li>
        </ul>
      </div>
      <div v-if="isProfileOrdersRoute" class="orders">
        <h2>orders</h2>
      </div>
      <!--      <div v-if="isProfileSettingsRoute" class="settings">-->
      <!--      </div>-->
    </div>
  </main>
</template>

<style scoped>
.container-small {
  margin-top: 50px;
  padding-top: 30px;
  padding-bottom: 30px;
  background-color: rgba(232, 230, 230, 0.4);
  border-radius: 50px;
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
}

.active {
  color: var(--blackish);
  background-color: white;
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
</style>
