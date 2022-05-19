//
// make fetch easier to use
//

import { useUserStore } from "./stores/user"

function injectToken(data) {
  let store = useUserStore()
  if (store.token !== "") {
    data.headers = {
      Authorization: `Bearer ${store.token}`,
      ...data.headers,
    }
  }
}

function injectContentType(data) {
  if ("body" in data) {
    data.headers = {
      "Content-Type": "application/json",
      ...data.headers,
    }
  }
}

export default {
  baseURL: "http://localhost:8080/api",
  setBaseURL(url) {
    this.baseURL = url
  },
  concatURL(url) {
    return `${this.baseURL}${url}`
  },
  handleErr(res) {
    return new Promise((resolve, reject) => {
      if (res.status >= 400) {
        return res.json().then((msg) =>
          reject({
            code: res.status,
            ...msg,
          })
        )
      } else {
        resolve(res)
      }
    })
  },
  get(url = "", data = {}) {
    injectToken(data)
    return fetch(this.concatURL(url), {
      method: "GET",
      ...data,
    }).then(this.handleErr)
  },
  post(url = "", data = {}) {
    injectToken(data)
    injectContentType(data)
    return fetch(this.concatURL(url), {
      method: "POST",
      ...data,
    }).then(this.handleErr)
  },
  patch(url = "", data = {}) {
    injectToken(data)
    injectContentType(data)
    return fetch(this.concatURL(url), {
      method: "PATCH",
      ...data,
    }).then(this.handleErr)
  },
}
