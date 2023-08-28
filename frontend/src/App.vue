<script lang="ts" setup>

import { useRouter, useRoute } from 'vue-router'
import {onMounted, ref} from "vue";
import routes from "./routes/router"
import Loading from "./views/loading.vue";
import Home from "./views/home.vue";

const isWait = ref(false)

function handleOpen(key:string, keyPath:string) {
  console.log(key, keyPath);
}

function handleClose(key :string, keyPath :string) {
  console.log(key, keyPath);
}

const router = useRouter()

onMounted(() => {
  router.push("/loading")
})

</script>

<template>
    <router-view v-slot="{ Component }">
      <transition name="el-fade-in-linear" appear>
        <component :is="Component"/>
      </transition>
    </router-view>


<!--  <home/>-->

</template>


<style>

.el-header, .el-footer {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

</style>