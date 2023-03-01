import { createRouter, createWebHistory } from 'vue-router'

import MainPage from './views/MainPage.vue'
import AboutView from './views/AboutView.vue'
import FeedbackView from './views/MainPage.vue'

const routes = [
    {
        path: '/',
        name: 'index',
        component: MainPage
    },
    {
        path: '/about',
        name: 'about',
        component: AboutView
    },
    {
        path: '/feedback',
        name: 'feedback',
        component: FeedbackView
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router