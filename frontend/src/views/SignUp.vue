<script setup>
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {signIn, signUp} from "../api/api";
import useVuelidate from "@vuelidate/core";
import {email, maxLength, minLength, required} from "@vuelidate/validators";
import {hideFooter, showFooter} from "../utils";

const router = useRouter()

const formData = reactive({
  email: "",
  phone: "",
  firstName: "",
  lastName: "",
  password: "",
  repeatPassword: ""
})

const rules = computed(() => {
  return {
    email: {email, required},
    phone: {min: minLength(9), max: maxLength(14), required},
    firstName: {required},
    lastName: {required},
    password: {min: minLength(6), required},
    repeatPassword: {min: minLength(6), required}
  }
})

const errMsg = ref('')

const v$ = useVuelidate(rules, formData)

const submitForm = async () => {
  const result = await v$.value.$validate()
  if (result) {
    try {
      await signUp(formData.email, formData.phone, formData.firstName, formData.lastName, formData.password, formData.repeatPassword)
      await signIn(formData.email, formData.password)
      await router.push({name: 'Home'})
    } catch (err) {
      errMsg.value = err.message
    }
  }
}

onMounted(() => {
  hideFooter()
})

onUnmounted(() => {
  showFooter()
})
</script>

<template>
  <main>
    <div class="container">
      <form @submit.prevent="submitForm">
        <h1>Create a user account</h1>
        <div>
          <p>Already have an account?</p>
          <router-link :to="{ name: 'SignIn' }">Sign in</router-link>
        </div>
        <div v-if="errMsg !== ''" class="err-msg">
          {{ errMsg }}
        </div>
        <span v-for="error in v$.email.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
        <input :placeholder="'Email'" name="email" :required="true" :type="'email'" :class="{'err': v$.email.$error}"
               v-model="formData.email"/>

        <span v-for="error in v$.phone.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
        <input :placeholder="'Phone'" name="phone" :required="true" :class="{'err': v$.phone.$error}"
               v-model="formData.phone"/>

        <span v-for="error in v$.firstName.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
        <input :placeholder="'First name'" name="firstName" :required="true" :class="{'err': v$.firstName.$error}"
               v-model="formData.firstName"/>

        <span v-for="error in v$.lastName.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
        <input :placeholder="'Last name'" name="lastName" :required="true" :class="{'err': v$.lastName.$error}"
               v-model="formData.lastName"/>

        <span v-for="error in v$.password.$errors" :key="error.$uid" class="err-span-msg">{{ error.$message }}</span>
        <input :placeholder="'Password'" name="password" :required="true" type="password"
               :class="{'err': v$.password.$error}" v-model="formData.password"/>

        <span v-for="error in v$.repeatPassword.$errors" :key="error.$uid" class="err-span-msg">{{
            error.$message
          }}</span>
        <input :placeholder="'Repeat password'" name="repeatPassword" :class="{'err': v$.repeatPassword.$error}"
               :required="true" type="password"
               v-model="formData.repeatPassword"/>

        <button type="submit">Next</button>
      </form>
    </div>
  </main>
</template>

<style scoped>
@import url("../assets/css/auth.css");
</style>
