<script setup>
import { NButton } from "naive-ui"
import http from "../http"
const props = defineProps({
  mallId: {
    type: Number,
    required: true,
  },
})

const exportExcel = async () => {
  const res = await http.get(`/excel?mallid=${props.mallId}`)
  const blob = await res.blob()
  const url = window.URL.createObjectURL(blob)
  const a = document.createElement("a")
  a.style.display = "none"
  a.href = url
  a.download = `${props.mallId}.xls`
  document.body.appendChild(a)
  a.click()
  window.URL.revokeObjectURL(url)
}
</script>

<template>
  <n-button @click="exportExcel">导出Excel</n-button>
</template>
