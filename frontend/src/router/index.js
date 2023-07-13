import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";
import Suppliers from "../views/Suppliers.vue";

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
    }
]

export default VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})