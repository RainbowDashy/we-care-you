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
        { default: () => "Home" }
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
        { default: () => "Malls" }
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
        { default: () => "New Mall" }
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
        { default: () => "Register" }
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
        { default: () => "Login" }
      ),
    key: "router-login",
    disabled: user.logined,
  },
]
</script>
<template>
  <n-menu :options="menuOptions" />
</template>
