<script setup>
import { useUserStore } from "../stores/user.js"
const user = useUserStore()

const login = async () => {
  try {
    await user.login()
  } catch (err) {
    console.log(err)
  }
}
</script>
<template>
  <div>
    <h1>Login Now!</h1>
    <p v-if="user.logined">logined</p>
    <input v-model="user.username" placeholder="username" />
    <input v-model="user.password" placeholder="password" />
    <div @click="login">login</div>
  </div>
</template>

<style></style>
