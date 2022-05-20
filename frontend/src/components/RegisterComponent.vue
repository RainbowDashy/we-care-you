<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import http from "../http.js"
import { useUserStore } from "../stores/user.js"
import { NButton, NForm, NFormItem, NInput } from "naive-ui"
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
    <n-form>
      <n-form-item>
        <n-input v-model:value="input.username" placeholder="username" />
      </n-form-item>
      <n-form-item>
        <n-input
          v-model:value="input.password"
          placeholder="password"
          type="password"
        />
      </n-form-item>
      <n-form-item>
        <n-input v-model:value="input.location" placeholder="location" />
      </n-form-item>
      <n-form-item>
        <n-button attr-type="button" @click="registerAndLogin">
          Register
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style></style>
