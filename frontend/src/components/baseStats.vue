<template>
  <div class="profile" v-if="data">
    <div class="profile_img">
      <img
        :src="store.getters.LCUAPIPrefix() + '/lol-game-data/assets/v1/profile-icons/' + data.profileIconId + '.jpg'"
        alt="profile icon"
      />
    </div>
    <div class="profile_stats">
      <h1 class="profile_name">
        <p style="white-space:nowrap;">{{ data.displayName }}</p>
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
      </h1>
      <p>Level: {{ data.summonerLevel }}</p>
<!--      <p>Last updated: {{ time }}</p>-->
    </div>
  </div>
</template>

<script lang="ts" setup>

  import {ISummoner} from "../types/types";
  import {useStore} from "vuex";
  const store = useStore()
  defineProps({
    data: {
      type: Object as () => ISummoner,
    },
  })

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
  grid-template-columns: min-content max-content;
  gap: 0.5rem;
  align-items: center;

  @media screen and (max-width: 1000px) {
    display: flex;
    flex-direction: column;
    padding: 0 2rem 0 2rem;
    text-align: center;

    &_name {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      // padding: .5rem;
      svg {
        margin-left: 0rem;
        margin: 0.5rem 0 1rem 0;
      }
    }
  }
  @media screen and (max-width: 500px) {
    flex-direction: row;
    text-align: left;
    padding: 0;

    &_name {
      display: flex;
      flex-direction: row;
    }
  }
  &_name {
    font-size: 2.4rem;
    display: flex;
    align-items: center;

    svg {
      //margin-left: 1rem;
      cursor: pointer;
    }
  }

  &_img {
    display: flex;
    align-items: center;
    img {
      width: 8rem;
      height: 8rem;
      border: 3px solid rgba(#72757e, 0.5);
    }
    padding: 1rem;
  }
}
</style>
