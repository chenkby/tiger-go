import { createRouter, createWebHashHistory } from 'vue-router'

const Home = { template: '<div>Home</div>' }
const About = { template: '<div>About</div>' }

const routes = [
  { path: '/', component: () => import('@/components/HelloWorld.vue') },
  { path: '/about', component: About },
  {
    path: '/article',
    component: () => import('@/modules/article/views/index.vue'),
  },
  {
    path: '/article/create',
    component: () => import('@/modules/article/views/form.vue'),
  },
  {
    path: '/article/view',
    component: () => import('@/modules/article/views/view.vue'),
  },
  // 404{ path: '/:pathMatch(.*)', component: NotFoundComponent },
]

const router = createRouter({
  // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHashHistory(),
  routes, // `routes: routes` 的缩写
})

export default router
