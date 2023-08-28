import Home from '../views/home.vue';
import { createRouter,createWebHashHistory } from 'vue-router'
import Personal from "../views/personal.vue";
import Test1 from "../views/test1.vue";
import Test2 from "../views/test2.vue";
import Loading from "../views/loading.vue";

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
                name: 'test1',
                path: 'test1',
                component: Test1
            },
            {
                name: 'test2',
                path: 'test2',
                component: Test2
            }
        ],
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default { router, routes };