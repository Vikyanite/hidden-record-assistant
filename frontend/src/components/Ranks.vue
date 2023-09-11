<template>
  <div class="ranks">
    <div>
      <div class="ranks_rank" v-if="rank_solo.division!='NA'">
        <div class="ranks_rank_img">
          <img
            :src="store.getters.LocalAssetPrefix() + '/assets/images/tier/' + rank_solo.tier + '.png'"
            alt=""
          />
        </div>
        <div class="ranks_rank_text">
          <h3 class="ranks_rank_title">Ranked SOLO/DUO</h3>
          <p style="color: var(--color-win)">
            {{ rank_solo.tier }} {{ rank_solo.division }}
          </p>
          <p>
            {{ rank_solo.leaguePoints }}LP /
            <span class="ranks_rank_color"
              >{{ rank_solo.wins }}W {{ rank_solo.losses }}L
            </span>
          </p>
          <p class="ranks_rank_color">
            Win ratio:
            {{
              (
                (rank_solo.wins / (rank_solo.wins + rank_solo.losses)) * 100
              ).toFixed(0)
            }}%
          </p>
        </div>
      </div>
      <div v-else>
        <div class="ranks_rank">
          <div class="ranks_rank_img">
            <img :src="store.getters.LocalAssetPrefix() + '/assets/images/tier/unranked.png'" alt="unranked"/>
          </div>
          <div class="ranks_rank_text">
            <h3 class="ranks_rank_title">Ranked SOLO/DUO</h3>
            <p style="color: var(--color-win)">UNRANKED</p>
          </div>
        </div>
      </div>

      <div class="ranks_rank" v-if="rank_flex.division!='NA'">
        <div class="ranks_rank_img">
          <img
            :src="store.getters.LocalAssetPrefix() + '/assets/images/tier/' + rank_flex.tier + '.png'"
            alt=""
          />
        </div>
        <div class="ranks_rank_text">
          <h3 class="ranks_rank_title">Ranked FLEX</h3>
          <p style="color: var(--color-win)">
            {{ rank_flex.tier }} {{ rank_flex.division }}
          </p>
          <p>
            {{ rank_flex.leaguePoints }}LP /
            <span class="ranks_rank_color"
              >{{ rank_flex.wins }}W {{ rank_flex.losses }}L</span
            >
          </p>
          <p class="ranks_rank_color">
            Win ratio:
            {{
              (
                (rank_flex.wins /
                  (rank_flex.wins +
                    rank_flex.losses)) *
                100
              ).toFixed(0)
            }}%
          </p>
        </div>
      </div>
      <div v-else>
        <div class="ranks_rank">
          <div class="ranks_rank_img">
            <img :src="store.getters.LocalAssetPrefix() + '/assets/images/tier/unranked.png'" alt="unranked"/>
          </div>
          <div class="ranks_rank_text">
            <h3 class="ranks_rank_title">Ranked FLEX</h3>
            <p style="color: var(--color-win)">UNRANKED</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

  import {useStore} from "vuex";
  import {model} from "../../wailsjs/go/models";
  const store = useStore()
  

  const props = defineProps <{
    rank_solo: model.RankDetails,
    rank_flex: model.RankDetails,
  }>()

</script>

<style scoped lang="scss">
.ranks {
  margin-top: 1rem;
  background-color: #242629;
  border: 1px solid rgba(#72757e, 0.2);
  font-size: 1.4rem;
  padding: 1rem 0 1rem 0;

  & > * {
    & > *:not(:last-child) {
      margin-bottom: 1rem;
      border-bottom: 1px solid rgba(#72757e, 0.2);
      padding-bottom: 1rem;
    }
  }

  &_rank {
    display: flex;
    align-items: center;
    padding-left: 1rem;

    &_img {
      margin-right: 1rem;
      img {
        width: 9rem;
        height: 8rem;
        object-fit: cover;
      }
    }

    &_title {
      font-size: 1.5rem;
      color: gray;
    }
    &_color {
      color: gray;
    }
  }
}
</style>
