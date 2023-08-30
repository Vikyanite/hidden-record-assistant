<script lang="ts" setup>

import { useRouter, useRoute } from 'vue-router'
import {onMounted, ref} from "vue";
import routes from "./routes/router"
import Loading from "./views/loading.vue";
import Home from "./views/home.vue";

const router = useRouter()

onMounted(() => {
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