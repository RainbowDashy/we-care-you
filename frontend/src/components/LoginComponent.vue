<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import { useUserStore } from "../stores/user.js"
import { NButton, NForm, NFormItem, NInput } from "naive-ui"
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
    <n-form :label-width="80">
      <n-form-item>
        <n-input
          v-model:value="input.username"
          placeholder="username"
        ></n-input>
      </n-form-item>
      <n-form-item>
        <n-input
          v-model:value="input.password"
          placeholder="password"
          type="password"
        ></n-input>
      </n-form-item>
      <n-form-item>
        <n-button attr-type="button" @click="login">Login</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style></style>
