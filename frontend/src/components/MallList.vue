<script setup>
import MallListItem from "./MallListItem.vue"
import { computed, onMounted, ref } from "vue"
import http from "../http"

defineProps({
  search: {
    type: String,
    default: "",
  },
})

const malls = ref([])
const filteredMalls = computed(() => {
  return malls.value.filter((mall) => {
    return mall.state == 1
  })
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

onMounted(refresh)
</script>

<template>
  <!-- <p>The search string is {{ search }}</p> -->
  <MallListItem
    v-for="mall in filteredMalls"
    :key="mall.id"
    :mall="mall"
  ></MallListItem>
</template>
