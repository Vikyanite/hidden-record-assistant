import Home from '../views/home.vue';
import { createRouter,createWebHashHistory } from 'vue-router'
import Personal from "../views/personal.vue";
import Test1 from "../views/test1.vue";
import Test2 from "../views/test2.vue";
import Loading from "../views/loading.vue";
import search from "../views/search.vue";

const routes = [
    {
        path: '/loading',
        component: Loading,
    },
    {
        path: '/home',
        component:Home,
        children: [
            {
                name: 'personal',
                path: 'personal',
                component: Personal
            },
            {
                name: 'search',
                path: 'search',
                component: search,
            },
        ],
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default { router, routes };