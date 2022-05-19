<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import http from "../http.js"
import { useUserStore } from "../stores/user.js"

const router = useRouter()
const user = useUserStore()
const input = reactive({
  username: "",
  password: "",
  location: "",
})
const registerAndLogin = async () => {
  try {
    await http.post("/users", {
      body: JSON.stringify(input),
    })
  } catch (err) {
    console.log(err)
    return
  }

  try {
    await user.login(input.username, input.password)
  } catch (err) {
    console.log(err)
    return
  }

  router.push("/")
}
</script>
<template>
  <div>
    <h1>Register Now!</h1>
    <form>
      <input
        v-model="input.username"
        placeholder="username"
        autocomplete="username "
      />
      <input
        v-model="input.password"
        placeholder="password"
        type="password"
        autocomplete="password"
      />
      <input v-model="input.location" placeholder="location" />
      <div @click="registerAndLogin">Register</div>
    </form>
  </div>
</template>

<style></style>
