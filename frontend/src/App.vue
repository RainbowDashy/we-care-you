<script setup>
import { RouterView } from "vue-router"
import { useUserStore } from "./stores/user"
import { onMounted } from "vue"
import NavigationComponent from "./components/NavigationComponent.vue"
import HeaderComponent from "./components/HeaderComponent.vue"

const user = useUserStore()

onMounted(() => {
  user.fetch()
})
</script>

<template>
  <header>
    <HeaderComponent />
  </header>
  <RouterView class="router-view" />
  <footer>
    <nav>
      <NavigationComponent />
    </nav>
  </footer>
</template>

<style>
.router-view {
  padding-top: 60px;
  padding-bottom: 50px;
}
header {
  position: fixed;
  top: 0;
  width: 100vw;
  background-color: white;
  box-shadow: 0 1px 3px #f0f1f2;
  z-index: 99;
}
footer {
  position: fixed;
  bottom: 0;
  background-color: rgb(245, 245, 245);
  width: 100vw;
  padding-top: 5px;
  height: 45px;
  z-index: 99;
}
</style>
