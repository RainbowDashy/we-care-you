<script setup>
import {
  NButton,
  NDatePicker,
  NDynamicInput,
  NInput,
  NInputNumber,
  zhCN,
  dateZhCN,
  NConfigProvider,
  useMessage,
  NSwitch,
} from "naive-ui"

import { useRouter } from "vue-router"
import { ref } from "vue"
import http from "../http"
const router = useRouter()
const customValue = ref([])
const endtime = ref()
const onCreate = () => {
  return {
    name: "",
    total: undefined,
    price: undefined,
    description: "",
    data: "",
  }
}
const message = useMessage()
const submit = async () => {
  const payload = {
    begintime: Date.now(),
    endtime: endtime.value,
    state: 1,
    items: customValue.value.map((v) => ({
      ...v,
      price: Math.trunc(v.price * 100),
    })),
  }
  try {
    await http.post("/malls", {
      body: JSON.stringify(payload),
    })
    message.success("团购创建成功")
  } catch (err) {
    console.log(err)
    message.error("团购创建失败")
    return
  }
  router.push("/malls")
}
</script>

<template>
  <main>
    <n-form-item>
      <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
        <n-date-picker v-model:value="endtime" type="datetime" />
      </n-config-provider>
      <br />
      <n-space>
        <n-switch size="medium">
          <template #unchecked> 是否可以预约 </template>
          <template #checked> 团购可预约 </template>
        </n-switch>
      </n-space>
      <br />
      <br />
      <n-dynamic-input v-model:value="customValue" :on-create="onCreate">
        <template #create-button-default>创建新的团购 </template>
        <template #default="{ value }">
          <div>
            <n-input placeholder="商品名称" v-model:value="value.name" />
            <n-input-number
              placeholder="商品总量"
              v-model:value="value.total"
            />
            <n-input-number placeholder="商品价格" v-model:value="value.price">
              <template #prefix>￥</template>
            </n-input-number>
            <n-input placeholder="商品描述" v-model:value="value.description" />
          </div>
        </template>
      </n-dynamic-input>
    </n-form-item>
    <p></p>
    <br />
    <div style="display: flex; justify-content: center">
      <n-form-item>
        <n-button color="#227700" @click="submit"> 提交 </n-button>
      </n-form-item>
    </div>
  </main>
</template>
