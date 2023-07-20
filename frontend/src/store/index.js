import {defineStore} from 'pinia'
import {reactive, ref} from "vue";

export const useFiltersStore = defineStore("filter", () => {
    const selectedCategory = ref(0)
    const selectedCategoryName = ref('')
    const searchFor = ref('')

    async function selectCategory(id, name) {
        selectedCategory.value = id
        selectedCategoryName.value = name
    }

    async function reset() {
        selectedCategory.value = 0
        selectedCategoryName.value = ''
        searchFor.value = ''
    }

    return {selectedCategory, searchFor, selectCategory, selectedCategoryName, reset}
})

export const useAuthStore = defineStore("auth", () => {
    const accessTokenRef = ref('')
    const refreshTokenRef = ref('')

    const idRef = ref(-1)
    const emailRef = ref('')
    const firstNameRef = ref('')
    const lastNameRef = ref('')

    async function setTokens(accessToken, refreshToken) {
        accessTokenRef.value = accessToken
        refreshTokenRef.value = refreshToken
    }

    async function setUser(id, email, firstName, lastName) {
        idRef.value = id
        emailRef.value = email
        firstNameRef.value = firstName
        lastNameRef.value = lastName
    }

    async function signOut() {
        await setUser(-1, "", "", "")
    }

    return {
        setTokens, setUser, signOut, accessTokenRef, refreshTokenRef, idRef, emailRef, firstNameRef, lastNameRef,
    }
}, {
    persist: true
})

export const useCartStore = defineStore('cart', () => {
    const products = ref([])

    function addProduct(product) {
        let index = products.value.findIndex(prod => prod.prod === product)
        if (index !== -1) {
            products.value[index].quantity++
        } else {
            products.value.push({
                prod: product, quantity: 1,
            })
        }
    }

    function reduceQuantity(product) {
        let index = products.value.findIndex(prod => prod.prod === product)
        products.value[index].quantity--
        if (products.value[index].quantity < 1) {
            products.value.splice(index, 1)
        }
    }

    function increaseQuantity(product) {
        let index = products.value.findIndex(prod => prod.prod === product)
        if (products.value[index].quantity >= 5) {
            return
        }
        products.value[index].quantity++
    }

    function getQuantity(product) {
        return products.value[products.value.findIndex(prod => prod.prod === product)].quantity
    }

    return {addProduct, reduceQuantity, increaseQuantity, getQuantity, products}
}, {
    persist: true
})
