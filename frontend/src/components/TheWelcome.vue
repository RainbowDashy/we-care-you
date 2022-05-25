<script setup>
import WelcomeItem from "./WelcomeItem.vue"
import DocumentationIcon from "./icons/IconDocumentation.vue"
import ToolingIcon from "./icons/IconTooling.vue"
import EcosystemIcon from "./icons/IconEcosystem.vue"
import CommunityIcon from "./icons/IconCommunity.vue"
import SupportIcon from "./icons/IconSupport.vue"
import { useUserStore } from "../stores/user.js"
import {NFormItem} from "naive-ui"
const user = useUserStore()
const begintime=new Date().toLocaleString()

</script>

<template>
<div id="building">
  <div >
    <h1 style=" color: #227700; text-align: center">
    {{user.username}}我们与你在一起
    </h1>
    <h2 class="cur_time">当前时间是{{begintime}}</h2>
    <img class="img" src="../photos/1.png" alt="loading..." 
    >
  </div>

<!--   <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>
 -->
  <!--   Vue’s
    <a target="_blank" href="https://vuejs.org/">official documentation</a>
    provides you with all information you need to get started.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <ToolingIcon />
    </template>
    <template #heading>Tooling</template>

    This project is served and bundled with
    <a href="https://vitejs.dev/guide/features.html" target="_blank">Vite</a>.
    The recommended IDE setup is
    <a href="https://code.visualstudio.com/" target="_blank">VSCode</a> +
    <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>.
    If you need to test your components and web pages, check out
    <a href="https://www.cypress.io/" target="_blank">Cypress</a> and
    <a
      href="https://docs.cypress.io/guides/component-testing/introduction"
      target="_blank"
      >Cypress Component Testing</a
    >.

    <br />

    More instructions are available in <code>README.md</code>.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <EcosystemIcon />
    </template>
    <template #heading>Ecosystem</template>

    Get official tools and libraries for your project:
    <a target="_blank" href="https://pinia.vuejs.org/">Pinia</a>,
    <a target="_blank" href="https://router.vuejs.org/">Vue Router</a>,
    <a target="_blank" href="https://test-utils.vuejs.org/">Vue Test Utils</a>,
    and
    <a target="_blank" href="https://github.com/vuejs/devtools">Vue Dev Tools</a
    >. If you need more resources, we suggest paying
    <a target="_blank" href="https://github.com/vuejs/awesome-vue"
      >Awesome Vue</a
    >
    a visit.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Community</template>

    Got stuck? Ask your question on
    <a target="_blank" href="https://chat.vuejs.org">Vue Land</a>, our official
    Discord server, or
    <a target="_blank" href="https://stackoverflow.com/questions/tagged/vue.js"
      >StackOverflow</a
    >. You should also subscribe to
    <a target="_blank" href="https://news.vuejs.org">our mailing list</a> and
    follow the official
    <a target="_blank" href="https://twitter.com/vuejs">@vuejs</a>
    twitter account for latest news in the Vue world.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <SupportIcon />
    </template>
    <template #heading>Support Vue</template>

    As an independent project, Vue relies on community backing for its
    sustainability. You can help us by
    <a target="_blank" href="https://vuejs.org/sponsor/">becoming a sponsor</a>.
  </WelcomeItem> -->
 </div>
</template>
<style scoped>
#building{
  background:url("../photos/2.png");
  width:100%;
  height:100%;
/*   background-position:0px 100px; */
  background-size:100% 100%;
  background-repeat: no-repeat;
}
.img{
  width:94%;
  height:auto;
  margin-left:3%;
  margin-top:10%;

/* position: absolute;
width:528px ;
height:330px ;
margin-left: 3%;
top: 40%;
height: 30%;
margin-top: -15%; */
}

.cur_time{ 
  margin-left:3%;
}
</style>  
