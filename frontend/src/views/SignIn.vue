<script setup>
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {getCustomer, signIn} from "../api/api";
import {useRouter} from "vue-router";
import useVuelidate from "@vuelidate/core";
import {email, minLength, required} from "@vuelidate/validators";
import {useAuthStore} from "../store";

const useAuth = useAuthStore()

const formData = reactive({
  email: "",
  password: ""
})

const rules = computed(() => {
  return {
    email: {email, required},
    password: {required, min: minLength(6)}
  }
})

const errMsg = ref('')

const v$ = useVuelidate(rules, formData)

const submitForm = async () => {
  const result = await v$.value.$validate()
  if (result) {
    signInCustomer(formData.email, formData.password)
  }
}
const router = useRouter()

onMounted(() => {
  mainHeightSetter()
  window.addEventListener('resize', mainHeightSetter)
})

onUnmounted(() => {
  window.removeEventListener('resize', mainHeightSetter)
})

async function mainHeightSetter() {
  const navHeight = document.querySelector('nav').offsetHeight
  const main = document.querySelector('main')
  main.style.height = window.innerHeight - navHeight + 'px'
  main.style.marginTop = -navHeight / 2 + 'px'
}

function signInCustomer(email, password) {
  signIn(email, password).then((response) => {
    useAuth.setTokens(response.access_token, response.refresh_token)
    getCustomer().then((response) => useAuth.setUser(response.id, response.email, response.phone, response.first_name, response.last_name))
    router.push('/')
  }).catch(err => {
    errMsg.value = err.message
  })
}
</script>

<template>
  <main>
    <form @submit.prevent="submitForm">
      <h1>Sign in</h1>
      <div>
        <p>Don't have an account?</p>
        <router-link :to="{ name: 'SignUp' }">Sign up</router-link>
      </div>
      <div v-if="errMsg !== ''" class="err-msg">
        {{ errMsg }}
      </div>
      <span v-for="error in v$.email.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
      <input :placeholder="'Email'" name="email" :required="true" :type="'email'" :class="{'err': v$.email.$error}"
             v-model="formData.email"/>

      <span v-for="error in v$.password.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
      <input :placeholder="'Password'" name="password" :required="true" type="password"
             :class="{'err': v$.password.$error}"
             v-model="formData.password"/>

      <button type="submit">Next</button>
    </form>
  </main>
</template>

<style scoped>
@import url("../assets/css/auth.css");
</style>
