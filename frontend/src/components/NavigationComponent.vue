<script setup>
import { NMenu } from "naive-ui"
import { NIcon, useMessage } from "naive-ui";
import { h } from "vue"
import { RouterLink } from "vue-router"
import { useUserStore } from "../stores/user"
import {
  ListOutline as ListIcon,
  HomeOutline as HomeIcon
} from "@vicons/ionicons5";

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
    icon: renderIcon(HomeIcon)
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
     icon: renderIcon(ListIcon)
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
    icon: renderIcon(ListIcon)
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
    icon: renderIcon(ListIcon)
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
    icon: renderIcon(ListIcon)
  },
]
function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}
</script>
<template>
  <n-menu :options="menuOptions" />
</template>
