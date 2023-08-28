<script setup lang="ts">
import {useRouter} from "vue-router";

const router = useRouter()

function ToPage2() {
  router.push("/test2")
}

</script>

<template>
  <div>
    test1 page
    <el-button @click="ToPage2">
      ToPage2
    </el-button>
  </div>

</template>

<style scoped>

</style>