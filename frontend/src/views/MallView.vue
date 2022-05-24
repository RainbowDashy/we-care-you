<script setup>
import { NButton } from "naive-ui"
import { useRoute, useRouter } from "vue-router"
import ItemList from "../components/ItemList.vue"

const route = useRoute()
const router = useRouter()
const mallId = parseInt(route.params.id)

const jumpToView = () => {
  router.push(`/malls/view/${mallId}`)
}
</script>

<template>
  <div>
    <n-button @click="jumpToView">查看购买记录</n-button>
    <ItemList :mallId="mallId" />
  </div>
</template>
