import Vue from 'vue'
import VueRouter from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import AdminView from '@/views/AdminView.vue'
import ArticleView from '@/views/ArticleView.vue'

import EditView from '@/views/children/Admin/Edit/EditView.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', redirect: '/page/1' },
  { path: '/page/:id', component: HomeView },
  { path: '/article/:id', component: ArticleView },
  {
    path: '/admin', component: AdminView, children: [
      {
        path: 'edit', component: EditView
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
