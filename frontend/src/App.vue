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
  <header></header>
  <RouterView class="router-view" />
  <footer>
    <nav>
      <NavigationComponent />
    </nav>
  </footer>
</template>

<style>
.router-view {
  padding-bottom: 50px;
}
footer {
  position: fixed;
  bottom: 0;
  background-color: white;
  width: 100vw;
  height: 50px;
}
</style>
