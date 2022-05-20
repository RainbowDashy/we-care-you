<script setup>
import { RouterView } from "vue-router"
import { useUserStore } from "./stores/user"
import { onMounted } from "vue"
import NavigationComponent from "./components/NavigationComponent.vue"

const user = useUserStore()

onMounted(() => {
  user.fetch()
})
</script>

<template>
  <header>
    <nav>
      <NavigationComponent />
    </nav>
  </header>
  <RouterView />
</template>
