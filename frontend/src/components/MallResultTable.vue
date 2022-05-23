<script setup>
import { NDataTable } from "naive-ui"
import { onMounted, ref } from "vue"
import http from "../http"

const createColumns = (items) => {
  let res = [
    {
      title: "User ID",
      key: "userid",
    },
  ]
  for (let item of items) {
    res.push({
      title: item.name,
      key: "item" + item.id,
    })
  }
  return res
}

const createRows = (orders) => {
  const ordersByUser = orders.sort((a, b) => {
    if (a.userid == b.userid) {
      return a.itemid - b.itemid
    }
    return a.userid - b.userid
  })

  const item2key = (id) => `item${id}`

  let rows = []
  let summary = { userid: "总计" }
  let o = {}
  for (let order of ordersByUser) {
    if (o.userid == undefined) {
      o.userid = order.userid
    } else if (o.userid != order.userid) {
      rows.push(o)
      o = {}
      o.userid = order.userid
    }
    if (o[item2key(order.itemid)] == undefined) {
      o[item2key(order.itemid)] = order.buycount
    } else {
      o[item2key(order.itemid)] += order.buycount
    }
    summary[item2key(order.itemid)] = summary[item2key(order.itemid)]
      ? summary[item2key(order.itemid)] + order.buycount
      : order.buycount
  }
  rows.push(o)
  rows.push(summary)
  return rows
}

const props = defineProps({
  mallId: {
    type: Number,
    required: true,
  },
})
const items = ref([])
const orders = ref([])
const tableColumns = ref([])
const tableRows = ref([])

const fetchItems = async () => {
  try {
    const res = await http.get(`/items?mallid=${props.mallId}`)
    items.value = await res.json()
  } catch (err) {
    console.log(err)
  }
}

const fetchOrders = async () => {
  const res = await http.get(`/orders?mallid=${props.mallId}`)
  orders.value = await res.json()
}

onMounted(async () => {
  await fetchOrders()
  await fetchItems()
  const columns = createColumns(items.value)
  const rows = createRows(orders.value)

  tableColumns.value = columns
  tableRows.value = rows
})
</script>

<template>
  <main>
    <h1>Mall Result</h1>
    <n-data-table :columns="tableColumns" :data="tableRows" />
  </main>
</template>
