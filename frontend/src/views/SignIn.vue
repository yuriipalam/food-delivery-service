<script setup>
import {computed, onMounted, reactive, ref} from "vue";
import {signIn} from "../api/api";
import {useRouter} from "vue-router";
import useVuelidate from "@vuelidate/core";
import {required, minLength} from "@vuelidate/validators";

const v$ = useVuelidate(rules, formData)

const formData = reactive({
  email: "",
  password: ""
})

const rules = computed(() => {
  return {
    email: required,
    password: minLength(6)
  }
})

const submitForm = async () => {
  const result = await v$.value.$validate()
  if (result) {
    alert('cool')
  } else {
    alert('fuck')
  }
}

const router = useRouter()

onMounted(async () => {
  await mainHeightSetter()
  window.addEventListener('resize', await mainHeightSetter)
})

async function mainHeightSetter() {
  const navHeight = document.querySelector('nav').offsetHeight
  const main = document.querySelector('main')
  main.style.height = window.innerHeight - navHeight + 'px'
  main.style.marginTop = -navHeight / 2 + 'px'
}

const errMsg = ref('')

const email = ref('')
const password = ref('')

function signInCustomer(email, password) {
  if (email !== "" && password !== "") {
    signIn(email, password)
        .then((response) => router.push('/'))
        .catch((error) => errMsg.value = error.message)
  } else {
    errMsg.value = 'All the fields has to be filled in!'
  }
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
      <input :placeholder="'Email'" name="email" :type="'email'" v-model="formData.email"/>
      <input :placeholder="'Password'" name="password" type="password" v-model="formData.password"/>
      <button type="submit">Next</button>
    </form>
  </main>
</template>

<style scoped>
@import url("../assets/css/auth.css");
</style>
