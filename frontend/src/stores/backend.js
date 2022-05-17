import { defineStore } from "pinia"

export const useBackendStore = defineStore("backend", {
  state: () => ({
    url: "127.0.0.1:8080",
  }),
})
