import { createRouter, createWebHashHistory } from 'vue-router'

const Home = { template: '<div>Homeffffffgdfgf2342343</div>' }
const routes = [{ path: '/login', component: Home }]
const router = createRouter({
  history: createWebHashHistory(),
  routes,
})
export default router
