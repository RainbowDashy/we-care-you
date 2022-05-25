<script setup>
import { ref } from "vue"
import { NInput,NIcon } from "naive-ui"
import MallList from "../components/MallList.vue"
import {
  Search
} from "@vicons/ionicons5"
let search = ref("")
</script>

<template>
  <div>
    <div class="search">
      <n-input placeholder="搜索" v-model:value="search">
        <template #prefix>
        <n-icon :component="Search" />
      </template>
      </n-input>
    </div>
    <div>
      <MallList :search="search"></MallList>
    </div>
  </div>
</template>
