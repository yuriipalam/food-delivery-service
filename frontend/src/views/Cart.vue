<script setup>

import OrderField from "../components/Supplier/OrderField.vue";
import GoButton from "../components/GoButton.vue";
import {useAuthStore, useCartStore} from "../store";
import {computed, reactive, ref} from "vue";
import {minLength, required} from "@vuelidate/validators";
import useVuelidate from "@vuelidate/core";

const useCart = useCartStore()
const useAuth = useAuthStore()

const formData = reactive({
  recipient: "",
  address: ""
})

const rules = computed(() => {
  return {
    recipient: {required, min: minLength(10)},
    address: {required, min: minLength(10)}
  }
})

const errMsg = ref('')

const v$ = useVuelidate(rules, formData)

const submitForm = async () => {
  const result = await v$.value.$validate()
  console.log(result)
  if (result) {
    // todo createOrder
  }
}
</script>

<template>
  <div class="container-small">
    <div class="cart">
      <h1>Your cart</h1>
      <div class="orders-list">
        <OrderField :product="product.product" v-for="product in useCart.products"/>
      </div>
      <div class="controls">
        <form @submit.prevent="submitForm">
          <div v-if="errMsg !== ''" class="err-msg">
            {{ errMsg }}
          </div>
          <input type="text" :value="useAuth.phoneRef" disabled>

          <span v-for="error in v$.address.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
          <input type="text" placeholder="Address" :class="{'err': v$.address.$error}" name="address" v-model="formData.address"/>

          <span v-for="error in v$.recipient.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
          <input type="text" placeholder="Recipient full name" :class="{'err': v$.recipient.$error}" name="recipient" v-model="formData.recipient">

          <div class="payment">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 32 32">
              <path fill="currentColor"
                    d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm0 26a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>
              <path fill="currentColor" d="M15 8h2v11h-2zm1 14a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 16 22z"/>
            </svg>
            <p>Payment is only by cash!</p>
          </div>
          <GoButton :type="'submit'">
            <span>Confirm order</span>
            <span>{{ useCart.getTotalPrice() }} HUF</span>
          </GoButton>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-size: 42px;
  font-weight: 400;
  text-align: center;
  margin-bottom: 40px;
  margin-top: 0;
}

.orders-list {
  overflow: auto;
  height: 35vh;
}

.cart {
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

.controls {
  display: flex;
  flex-direction: column;
  margin-right: 5%;
  margin-left: 5%;
}

form {
  display: flex;
  flex-direction: column;
}

.payment {
  justify-content: center;
  align-items: center;
  display: flex;
  color: #8f1515;
  grid-gap: 10px;
  margin-bottom: 5px;
}

input {
  flex-grow: 1;
  padding: 10px;
  font-size: 16px;
  border: none;
  border-bottom: 2px solid rgba(39, 45, 47, 0.2);
  outline: none;
  background-color: transparent;
  margin-bottom: 20px;
}
input:focus {
  border-bottom: 2px solid rgba(39, 45, 47, 0.3);
}
input.err {
  border-bottom: 2px solid #ff4a4a;
}

input, span {
  margin-right: 5%;
  margin-left: 5%;
}
</style>
