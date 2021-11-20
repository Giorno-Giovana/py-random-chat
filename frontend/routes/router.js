import Vue from 'vue';
import Router from 'vue-router';
import Index from '../pages/index'
import Login from '../pages/login'
import Register from '../pages/register'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
  ]
});

export default router