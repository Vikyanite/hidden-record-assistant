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

@import "./styles/search_page.scss"; // search_page styles

#app {
  --el-menu-base-level-padding: 5px;
  --el-menu-icon-width: 25px;
  --el-menu-item-height: 30px;

  --el-menu-item-margin: 5px;
}



@import url("https://fonts.googleapis.com/css2?family=Poppins&display=swap");

:root {
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
  overflow-x: hidden;

  @media screen and (max-width: 75em) {
    font-size: 56.25%;
  }

  @media screen and (max-width: 56.25em) {
    font-size: 50%;
  }

  @media screen and (max-width: 37.5em) {
    font-size: 50%;
  }

  @media screen and (min-width: 121.875em) {
    font-size: 65%;
  }
}
body {
  font-family: "Poppins", sans-serif;
  font-weight: 400;
  line-height: 1.7;
  overflow-x: hidden;
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
.loading_spin{
  width: 4rem;
  height: 4rem;
  margin: 2rem 0 2rem 0;
}

</style>