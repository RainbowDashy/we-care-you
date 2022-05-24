<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import http from "../http.js"
import { useUserStore } from "../stores/user.js"
import { NButton, NForm, NFormItem, NInput,NIcon } from "naive-ui"
import { useMessage } from "naive-ui"
import {
  BagCheckOutline,
  PersonAddOutline,
  LocationOutline,
} from "@vicons/ionicons5"
const router = useRouter()
const user = useUserStore()

const input = reactive({
  username: "",
  password: "",
  location: "",
})
const message=useMessage()
const registerAndLogin = async () => {
  try {
    await http.post("/users", {
      body: JSON.stringify(input),
    })
    message.success("注册成功")
  } catch (err) {
    console.log(err)
    message.error("注册失败")
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
  <div >
    <h1>注册</h1>
    <h2>创建新的账户</h2>
    <n-form>
      <n-form-item>
        <n-input v-model:value="input.username" placeholder="用户名" >
         <template #prefix>
        <n-icon :component="PersonAddOutline" />
         </template>
        </n-input>
      </n-form-item>
      <n-form-item>
        <n-input
          v-model:value="input.password"
          placeholder="密码"
          type="password"
        >
        <template #prefix>
        <n-icon :component="BagCheckOutline" />
      </template>
      </n-input>
      </n-form-item>
      <n-form-item>
        <n-input v-model:value="input.location" placeholder="地址" >
        <template #prefix>
        <n-icon :component="LocationOutline" />
         </template>
         </n-input>
      </n-form-item>
      <div style="display:flex;justify-content: center">
      <n-form-item>
        <n-button
          attr-type="button"
          color="#426E52"
          size="large"
          @click="registerAndLogin"
        >
          注册
        </n-button>
      </n-form-item>
      </div>
    </n-form>
  </div>
</template>

<style scoped>
h1 {
  color: #227700;
  text-align: center;
}
h2 {
  color: #8b8b8b;
  text-align: center;
}
 
</style>
