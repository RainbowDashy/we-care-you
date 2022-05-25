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
  <div id='building'>
    <h1>注册</h1>
    <p>创建新的账户</p>
    <n-form>
      <n-form-item>
        <n-input v-model:value="input.username" placeholder="用户名" />
      </n-form-item>
      <n-form-item>
        <n-input
          v-model:value="input.password"
          placeholder="密码"
          type="password"
        />
      </n-form-item>
      <n-form-item>
        <n-input v-model:value="input.location" placeholder="地址" />
      </n-form-item>
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
    </n-form>
  </div>
  
</template>


<style scoped>
h1 {
  color: #007702;
}
p {
  color: #120762;
  font-size: 10px;
  text-indent: 2em;
  line-height: 15px;
}


#building{
  background:url("../photos/3.jpg");
  width:100%;
  height:100%;
  position:fixed;
  background-size:100% 30%;
  background-repeat: no-repeat;
}



</style> 

