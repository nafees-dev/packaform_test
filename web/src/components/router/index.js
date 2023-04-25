import { createWebHistory, createRouter } from "vue-router";
import HelloWorld  from '../views/HelloWorld'
import Order  from '../views/Order'

const routes = [
    {
        path: '/',
        name: 'home',
        component: HelloWorld,
    },
    {
        path: '/order',
        name: 'home',
        component: Order,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
  export default router;