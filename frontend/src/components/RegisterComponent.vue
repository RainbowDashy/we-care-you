<script setup>
import http from "../http.js";
import { useUserStore } from "../stores/user.js"
const user = useUserStore()
const register = async () => {
  try {
    const res = await http.post("/users", {
      body: JSON.stringify(user.$state),
    });
    console.log(await res.json())
  } catch(err) {
    console.log(err);
  }
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
