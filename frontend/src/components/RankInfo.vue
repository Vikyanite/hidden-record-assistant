<script setup lang="ts">

import {model} from "../../wailsjs/go/models";
import {useStore} from "vuex";
import Popper from "vue3-popper";

const store = useStore()

defineProps<{
  data: model.RankDetails,
}>()

</script>

<template>
  <div class="rank">
      <div class="rank_img">
        <h3 class="rank_title"> {{ data.queueType == "RANKED_SOLO_5x5" ? "单排 / 双排" : "灵活组排" }} </h3>
        <Popper placement="bottom" arrow hover>
          <img
              :src="data.tierOb.iconPath"
              :alt="data.tier"
          />
          <template #content>
            <div style="text-align: left">
              <p>
                {{ data.leaguePoints }}胜点
              </p>
              <p class="rank_color">
                  {{ data.wins }}胜 {{ data.losses }}负
              </p>
              <p class="rank_color" v-if="data.division!='NA'">
                胜率: {{((data.wins / (data.wins + data.losses)) * 100).toFixed(0) }}%
              </p>

            </div>
          </template>
        </Popper>
        <p class="rank_division">
          {{ data.tierOb.name }} {{ data.division!='NA' ? data.division : '' }}
        </p>
      </div>
  </div>

</template>

<style scoped lang="scss">
  .rank {
    &_img {
      img {
        width: 5rem;
        height: 5rem;
        margin: -10px 0 -15px 0;
        object-fit: cover;
      }
    }
    &_division {
      color: var(--color-win);
      font-size: 1rem;
    }
    &_title {
      font-size: 1.5rem;
      color: gray;
    }
    &_color {
      color: gray;
    }
  }

</style>