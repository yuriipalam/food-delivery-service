import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";
import Suppliers from "../views/Suppliers.vue";
import Categories from "../views/Categories.vue";
import Category from "../views/Category.vue"
import Supplier from "../views/Supplier.vue";
import SignUp from "../views/SignUp.vue";
import SignIn from "../views/SignIn.vue";
import Profile from "../views/Profile.vue";
import {useAuthStore} from "../store";
import Cart from "../views/Cart.vue";
import NotFound404 from "../views/NotFound404.vue";
import SomethingWentWrong from "../views/SomethingWentWrong500.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        meta: {
            title: 'FoodJet - Home'
        },
        component: Home
    },
    {
        path: '/suppliers',
        name: 'Suppliers',
        meta: {
            title: 'FoodJet - Suppliers Flow'
        },
        component: Suppliers
    },
    {
        path: '/suppliers/:id',
        name: 'Supplier',
        meta: {
            title: 'FoodJet - Explore Supplier'
        },
        component: Supplier
    },
    {
        path: '/categories',
        name: 'Categories',
        meta: {
            title: 'FoodJet - Categories'
        },
        component: Categories
    },
    {
        path: '/categories/:id',
        name: 'Category',
        meta: {
            title: 'FoodJet - Explore Category'
        },
        component: Category
    },
    {
        path: '/sign-up',
        name: 'SignUp',
        meta: {
            title: 'FoodJet - Sign Up'
        },
        component: SignUp
    },
    {
        path: '/sign-in',
        name: 'SignIn',
        meta: {
            title: 'FoodJet - Sign In'
        },
        component: SignIn
    },
    {
        path: '/profile',
        name: 'Profile',
        meta: {
            requiresAuth: true,
            title: 'FoodJet - Profile'
        },
        component: Profile
    },
    {
        path: '/profile/orders',
        name: 'ProfileOrders',
        meta: {
            requiresAuth: true,
            title: 'FoodJet - Orders'
        },
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
        meta: {
            title: 'FoodJet - Cart'
        },
        component: Cart
    },
    {
        path: '/not-found',
        name: '404',
        meta: {
            title: 'FoodJet - Not Found'
        },
        component: NotFound404
    },
    {
        path: '/something-went-wrong',
        name: '500',
        meta: {
            title: 'FoodJet - Something Went Wrong'
        },
        component: SomethingWentWrong
    },
    // default redirect for vue router
    {
        path: '/:pathMatch(.*)*',
        component: NotFound404
    }
]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})

export default router

// check if user is authenticated only when want to know this
router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth && useAuthStore().idRef === -1) {
        next({name: 'SignIn'})
    } else {
        next()
    }
})

// takes title from meta and sets it for each route
router.beforeEach((to, from, next) => {
    const title = to.meta.title
    if (title) {
        document.title = title
    }
    next()
})
