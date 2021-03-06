import { createRouter, createWebHistory } from "vue-router"
import HomeView from "../views/HomeView.vue"
import MallView from "../views/MallView.vue"
import MallsView from "../views/MallsView.vue"
import RegisterView from "../components/RegisterComponent.vue"
import LoginView from "../components/LoginComponent.vue"
import MallResultView from "../views/MallResultView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/malls",
      name: "malls",
      component: MallsView,
    },
    {
      path: "/malls/:id",
      name: "mall",
      component: MallView,
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/register",
      name: "register",
      component: RegisterView,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
    },
    {
      path: "/mall/new",
      name: "new mall",
      component: () => import("../components/NewMall.vue"),
    },
    {
      path: "/malls/view/:id",
      name: "view-mall",
      component: MallResultView,
    },
  ],
})

export default router
