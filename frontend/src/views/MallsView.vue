<script setup>
import { ref } from "vue"
import { NInput } from "naive-ui"
import MallList from "../components/MallList.vue"
let search = ref("")
</script>

<template>
  <div>
    <div class="search">
      <n-input round placeholder="搜索" v-model:value="search" />
    </div>
    <div>
      <MallList :search="search"></MallList>
    </div>
  </div>
</template>
