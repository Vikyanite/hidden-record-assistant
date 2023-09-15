
import { createRouter,createWebHashHistory } from 'vue-router'

import Home from '../views/home.vue';

const routes = [
    {
        path: '/loading',
        component: () => import('../views/loading.vue'),
    },
    {
        path: '/home',
        component:Home,
        children: [
            {
                name: 'search',
                path: 'search',
                component: () => import('../views/search.vue'),
            },
            {
                name: 'result',
                path: 'result/:name',
                component: () => import('../views/result.vue'),
            },
            {
                name: 'personal',
                path: 'personal',
                component: () => import('../views/personal.vue'),
            }
        ],
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default { router, routes };