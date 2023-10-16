<script setup lang="ts">
import {useRouter} from "vue-router";
import Result from "../components/Result.vue";
import {model} from "../../wailsjs/go/models";
import {onActivated, onMounted, Ref, ref} from "vue";
import {GetSummonerByName} from "../../wailsjs/go/backend/WailsApp";
import {useStore} from "vuex";
import {DefaultSummoner} from "../store";
import {RefreshRight} from "@element-plus/icons-vue";

const store = useStore()
const router = useRouter()
const loading = ref(true)
const summoner = ref(DefaultSummoner)

onMounted(LoadFunc)

function LoadFunc() {
  const name = router.currentRoute.value.params.name as string
  loading.value = true
  GetSummonerByName(name).then((res: model.DisplaySummoner) => {
    summoner.value = res
  }).catch((err: any) => {
    console.log(err)
  }).finally(() => {
    loading.value = false
  })
}

</script>

<template>
  <div>
    <Result  v-loading.fullscreen.lock="loading" :summoner="summoner"/>
    <el-button
        v-if="!loading"
        @click="LoadFunc"
        circle
        class="refresh-button"
    >
    <el-icon style="width: 25px;height: 25px;">
      <RefreshRight style="width: 25px;height: 25px;"/>
    </el-icon>
  </el-button>
  </div>

</template>

<style scoped lang="scss">
.refresh-button {
  width: 50px;
  height: 50px;
  position: fixed;
  right: 40px;
  bottom: 40px;
  z-index: 999;
}
</style>