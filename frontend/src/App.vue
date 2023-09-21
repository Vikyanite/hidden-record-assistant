<script lang="ts" setup>

import { useRouter, useRoute } from 'vue-router'
import {nextTick, onMounted, ref} from "vue";
import routes from "./routes/router"
import Loading from "./views/loading.vue";
import Home from "./views/home.vue";
import {EventsOn} from "../wailsjs/runtime";
import {PieChart} from "@element-plus/icons-vue";

const router = useRouter()

onMounted(() => {
  nextTick(() => {
    EventsOn("disconnected", (...data: any) => {
      console.log("event: disconnected")
      router.push("/loading")
      console.log("to loading")
    })
  })

  router.push("/loading")
})



</script>

<template>
    <router-view v-slot="{ Component }" class="app">
      <transition name="el-fade-in-linear" appear>
        <component :is="Component"/>
      </transition>
    </router-view>
<!--  <home/>-->

</template>


<style >

@import "./styles/search_page.scss";

:root {
  --wails-window-height: 720px;
  --wails-window-width: 1280px;
  --wails-nav-height: 40px;

  --wails-app-height: calc(var(--wails-window-height) - var(--wails-nav-height));

  --color-win: rgb(32,178,170);
  --color-lose: rgb(240, 128, 128);
  --color-remake: rgb(192,192,192);

  --color-win07: rgba(32, 178, 170, 0.7);
  --color-win01: rgba(32, 178, 170, 0.1);

  --color-lose07: rgba(240, 128, 128, 0.7);
  --color-lose01: rgba(240, 128, 128, 0.1);

  --color-remake07: rgba(192,192,192,.7);
  --color-remake01: rgba(192,192,192,.1);

  --color-yellow: rgb(207, 179, 60);
}

html {
  font-size: 62.5%;
  scroll-behavior: smooth;
}

#app {
  --el-menu-base-level-padding: 5px;
  --el-menu-icon-width: 25px;
  --el-menu-item-height: 30px;

  --el-menu-item-margin: 0;
}

body {


  font-family: "Poppins", sans-serif;
  font-weight: 400;
  line-height: 1.7;

  box-sizing: border-box;
  color: #fffffe;
}
* {
  margin: 0;
  padding: 0;
}

*,
*::after,
*::before {
  box-sizing: inherit;
}

</style>