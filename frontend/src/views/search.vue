<script setup lang="ts">
  import {useRouter} from "vue-router";
  import {Search} from "@element-plus/icons-vue";
  import {ref} from "vue";

  const router = useRouter()
  const name = ref("")
</script>

<template>
  <div>
    <el-input v-model="name" placeholder="召唤师搜索" @keydown.enter="router.push('/home/result/' + name)">
      <template #prefix>
        <el-icon class="el-input__icon"><Search /></el-icon>
      </template>
    </el-input>
  </div>

</template>

<style scoped>

</style>