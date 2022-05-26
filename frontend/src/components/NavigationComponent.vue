<script setup>
import { NMenu } from "naive-ui"
import { NIcon } from "naive-ui"
import { computed, h } from "vue"
import { RouterLink } from "vue-router"
import { useUserStore } from "../stores/user"
import {
  HomeOutline,
  StorefrontOutline,
  CartOutline,
  PersonOutline,
  PersonAddOutline,
} from "@vicons/ionicons5"

const user = useUserStore()
const menuOptions = computed(() => [
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
    icon: renderIcon(HomeOutline),
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
        { default: () => "团购" }
      ),
    key: "router-malls",
    icon: renderIcon(CartOutline),
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
        { default: () => "开团" }
      ),
    key: "router-mall-new",
    icon: renderIcon(StorefrontOutline),
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
    icon: renderIcon(PersonAddOutline),
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
    icon: renderIcon(PersonOutline),
  },
])
function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
}
</script>
<template>
  <n-menu class="navigation" :options="menuOptions" mode="horizontal" />
</template>

<style>
.navigation .n-menu-item-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 !important;
}
.navigation .n-menu-item-content__icon {
  margin-right: 0 !important;
}
.navigation {
  display: flex !important;
  flex-direction: row;
  justify-content: space-around;
}
</style>
