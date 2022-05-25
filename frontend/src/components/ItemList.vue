<script setup>
import { NImage, NButton, NSpace, NInputNumber, useMessage } from "naive-ui"
import { computed, onMounted, ref } from "vue"
import http from "../http"
import { useUserStore } from "../stores/user"

const props = defineProps({
  mallId: {
    type: Number,
    required: true,
  },
})
const message = useMessage()
const items = ref([])
const orders = ref([])
const user = useUserStore()

const totalPrice = computed(() => {
  return (
    orders.value.reduce(
      (total, order) => total + order.price * order.buycount,
      0
    ) / 100
  )
})

const fetchItems = async () => {
  try {
    const res = await http.get(`/items?mallid=${props.mallId}`)
    items.value = await res.json()
    orders.value = items.value.map((item) => ({
      mallid: props.mallId,
      userid: user.id,
      itemid: item.id,
      buycount: 0,
      price: item.price,
    }))
  } catch (err) {
    console.log(err)
  }
}

onMounted(fetchItems)

const buy = async () => {
  try {
    await http.post("/orders", {
      body: JSON.stringify({ orders: orders.value }),
    })
    message.success("下单成功")
  } catch (err) {
    console.log(err)
    message.error("下单失败，请重新下单")
  }
}
</script>

<template>
  <div>
    <template v-for="(item, index) in items" :key="item.id">
      <div class="item-container">
        <div class="item-image">
          <n-image
            src="https://imgservice.suning.cn/uimg1/b2c/image/5NBrypvkDaxG0-o-qafX9A.jpg"
          />
        </div>
        <div class="item-description">
          <h3>{{ item.name }}</h3>
          <p>{{ item.description }}</p>
          <p>¥{{ item.price / 100 }}</p>
        </div>
        <div class="item-action">
          <n-input-number
            placeholder="购买数量"
            v-model:value="orders[index].buycount"
            :min="0"
            :max="item.total"
          />
        </div>
      </div>
    </template>
  </div>
  <div>
    <n-space justify="space-between" align="center">
      <div>总价：¥{{ totalPrice }}</div>
      <n-button @click="buy">下单</n-button>
    </n-space>
  </div>
</template>

<style>
.item-container {
  display: flex;
  align-items: center;
  height: 5rem;
}
.item-description {
  flex-grow: 1;
}
.item-action {
  width: 30vw;
}
.item-image,
.item-image img {
  width: 5rem;
  height: 5rem;
}
.item-description p,
.item-description h3 {
  margin: 0;
}
</style>
