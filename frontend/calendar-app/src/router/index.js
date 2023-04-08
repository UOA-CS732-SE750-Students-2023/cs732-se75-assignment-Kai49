import { createRouter, createWebHistory } from 'vue-router'
import Calendar from '../components/Calendar'

const routes = [
  {
    path: '/',
    name: 'Calendar',
    component: Calendar
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
