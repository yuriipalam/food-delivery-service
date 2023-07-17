import * as VueRouter from "vue-router"
import Home from "../views/Home.vue";
import Suppliers from "../views/Suppliers.vue";
import Categories from "../views/Categories.vue";
import ContactUs from "../views/ContactUs.vue";
import Placeholder from "../components/Placeholder.vue";
import Supplier from "../views/Supplier.vue";

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
        path: '/contact-us',
        name: 'ContactUs',
        component: ContactUs
    },
    {
        path: '/placeholder',
        name: 'Placeholder',
        component: Placeholder
    }
]

export default VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes
})
