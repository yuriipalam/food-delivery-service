import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";
import Suppliers from "../views/Suppliers.vue";
import Categories from "../views/Categories.vue";
import Category from "../views/Category.vue"
import ContactUs from "../views/ContactUs.vue";
import Placeholder from "../components/Placeholder.vue";
import Supplier from "../views/Supplier.vue";
import SignUp from "../views/SignUp.vue";
import SignIn from "../views/SignIn.vue";
import Profile from "../views/Profile.vue";
const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/suppliers',
        name: 'Suppliers',
        component: Suppliers
    },
    {
        path: '/suppliers/:id',
        name: 'Supplier',
        component: Supplier
    },
    {
        path: '/categories',
        name: 'Categories',
        component: Categories
    },
    {
        path: '/categories/:id',
        name: 'Category',
        component: Category
    },
    {
        path: '/contact-us',
        name: 'ContactUs',
        component: ContactUs
    },
    {
        path: '/sign-up',
        name: 'SignUp',
        component: SignUp
    },
    {
        path: '/sign-in',
        name: 'SignIn',
        component: SignIn
    },
    {
        path: '/profile',
        name: 'Profile',
        meta: {requiresAuth: true},
        component: Profile
    },
    {
        path: '/profile/orders',
        name: 'ProfileOrders',
        meta: {requiresAuth: true},
        component: Profile
    },
    // {
    //     path: '/profile/settings',
    //     name: 'ProfileSettings',
    //     meta: {requiresAuth: true},
    //     component: Profile
    // },
    {
        path: '/cart',
        name: 'Cart',
        component: Cart
    }
]

import {useAuthStore} from "../store";
import OrderField from "../components/Supplier/OrderField.vue";
import Cart from "../views/Cart.vue";

const router = VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})

export default router

router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth && useAuthStore().idRef === -1) {
        next({name: 'SignIn'})
    } else {
        next()
    }
})
