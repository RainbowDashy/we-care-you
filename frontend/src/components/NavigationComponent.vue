<script setup>
import { NMenu } from "naive-ui"
import { h } from "vue"
import { RouterLink } from "vue-router"
import { useUserStore } from "../stores/user"
const user = useUserStore()
const menuOptions = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "home",
          },
        },
        { default: () => "主页" }
      ),
    key: "router-home",
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "malls",
          },
        },
        { default: () => "商品团购" }
      ),
    key: "router-malls",
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "new mall",
          },
        },
        { default: () => "发起团购" }
      ),
    key: "router-mall-new",
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "register",
          },
        },
        { default: () => "注册" }
      ),
    key: "router-register",
    disabled: user.logined,
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "login",
          },
        },
        { default: () => "登录" }
      ),
    key: "router-login",
    disabled: user.logined,
  },
]
</script>
<template>
  <n-menu :options="menuOptions" />
</template>
