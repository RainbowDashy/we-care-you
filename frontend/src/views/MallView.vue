<script setup>
import { useRoute } from "vue-router"
import ItemList from "../components/ItemList.vue"

const route = useRoute()
const mallId = parseInt(route.params.id)
</script>

<template>
  <ItemList :mallId="mallId" />
</template>
