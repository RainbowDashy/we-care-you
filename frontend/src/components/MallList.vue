<script setup>
import MallListItem from "./MallListItem.vue"
import { computed, onMounted, ref } from "vue"
import http from "../http"
import { NSpace, NSwitch } from "naive-ui"
import { useUserStore } from "../stores/user"
defineProps({
  search: {
    type: String,
    default: "",
  },
})

const founderFilter = ref(false)
const customerFilter = ref(false)
const user = useUserStore()

const malls = ref([])
const userOrder = ref(new Set())
const filteredMalls = computed(() => {
  let res = malls.value.filter((mall) => {
    return mall.state == 1
  })
  if (founderFilter.value) {
    res = res.filter((mall) => {
      return mall.userid == user.id
    })
  }
  if (customerFilter.value) {
    res = res.filter((mall) => {
      return userOrder.value.has(mall.id)
    })
  }
  return res
})

const refresh = async () => {
  try {
    const res = await http.get("/malls")
    const data = await res.json()
    malls.value = data
  } catch (err) {
    console.log(err)
  }
}

const fetchOrders = async () => {
  const res = await http.get(`/orders?userid=${user.id}`)
  const orders = await res.json()
  for (let order of orders) {
    userOrder.value.add(order.mallid)
  }
}

onMounted(async () => {
  await refresh()
  await fetchOrders()
})
</script>

<template>
  <!-- <p>The search string is {{ search }}</p> -->
  <n-space>
    <n-switch v-model:value="customerFilter">
      <template #checked> 我参加的 </template>
      <template #unchecked> 我参加的 </template>
    </n-switch>
    <n-switch v-model:value="founderFilter">
      <template #checked> 我发起的 </template>
      <template #unchecked> 我发起的 </template>
    </n-switch>
  </n-space>
  <MallListItem
    v-for="mall in filteredMalls"
    :key="mall.id"
    :mall="mall"
  ></MallListItem>
</template>
