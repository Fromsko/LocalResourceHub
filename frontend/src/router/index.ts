import { createRouter, createWebHashHistory, type RouteRecordRaw } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'

const routes: Array<RouteRecordRaw> = [
  {
    name: 'Admin',
    path: '/admin',
    meta: {
      title: 'Admin',
      isPage: true
    },
    component: () => import('@components/Dashboard-finished.vue')
  },
  {
    name: 'Home',
    path: '/',
    meta: {
      title: 'Home',
      isPage: true,
      layout: 'default'
    },
    component: () => import('@/layouts/DefaultLayout.vue'),
    children: [
      {
        name: 'UploadResource',
        path: '',
        meta: {
          title: 'Upload Resource',
          isPage: true
        },
        component: () => import('@components/UploadResource.vue')
      },
      {
        name: 'FileList',
        path: 'files',
        meta: {
          title: 'File List',
          isPage: true
        },
        component: () => import('@components/FileList.vue')
      }
    ]
  }
]

const router = createRouter({
  routes: routes,
  history: createWebHashHistory()
})

router.beforeEach((to, from, next) => {
  if (to.meta.layout === 'default' && to.matched.length > 0) {
    const matchedRoute = to.matched[0]
    if (matchedRoute && matchedRoute.components) {
      matchedRoute.components = { default: matchedRoute.components.default, layout: DefaultLayout }
    }
  }
  next()
})

export default router
