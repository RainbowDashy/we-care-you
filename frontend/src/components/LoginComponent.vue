<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import { useUserStore } from "../stores/user.js"
import { NButton, NForm, NFormItem, NInput, useMessage, NIcon } from "naive-ui"
import { BagCheckOutline, PersonOutline } from "@vicons/ionicons5"

const user = useUserStore()
const router = useRouter()

let input = reactive({
  username: "",
  password: "",
})

const message = useMessage()

const login = async () => {
  try {
    await user.login(input.username, input.password)
    message.success("登录成功")
  } catch (err) {
    console.log(err)
    message.error("登录失败，请输入正确的用户名密码")
    return
  }
  router.push("/")
}
</script>
<template>
<div id="building"> 
    <h1>欢迎使用</h1>
    <h2>账户登录</h2>
    <n-form :label-width="80">
      <n-form-item>
        <n-input v-model:value="input.username" placeholder="用户名">
          <template #prefix>
            <n-icon :component="PersonOutline" />
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
      <div style="display: flex; justify-content: center">
        <n-form-item>
          <n-button
            attr-type="button"
            color="#426E52"
            size="large"
            @click="login"
          >
            登录
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
#building{
  background:url("../photos/3.jpg") scroll top rgba(255, 255, 255, 0.5);
  width:100%;
  height:10%;
/*   background-position:0px 100px; */
  background-size:100% 100%;
  background-repeat: no-repeat;
}
</style>
