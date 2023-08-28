<script setup lang="ts">
import {useRouter} from "vue-router";

const router = useRouter()

function ToHome() {
  router.push("/")
}

</script>

<template>
  <div>
    test2 page
    <el-button @click="ToHome">
      ToHome
    </el-button>
  </div>

</template>

<style scoped>

</style>