<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {InitConnector} from "../../wailsjs/go/main/App";


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


function Init() {
  result.value = Result.info
  setTimeout(() => {
    InitConnector()
        .then(() => {
          result.value = Result.success
          setTimeout(() => {
            router.push("/home")
          }, 1000)
          //router.push("/home")

        })
        .catch((err: any) => {
          result.value = Result.error
          Result.error.subTitle = err.toString()
        })
  }, 2000)
}

</script>

<template>
  <div>
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
.myfont {
  font-size: 50px;
  color: #333333;
}
</style>