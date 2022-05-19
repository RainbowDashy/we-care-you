<script setup>
import { RouterLink, RouterView } from "vue-router"
import { useUserStore } from "./stores/user"
import { onMounted } from "vue"

const user = useUserStore()

onMounted(() => {
  user.fetch()
})
</script>

<template>
  <header>
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
      <RouterLink to="/register" v-if="!user.logined">Register</RouterLink>
      <RouterLink to="/login" v-if="!user.logined">Login</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>
