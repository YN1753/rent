import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import CityList from '../views/CityList.vue'
import LoginAccount from '../views/LoginAccount.vue'
import LoginEmail from '../views/LoginEmail.vue'
import User from '../views/User.vue'
import Register from '../views/Register.vue'
import ForgotPassword from '../views/ForgotPassword.vue'

const routes = [
  { path: '/', redirect: '/home' },
  { path: '/home', component: Home },
  { path: '/city', component: CityList },
  { path: '/login', redirect: '/login/account' },
  { path: '/login/account', component: LoginAccount },
  { path: '/login/email', component: LoginEmail },
  { path: '/user', component: User },
  { path: '/register', component: Register },
  { path: '/forget', component: ForgotPassword }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
