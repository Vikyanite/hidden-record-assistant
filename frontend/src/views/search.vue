<script lang="ts" setup>
import {onBeforeMount, reactive, ref} from "vue";

import Ranks from "../components/Ranks.vue";
import RecentSummary from "../components/RecentSummary.vue";

import BaseStats from "../components/BaseStats.vue";
import Match from "../components/Match.vue";
import {model} from "../../wailsjs/go/models";
import {useStore} from "vuex";

const summoner = ref(new model.Summoner())

const store = useStore()

onBeforeMount(() => {
  summoner.value = store.state.currentSummoner
})
</script>

<template>
  <div class="about">
    <div class="about_left">
      <BaseStats
        :data="summoner.dataAccount"
        class="about_left_profile"
      >
      </BaseStats>

      <Ranks
        class="about_left_rank"
        :rank_solo="summoner.dataRank_solo"
        :rank_flex="summoner.dataRank_flex"
      ></Ranks>

                                                              <!-- 这部分是top熟练度的英雄-->
<!--      <div v-if="account.champs_points.length == 10" class="about_left_champs">
        &lt;!&ndash;pun 10, pt ca presupun ca toate lumea are cel putin 10 campioni cu care au jucat macar odata&ndash;&gt;
        <champsPoints
          :champs_points="account.champs_points"
          :region="data.region"
          :summonerID="account.dataAccount.id"
        ></champsPoints>
      </div>
      <div
        class="about_left_champs"
        v-else
        style="display: flex; align-items: center; justify-content: center"
      >
        <svg
          class="loading_spin"
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          aria-hidden="true"
          role="img"
          width="1em"
          height="1em"
          preserveAspectRatio="xMidYMid meet"
          viewBox="0 0 24 24"
        >
          <circle cx="12" cy="2" r="0" fill="currentColor">
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(45 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.125s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(90 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.25s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(135 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.375s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(180 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.5s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(225 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.625s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(270 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.75s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
          <circle
            transform="rotate(315 12 12)"
            cx="12"
            cy="2"
            r="0"
            fill="currentColor"
          >
            <animate
              attributeName="r"
              values="0;2;0;0"
              dur="1s"
              repeatCount="indefinite"
              begin="0.875s"
              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"
              calcMode="spline"
            />
          </circle>
        </svg>
      </div>-->
    </div>

    <div class="about_right">
      <div
        class="about_right_summary"
      >
        <RecentSummary :data="summoner.matchData"/>
      </div>

<!--      &lt;!&ndash;LOADING PAGE&ndash;&gt;-->
<!--      <div-->

<!--        style="display: flex; align-items: center; justify-content: center"-->
<!--      >-->
<!--        <svg-->
<!--          class="loading_spin"-->
<!--          xmlns="http://www.w3.org/2000/svg"-->
<!--          xmlns:xlink="http://www.w3.org/1999/xlink"-->
<!--          aria-hidden="true"-->
<!--          role="img"-->
<!--          width="1em"-->
<!--          height="1em"-->
<!--          preserveAspectRatio="xMidYMid meet"-->
<!--          viewBox="0 0 24 24"-->
<!--        >-->
<!--          <circle cx="12" cy="2" r="0" fill="currentColor">-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(45 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.125s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(90 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.25s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(135 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.375s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(180 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.5s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(225 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.625s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(270 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.75s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--          <circle-->
<!--            transform="rotate(315 12 12)"-->
<!--            cx="12"-->
<!--            cy="2"-->
<!--            r="0"-->
<!--            fill="currentColor"-->
<!--          >-->
<!--            <animate-->
<!--              attributeName="r"-->
<!--              values="0;2;0;0"-->
<!--              dur="1s"-->
<!--              repeatCount="indefinite"-->
<!--              begin="0.875s"-->
<!--              keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8"-->
<!--              calcMode="spline"-->
<!--            />-->
<!--          </circle>-->
<!--        </svg>-->
<!--      </div>-->

      <div class="about_right_matchesList">
        <div v-for="(i, index) in summoner.displayMatchHistory" :key="index">
          <div v-if="summoner.displayMatchHistory[index]">
            <Match
              :data="summoner.displayMatchHistory[index]"
            ></Match>
          </div>
        </div>
      </div>

    </div>
  </div>

</template>


<style scoped lang="scss">

.about {
  margin-left: 15%;
  margin-right: 15%;
  margin-top: 5rem;

  display: grid;
  grid-template-columns: 35% 65%;
  gap: 3rem;

  @media screen and (min-width: 2000px) {
    margin-left: 20%;
    margin-right: 20%;
  }

  @media screen and (max-width: 1500px) {
    margin-left: 5%;
    margin-right: 5%;
  }
  @media screen and (max-width: 1350px) {
    margin-left: 2%;
    margin-right: 4%;
  }
  @media screen and (max-width: 1000px) {
    margin-left: 2%;
    margin-right: 2%;
  }

  @media screen and (max-width: 1200px) {
    grid-template-columns: 30% 70%;
  }

  @media screen and (max-width: 1000px) {
    grid-template-columns: 1fr;
    margin-top: 2rem;
  }

  &_right {
    &_summary,
    &_live {
      color: #fffffe;
      background-color: #242629;
      border: 1px solid rgba(114,117,126, 0.2);
      font-size: 1.4rem;
      margin-bottom: 2rem;
    }
    &_matchesList {
      color: #fffffe;
      background-color: #242629;
      border: 1px solid rgba(114,117,126, 0.2);
      font-size: 1.4rem;
      margin-bottom: 1rem;
    }
  }

  &_left {
    display: flex;
    flex-direction: column;

    &_liveBtn {
      padding: 1rem;
      margin-top: 1rem;
      background-color: var(--color-win);
      border: 1px solid rgba(114,117,126, 0.2);
      font-family: inherit;
      font-weight: bold;
      color: #fffffe;
      cursor: pointer;

      @media screen and (max-width: 500px) {
        font-weight: normal;
      }
    }

    &_champs {
      @media screen and (max-width: 1000px) {
        grid-column: 1/-1;
      }
    }

    @media screen and (max-width: 1000px) {
      display: grid;
      grid-template-columns: max-content 1fr;
      grid-template-rows: max-content max-content;
      gap: 1rem;
      align-items: stretch;

      &_liveBtn {
        grid-column: 1/-1;
        grid-row: 2/-1;
      }

      & > * {
        height: 100%;
        margin: 0;
      }
    }

    @media screen and (max-width: 500px) {
      grid-template-columns: 1fr;
    }
  }
}
</style>
