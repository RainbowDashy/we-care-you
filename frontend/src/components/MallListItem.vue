<script setup>
import { computed, onMounted, ref } from "vue"
import http from "../http"
import { NCard, NTime, NSpace, NButton } from "naive-ui"

const props = defineProps({
  mall: {
    require: true,
  },
})

const items = ref([])

const mallDescription = computed(() => {
  return items.value.map((item) => item.name).join(" ")
})

const fetchItems = async () => {
  try {
    const res = await http.get(`/items?mallid=${props.mall.id}`)
    const data = await res.json()
    items.value = data
  } catch (err) {
    console.log(err)
  }
}

onMounted(fetchItems)
</script>

<template>
  <div class="mall">
    <n-card :title="`拼团${mall.id}`" size="small">
      <n-space justify="space-between">
        <p>开始：<n-time :time="mall.begintime" format="M月d日H时mm分" /></p>
        <p>结束：<n-time :time="mall.endtime" format="M月d日H时mm分" /></p>
      </n-space>
      <p>团长：用户{{ mall.userid }}</p>
      <p>{{ mallDescription }}</p>
      <template #action>
        <n-button>浏览</n-button>
      </template>
    </n-card>
  </div>
</template>
