<script setup>
import { useRoute, useRouter } from "vue-router"
import MallResultTable from "../components/MallResultTable.vue"
import ExcelExportComponent from "../components/ExcelExportComponent.vue"
import { NButton } from "naive-ui"
const router = useRouter()
const route = useRoute()
const mallId = parseInt(route.params.id)
const returnToPurchase = () => {
  router.push(`/malls/${mallId}`)
}
</script>

<template>
  <n-button @click="returnToPurchase">返回购买</n-button>
  <MallResultTable :mallId="mallId" />
  <ExcelExportComponent :mallId="mallId" />
</template>
