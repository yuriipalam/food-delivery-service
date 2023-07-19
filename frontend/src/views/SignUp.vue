<script setup>
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {signUp} from "../api/api";

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

const email = ref('')
const phone = ref('')
const firstName = ref('')
const lastName = ref('')
const password = ref('')
const repeatPassword = ref('')

const errMsg = ref('')

function signUpCustomer(email, phone, firstName, lastName, password, repeatPassword) {
  if (phone !== "" && firstName !== "" && lastName !== "" && password !== "" && repeatPassword !== "") {
    signUp(email, phone, firstName, lastName, password, repeatPassword)
        .then((response) => router.push('/'))
        .catch((error) => errMsg.value = error.message)
  } else {
    errMsg.value = 'All the fields has to be filled in!'
  }
}
</script>

<template>
  <main>
    <form>
      <h1>Create a user account</h1>
      <div>
        <p>Already have an account?</p>
        <router-link :to="{ name: 'SignIn' }">Sign in</router-link>
      </div>
      <div v-if="errMsg !== ''" class="err-msg">
        {{ errMsg }}
      </div>
      <input :placeholder="'Email'" name="email" :required="true" :type="'email'" v-model="email"/>
      <input :placeholder="'Phone'" name="phone" :required="true" v-model="phone"/>
      <input :placeholder="'First name'" name="firstName" :required="true" v-model="firstName"/>
      <input :placeholder="'Last name'" name="lastName" :required="true" v-model="lastName"/>
      <input :placeholder="'Password'" name="password" :required="true" type="password" v-model="password"/>
      <input :placeholder="'Repeat password'" name="repeatPassword" :required="true" type="password"
             v-model="repeatPassword"/>
      <button type="submit" @click.prevent="signUpCustomer(email, phone, firstName, lastName, password, repeatPassword)">Next</button>
    </form>
  </main>
</template>

<style scoped>
@import url("../assets/css/auth.css");
</style>
