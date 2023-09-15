<script setup lang="ts">
import {useRouter} from "vue-router";
import Result from "../components/result.vue";
import {model} from "../../wailsjs/go/models";
import {onActivated, Ref, ref} from "vue";
import {GetSummonerByName} from "../../wailsjs/go/backend/WailsApp";
import {useStore} from "vuex";
import {DefaultSummoner} from "../store";

const store = useStore()
const router = useRouter()
const loading = ref(true)
const summoner = ref(DefaultSummoner)

onActivated(() => {
  const name = router.currentRoute.value.params.name as string
  loading.value = true
  GetSummonerByName(name).then((res: model.Summoner) => {
    summoner.value = res
  }).catch((err: any) => {
    console.log(err)
  }).finally(() => {
    loading.value = false
  })

})

</script>

<template>
    <Result v-loading="loading" :summoner="summoner"/>
</template>

<style scoped lang="scss">

</style>