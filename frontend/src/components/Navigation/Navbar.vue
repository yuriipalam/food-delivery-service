<script setup>
import NavLink from "./NavLink.vue";
import PrimaryButton from "../PrimaryButton.vue";
import {useRouter} from "vue-router";
import {useAuthStore} from "../../store";

const router = useRouter()
const useAuth = useAuthStore()

const props = defineProps({
  isLight: Boolean,
})

function getFill() {
  if (props.isLight) {
    return 'var(--milky)'
  }
  return 'var(--blackish)'
}
</script>

<template>
  <nav>
    <div class="container">
      <NavLink :route-name="'Home'" :is-light="props.isLight"><strong><b>FoodJet</b></strong></NavLink>
      <ul>
        <li>
          <NavLink :route-name="'Suppliers'" :is-light="props.isLight">Suppliers</NavLink>
        </li>
        <li>
          <NavLink :route-name="'Categories'" :is-light="props.isLight">Categories</NavLink>
        </li>
        <li>
          <NavLink :route-name="'ContactUs'" :is-light="props.isLight">Contact us</NavLink>
        </li>
      </ul>
      <div class="navbar-content">
        <NavLink :route-name="'Placeholder'">
          <!-- default height for svg is 33 !-->
          <svg xmlns="http://www.w3.org/2000/svg" width="26" viewBox="0 0 29 33" :fill="getFill()">
            <path
                d="M28.509 4.2C28.3679 4.00623 28.1937 3.85081 27.9981 3.74422C27.8025 3.63763 27.59 3.58235 27.375 3.58209H6.6465L6.354 1.49731C6.29577 1.0791 6.11518 0.699114 5.84437 0.424943C5.57355 0.150771 5.23006 0.000195169 4.875 0H1.5C1.10218 0 0.720644 0.188699 0.43934 0.524585C0.158035 0.860471 0 1.31603 0 1.79104C0 2.26606 0.158035 2.72162 0.43934 3.0575C0.720644 3.39339 1.10218 3.58209 1.5 3.58209H3.6045L6.3945 23.5773L6.462 23.7994L6.543 24.0699L6.723 24.3904L6.8655 24.591L7.155 24.8239L7.323 24.9403C7.49745 25.0274 7.68443 25.073 7.8735 25.0746H24.375C24.7728 25.0746 25.1544 24.8859 25.4357 24.55C25.717 24.2142 25.875 23.7586 25.875 23.2836C25.875 22.8086 25.717 22.353 25.4357 22.0171C25.1544 21.6812 24.7728 21.4925 24.375 21.4925H9.1455L8.8965 19.7015H25.875C26.236 19.7016 26.5848 19.5462 26.8576 19.2639C27.1304 18.9817 27.3087 18.5914 27.36 18.1648L28.86 5.62746C28.8904 5.3734 28.8748 5.11451 28.8142 4.86826C28.7537 4.62201 28.6496 4.39413 28.509 4.2ZM25.6455 7.16418L25.218 10.7463H19.875V7.16418H25.6455ZM18.375 7.16418V10.7463H13.875V7.16418H18.375ZM18.375 12.5373V16.1194H13.875V12.5373H18.375ZM12.375 7.16418V10.7463H7.875L7.653 10.8L7.146 7.16418H12.375ZM7.896 12.5373H12.375V16.1194H8.3955L7.896 12.5373ZM19.875 16.1194V12.5373H25.002L24.5745 16.1194H19.875Z"
            />
            <path
                d="M10.125 32.2388C11.3676 32.2388 12.375 31.036 12.375 29.5522C12.375 28.0685 11.3676 26.8657 10.125 26.8657C8.88236 26.8657 7.875 28.0685 7.875 29.5522C7.875 31.036 8.88236 32.2388 10.125 32.2388Z"
            />
            <path
                d="M23.625 32.2388C24.8676 32.2388 25.875 31.036 25.875 29.5522C25.875 28.0685 24.8676 26.8657 23.625 26.8657C22.3824 26.8657 21.375 28.0685 21.375 29.5522C21.375 31.036 22.3824 32.2388 23.625 32.2388Z"
            />
          </svg>
        </NavLink>
        <PrimaryButton v-if="useAuth.idRef === -1" @click="router.push({name: 'SignUp'})" :class="'header-button'">Sign up</PrimaryButton>
        <PrimaryButton v-if="useAuth.idRef !== -1" @click="router.push({name: 'Profile'})" :class="'header-button'">Profile</PrimaryButton>
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav {
  padding-top: 20px;
  padding-bottom: 20px;
}

nav > .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar-content {
  display: flex;
  align-items: center;
}

ul {
  margin: 0 0;
  padding-left: 0;
  list-style: none;
  display: flex;
}

ul li {
  margin-right: 100px;
}

ul li:last-child {
  margin-right: 0;
}

svg {
  vertical-align: middle;
  margin-right: 50px;
}
</style>
