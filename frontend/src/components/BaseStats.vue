<template>
  <div class="profile">
    <div class="profile_top">
      <div class="profile_top_stats">
        <div class="profile_top_stats_img">
          <img
              :src="store.getters.LCUAPIPrefix('/lol-game-data/assets/v1/profile-icons/' + data.profileIconId + '.jpg')"
              alt="profile icon"
          />
        </div>
        <div class="profile_top_stats_name">
          <p >{{ data.displayName }}</p>
          <!--        这里是收藏召唤师功能-->
          <!--        <svg -->
          <!--          class="bookmark"-->
          <!--          :class="saved == true ? 'hidden' : ''"-->
          <!--          @click="saveToLocalStorage()"-->
          <!--          xmlns="http://www.w3.org/2000/svg"-->
          <!--          xmlns:xlink="http://www.w3.org/1999/xlink"-->
          <!--          aria-hidden="true"-->
          <!--          role="img"-->
          <!--          width="1em"-->
          <!--          height="1em"-->
          <!--          preserveAspectRatio="xMidYMid meet"-->
          <!--          viewBox="0 0 16 16"-->
          <!--        >-->
          <!--          <g fill="var(&#45;&#45;color-win)">-->
          <!--            <path-->
          <!--              d="M2 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v13.5a.5.5 0 0 1-.777.416L8 13.101l-5.223 2.815A.5.5 0 0 1 2 15.5V2zm2-1a1 1 0 0 0-1 1v12.566l4.723-2.482a.5.5 0 0 1 .554 0L13 14.566V2a1 1 0 0 0-1-1H4z"-->
          <!--            />-->
          <!--          </g>-->
          <!--        </svg>-->
          <!--        <svg class="bookmark_remove"  :class="saved == false ? 'hidden' : ''" @click="deleteFromLocalStorage()" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 16 16"><g fill="var(&#45;&#45;color-win)"><path d="M2 2v13.5a.5.5 0 0 0 .74.439L8 13.069l5.26 2.87A.5.5 0 0 0 14 15.5V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2z"/></g></svg>-->
          <p>Level: {{ data.summonerLevel }}</p>

          <!--      <p>Last updated: {{ time }}</p>-->
        </div>
      </div>
      <div></div>
    </div>
    <div class="profile_bottom">
      <div></div>
      <div class="profile_bottom_ranks">
        <RankInfo :data="solo"/>
        <RankInfo :data="flex"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>

  import {useStore} from "vuex";
  import {model} from "../../wailsjs/go/models";
  import RankInfo from "./RankInfo.vue";
  const store = useStore()
  defineProps<{
    solo: model.RankDetails,
    flex: model.RankDetails,
    data: model.SummonerBaseInfo,
  }>()

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.hidden{
  display: none;
}
.bookmark{
  background-color: transparent !important;
}
.bookmark_remove{
  width: 50% !important;
}

.profile {
  color: #fffffe;
  background-color: #242629;
  border: 1px solid rgba(#72757e, 0.2);
  font-size: 1.4rem;

  display: grid;
  grid-template-columns: 61.8% 38.2%;

  &_top {
    display: grid;
    grid-template-rows: 61.8% 38.2%;

    &_stats {
      display: grid;
      grid-template-columns: max-content max-content;

      align-content: center;
      justify-content: center;

      &_name {
        display: flex;
        flex-direction: column;
        font-size: 2.0rem;
        white-space: nowrap;
        font-weight: bold;
        justify-content: center;

        svg {
          //margin-left: 1rem;
          cursor: pointer;
        }
      }

      &_img {
        img {
          width: 8rem;
          height: 8rem;
          border: 3px solid rgba(#72757e, 0.5);
        }
        padding: 1rem;
      }
    }

  }
  &_bottom {
    display: grid;
    grid-template-rows: 38.2% 61.8%;

    &_ranks {
      display: grid;
      grid-template-columns: 50% 50%;
      font-size: 1.4rem;
      align-content: center;

      margin-top: -20px;
    }
  }


}
</style>
