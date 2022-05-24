<script setup>
import {
  NButton,
  NDatePicker,
  NDynamicInput,
  NInput,
  NInputNumber,
} from "naive-ui"
import { ref } from "vue"
import http from "../http"
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
  } catch (err) {
    console.log(err)
    return
  }
}
</script>

<template>
  <main>
    <n-date-picker v-model:value="endtime" type="datetime" />
    <n-dynamic-input v-model:value="customValue" :on-create="onCreate">
      <template #create-button-default>创建新的团购 </template>
      <template #default="{ value }">
        <div>
          <n-input placeholder="商品名称" v-model:value="value.name" />
          <n-input-number placeholder="商品总量" v-model:value="value.total" />
          <n-input-number placeholder="商品价格" v-model:value="value.price">
            <template #prefix>￥</template>
          </n-input-number>
          <n-input placeholder="商品描述" v-model:value="value.description" />
        </div>
      </template>
    </n-dynamic-input>
    <n-button color="#227700" @click="submit"> 提交 </n-button>
  </main>
</template>
