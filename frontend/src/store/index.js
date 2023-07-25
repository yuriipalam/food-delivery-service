import {defineStore} from 'pinia'
import {ref} from "vue";

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
    const phoneRef = ref('')
    const firstNameRef = ref('')
    const lastNameRef = ref('')

    async function setTokens(accessToken, refreshToken) {
        accessTokenRef.value = accessToken
        refreshTokenRef.value = refreshToken
    }

    async function setUser(id, email, phone, firstName, lastName) {
        idRef.value = id
        emailRef.value = email
        phoneRef.value = phone
        firstNameRef.value = firstName
        lastNameRef.value = lastName
    }

    async function signOut() {
        await setUser(-1, "", "", "", "")
        await setTokens('', '')
    }

    return {
        setTokens,
        setUser,
        signOut,
        accessTokenRef,
        refreshTokenRef,
        idRef,
        emailRef,
        phoneRef,
        firstNameRef,
        lastNameRef,
    }
}, {
    persist: true
})

export const useCartStore = defineStore('cart', () => {
    const products = ref({})
    const supplierIDs = ref([])
    const err = ref('')

    function addProduct(product) {
        const id = product.id

        if (products.value[id]) {
            if (products.value[id].quantity < 8) {
                products.value[id].quantity++
            }
            return
        } else if (!supplierIDs.value.includes(product.supplier_id)) {
            if (supplierIDs.value.length >= 2) {
                err.value = 'You can order from at most two suppliers'
                return
            }
            supplierIDs.value.push(product.supplier_id)
        }

        products.value[id] = {
            product: product, quantity: 1
        }
    }

    function reduceQuantity(id) {
        if (products.value[id].quantity <= 1) {
            const currProduct = products.value[id].product
            // clearing supplier list
            if (Object.keys(products.value).filter(iterID => products.value[iterID].product.supplier_id === currProduct.supplier_id).length === 1) {
                supplierIDs.value = supplierIDs.value.filter(iterID => iterID !== currProduct.supplier_id)
            }
            delete products.value[id]
            // if (products.value.length === 0) {
            //     supplierIDs.value = []
            // }
        } else {
            products.value[id].quantity--
        }
    }

    function increaseQuantity(id) {
        if (products.value[id].quantity >= 8) {
            return
        }
        products.value[id].quantity++
    }

    function getQuantity(id) {
        return products.value[id].quantity
    }

    function getProductTotalPrice(id) {
        return Math.round(getQuantity(id) * products.value[id].product.price * 100) / 100
    }

    function getTotalPrice() {
        return Object.keys(products.value).map(key => getProductTotalPrice(parseInt(key))).reduce((total, current) => total + current, 0).toFixed(2)
    }

    function clearCart() {
        products.value = {}
        supplierIDs.value = []
    }

    return {
        addProduct,
        reduceQuantity,
        increaseQuantity,
        getQuantity,
        getProductTotalPrice,
        clearCart,
        getTotalPrice,
        err,
        supplierIDs,
        products
    }
}, {
    persist: true
})
