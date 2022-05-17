<script setup>
import { useUserStore } from "../stores/counter.js"
const user = useUserStore()
const register = async () => {
  console.log("reg")
  const res = await fetch("http://127.0.0.1:8080/api/users", {
    method: "POST",
    body: JSON.stringify(user.$state),
  })
  console.log(res)
}
</script>
<template>
  <div>
    <h1>Register Now!</h1>
    <input v-model="user.username" placeholder="username" />
    <input v-model="user.password" placeholder="password" />
    <div @click="register">Register</div>
  </div>
</template>

<style></style>
