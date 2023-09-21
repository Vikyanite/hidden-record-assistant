<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {InitBackend} from "../../wailsjs/go/backend/WailsApp";
import {useStore} from "vuex";

const router = useRouter()

onMounted(() => {
  Init()
})

enum Status {
  success = "success",
  info = "info",
  error = "error"
}

const Result = {
  success: {
    title: "请求成功",
    subTitle: "欢迎使用",
    status: Status.success
  },
  info: {
    title: "加载中",
    subTitle: "正在请求后台服务",
    status: Status.info
  },
  error: {
    title: "出错啦",
    subTitle: "",
    status: Status.error
  }
}

const result = ref(Result.info)

const store = useStore()

function Init() {
  result.value = Result.info
  setTimeout(() => {
    InitBackend()
        .then((res) => {
          result.value = Result.success
          store.commit('SetState', res);
          setTimeout(() => {
            router.push("/home")
          }, 1000)
          //router.push("/home")

        })
        .catch((err: any) => {
          result.value = Result.error
          Result.error.subTitle = err.toString()
        })
  }, 1000)
}

</script>

<template>
  <div class="load-page">
    <el-result
        :icon=result.status
        :title=result.title
        :subTitle=result.subTitle
        v-loading="result.status==Status.info"
    >
      <template #extra>
        <el-button v-show="result.status==Status.error" @click="Init" type="primary" size="default">重试</el-button>
      </template>
    </el-result>
  </div>
</template>

<style scoped>

.el-result {
  height: 100%;
}

.load-page {
  height: 100%;
}

</style>