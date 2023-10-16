<script lang="ts" setup>
import {computed, ref} from "vue";

import RecentSummary from "./RecentSummary.vue";

import BaseStats from "./BaseStats.vue";
import Match from "./Match.vue";

import {model} from "../../wailsjs/go/models";
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import {GetMatchRecordsByPuuid} from "../../wailsjs/go/backend/WailsApp";


const props = defineProps<{
  summoner: model.DisplaySummoner,
}>()

const store = useStore()

const router = useRouter()

const loading = ref(false)
const noMore = ref(false)
const disabled = computed(() => loading.value || noMore.value)

const nowMatchIndex = ref(0)
const endIndex = ref(19)

const delta = 10

const loadHandler = () => {
  loading.value = true

  GetMatchRecordsByPuuid(props.summoner.dataAccount.puuid, endIndex.value+1, endIndex.value+delta).then((res: model.DisplayMatch[]) => {
    if (res.length === 0) {
      noMore.value = true
      return
    } else {
      noMore.value = false
      props.summoner.displayMatchHistory.push(...res)
      endIndex.value = endIndex.value+delta
    }
  }).catch((err: any) => {
    console.log(err)
  }).finally(() => {
    loading.value = false
  })
}

</script>

<template>
  <div class="about">
    <div class="about_left">
      <BaseStats
          :data="summoner.dataAccount"
          :solo="summoner.dataRank_solo"
          :flex="summoner.dataRank_flex"
      />
      <RecentSummary :data="summoner.matchData"/>
    </div>

    <div class="about_right">
      <div class="about_right_matchesList">
        <el-scrollbar
            height="100%"
        >
          <div
              v-infinite-scroll="loadHandler"
              :infinite-scroll-disabled="disabled"
              infinite-scroll-immediate="false"
          >
            <Match v-for="(i, index) in summoner.displayMatchHistory" :key="index"
                :data="summoner.displayMatchHistory[index]"
            />
          </div>
          <p v-if="loading">Loading...</p>
          <p v-if="noMore">No more</p>
        </el-scrollbar>
      </div>
    </div>
  </div>

</template>

<style scoped lang="scss">
.about {

  display: grid;
  grid-template-columns: 38.2% 61.8%;
  gap: 1rem;

  &_right {
    &_matchesList {
      height: calc(var(--wails-app-height) - var(--el-main-padding)*2);
      color: #fffffe;
      background-color: #242629;
      border: 1px solid rgba(114,117,126, 0.2);
      font-size: 1.4rem;
    }
  }

  &_left {
    display: grid;
    grid-template-rows: 13% 87%;
  }
}
</style>
