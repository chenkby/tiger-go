import {
  createRouter,
  createWebHashHistory,
  createWebHistory
} from 'vue-router'

const Home = { template: '<div>Home</div>' }
const About = { template: '<div>About</div>' }

const routes = [
  {
    path: '/',
    component: () => import('@/views/layouts/layout.vue'),
    children: [
      {
        path: '/article',
        component: () => import('@/modules/article/views/index.vue')
      },
      {
        path: '/article/create',
        component: () => import('@/modules/article/views/form.vue')
      },
      {
        path: '/article/update',
        component: () => import('@/modules/article/views/form.vue')
      },
      {
        path: '/article/view',
        component: () => import('@/modules/article/views/view.vue')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login.vue')
  },
  { path: '/:pathMatch(.*)', component: import('@/views/404.vue') }
  // 404{ path: '/:pathMatch(.*)', component: NotFoundComponent },
]

const router = createRouter({
  // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
  history: createWebHistory(),
  routes, // `routes: routes` 的缩写
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      console.log('savePosition', savedPosition)
      return savedPosition
    }
    return { left: 0, top: 0 }
  }
})

export default router
