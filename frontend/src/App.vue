<script lang="ts" setup>

import { useRouter, useRoute } from 'vue-router'
import {nextTick, onMounted, ref} from "vue";
import routes from "./routes/router"
import Loading from "./views/loading.vue";
import Home from "./views/home.vue";
import {EventsOn} from "../wailsjs/runtime";

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


<style>

</style>