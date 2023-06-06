import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginPanel from "../views/LoginPanel.vue"
import RegisterPage from "../views/RegisterPage.vue"
import MainPage from "../views/MainPage.vue"
import Personal from "../views/Personal.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: "login",
    component: LoginPanel,
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterPage
  },
  {
    path: '/index',
    name: 'index',
    component: MainPage
  },
  {
    path: '/personal',
    name: 'personal',
    component: Personal
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.path === from.path) {
    console.log(666)
    return false;
  }
  const token = window.localStorage.getItem('token')
  if (to.path == '/login') next()
  if ( !token ) {
    next('/login')
  } else {
    next()
  }
})
export default router
