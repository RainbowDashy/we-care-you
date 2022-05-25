<script setup>
import { computed, onMounted, ref } from "vue"
import http from "../http"
import { NCard, NTime, NSpace, NButton } from "naive-ui"
import { useRouter } from "vue-router"

const props = defineProps({
  mall: {
    require: true,
  },
  search: String,
})

const items = ref([])

const mallDescription = computed(() => {
  return items.value.map((item) => item.name).join(" ")
})

const show = computed(() => {
  return (
    items.value.filter((item) => {
      return (
        item.name.includes(props.search) ||
        item.description.includes(props.search)
      )
    }).length > 0
  )
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

const router = useRouter()

const goToMallView = () => {
  router.push(`/malls/${props.mall.id}`)
}

onMounted(fetchItems)
</script>

<template>
  <div class="mall" v-show="show">
    <n-card :title="`拼团${mall.id}`" size="small">
      <n-space justify="space-between">
        <p>开始：<n-time :time="mall.begintime" format="M月d日H时mm分" /></p>
        <p>结束：<n-time :time="mall.endtime" format="M月d日H时mm分" /></p>
      </n-space>
      <p>团长：用户{{ mall.userid }}</p>
      <p>{{ mallDescription }}</p>
      <template #action>
        <n-button @click="goToMallView">浏览</n-button>
      </template>
    </n-card>
  </div>
</template>

