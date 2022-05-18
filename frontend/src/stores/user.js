import { defineStore } from "pinia"
import http from "../http"

export const useUserStore = defineStore("user", {
  state: () => ({
    username: "",
    password: "",
    location: "",
    token: window.localStorage.getItem("token") || "",
  }),
  getters: {
    logined: (state) => state.token !== "",
  },
  actions: {
    async login() {
      let res = await http.post("/login", {
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      })
      let data = await res.json()
      this.token = data.token
      window.localStorage.setItem("token", data.token)
    },
  },
})
