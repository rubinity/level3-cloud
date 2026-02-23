import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateView from '../views/CreateView.vue'
import DeleteView from '../views/DeleteView.vue'
import VueView from '../views/VueView.vue'
import ListView from '../views/ListView.vue'
import ConnectionView from '../views/ConnectionView.vue'
import AuthView from '../views/AuthView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/create',
    name: 'create',
    component: CreateView
  },
  {
    path: '/delete',
    name: 'delete',
    component: DeleteView
  },
    {
    path: '/connection',
    name: 'connection',
    component: ConnectionView
  },
    {
    path: '/list',
    name: 'list',
    component: ListView
  },
    {
      // info about view
    path: '/vue',
    name: 'vue',
    component: VueView
  },
      {
      // info about view
    path: '/auth',
    name: 'AuthView',
    component: AuthView
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
