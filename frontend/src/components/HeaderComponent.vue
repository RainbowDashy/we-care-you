<script setup>
import { NIcon, NButton } from "naive-ui"
import { ChevronBackOutline, ChevronForwardOutline } from "@vicons/ionicons5"
import { useRouter } from "vue-router"

const router = useRouter()
</script>

<template>
  <div class="header">
    <n-button text class="header-button" @click="router.back">
      <n-icon><ChevronBackOutline /></n-icon>
    </n-button>
    <span class="header-title">We Care You</span>
    <n-button text class="header-button" @click="router.forward">
      <n-icon><ChevronForwardOutline /></n-icon>
    </n-button>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 24px;
  height: 50px;
}
.header-title {
  font-family: cursive;
}

.header-button {
  font-size: 24px;
}
</style>
