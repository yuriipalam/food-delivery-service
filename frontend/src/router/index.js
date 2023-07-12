import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    }
]

export default VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})