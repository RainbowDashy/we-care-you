<script setup>
import { useRoute } from "vue-router"
import MallResultTable from "../components/MallResultTable.vue"
import ExcelExportComponent from "../components/ExcelExportComponent.vue"

const route = useRoute()
const mallId = parseInt(route.params.id)
</script>

<template>
  <MallResultTable :mallId="mallId" />
  <ExcelExportComponent :mallId="mallId" />
</template>
