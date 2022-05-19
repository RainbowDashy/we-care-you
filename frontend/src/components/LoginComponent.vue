<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import { useUserStore } from "../stores/user.js"
const user = useUserStore()
const router = useRouter()

let input = reactive({
  username: "",
  password: "",
})

const login = async () => {
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
    <h1>Login Now!</h1>
    <input v-model="input.username" placeholder="username" />
    <input v-model="input.password" placeholder="password" type="password" />
    <div @click="login">login</div>
  </div>
</template>

<style></style>
