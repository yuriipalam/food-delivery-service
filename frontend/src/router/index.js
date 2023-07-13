import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";
import Suppliers from "../views/Suppliers.vue";
import Categories from "../views/Categories.vue";
import ContactUs from "../views/ContactUs.vue";

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
        path: '/categories',
        name: 'Categories',
        component: Categories
    },
    {
        path: '/contact-us',
        name: 'ContactUs',
        component: ContactUs
    }
]

export default VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})