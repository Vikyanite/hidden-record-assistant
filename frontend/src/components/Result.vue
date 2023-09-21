<script lang="ts" setup>
import {ref} from "vue";

import RecentSummary from "./RecentSummary.vue";

import BaseStats from "./BaseStats.vue";
import Match from "./Match.vue";

import {model} from "../../wailsjs/go/models";
import {useStore} from "vuex";
import {useRouter} from "vue-router";


const props = defineProps<{
  summoner: model.Summoner,
}>()

const store = useStore()

const router = useRouter()

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
        <el-scrollbar height="100%">
          <div v-for="(i, index) in summoner.displayMatchHistory" :key="index">
            <Match
                :data="summoner.displayMatchHistory[index]"
            />
          </div>
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
