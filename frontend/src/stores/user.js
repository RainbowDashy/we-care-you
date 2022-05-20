import { defineStore } from "pinia"
import http from "../http"

export const useUserStore = defineStore("user", {
  state: () => ({
    id: 0,
    username: "",
    location: "",
    token: "",
  }),
  getters: {
    logined: (state) => state.token !== "",
  },
  actions: {
    async fetch() {
      let token = window.localStorage.getItem("token")
      if (!token) {
        return
      }
      this.token = token
      try {
        let res = await http.get("/users")
        let data = await res.json()
        this.id = data.id
        this.username = data.username
        this.location = data.location
      } catch {
        // token is expired
        this.token = ""
        window.localStorage.removeItem("token")
      }
    },
    async login(username, password) {
      let res = await http.post("/login", {
        body: JSON.stringify({
          username: username,
          password: password,
        }),
      })
      let data = await res.json()
      window.localStorage.setItem("token", data.token)

      // fetch other user information
      this.fetch()
    },
  },
})
