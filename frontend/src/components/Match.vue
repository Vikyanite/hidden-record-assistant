<template>
  <div class="match" style="display:flex; flex-direction:column;"> 
    <div style="width: 100%" :class="data.gameDurationInt<250 ? 'remake' : (data.participants[data.poz].stats.win ? 'win' : 'lose')">
      <div class="match_details">
        <div class="match_details_time">
          <p>{{ data.queueDescription }}</p>

          <p v-if="data.gameDurationInt<250" style="color: var(--color-remake)">Remake</p>
          <p v-else-if="data.participants[data.poz].stats.win" style="color: var(--color-win)">Win</p>
          <p v-else style="color: var(--color-lose)">Defeat</p>
          <p v-if="data.gameDurationInt">{{ data.gameDuration }}</p>
          <p>{{ data.gameTimeAgo }}前</p>
        </div>

        <div class="match_details_champ" >
<!--          <p v-if="data.realChampsNames[data.poz]">{{ data.realChampsNames[data.poz] }}</p>-->
          <img
            :src="store.getters.LCUAPIPrefix(data.participants[data.poz].championOb.squarePortraitPath)"
            alt="champ img"
          />
          <p class="match_details_champ_level">{{ data.participants[data.poz].stats.champLevel }}</p>
        </div>

        <div class="match_details_spells">
          <Popper placement="top" arrow hover>
            <img
              :src=" store.getters.LCUAPIPrefix(data.spellD.iconPath)"
              alt="spell img"
              class="match_details_spells_spellD"
            />

            <template #content>
              <div style="max-width: 400px">
                <p class="spell_name">{{ data.spellD.name }}</p>
                {{ data.spellD.description }}
              </div>
            </template>
          </Popper>

          <Popper placement="top" arrow hover>
            <img
              :src="store.getters.LCUAPIPrefix(data.runePrimary.iconPath)"
              alt="runes img"
              class="match_details_spells_runePrimary"
            />
            <template #content>
              <p class="spell_name">{{ data.runePrimary.name }}</p>
              <p
                style="max-width: 400px"
                v-html="data.runePrimary.longDesc"
              ></p>
            </template>
          </Popper>

          <Popper placement="top" arrow hover>
            <img
                referrerpolicy="no-referrer"
                :src=" store.getters.LCUAPIPrefix(data.spellF.iconPath)"
              alt="spell img"
              class="match_details_spells_spellF"
            />

            <template #content>
              <div style="max-width: 400px">
                <p class="spell_name">{{ data.spellF.name }}</p>
                {{ data.spellF.description }}
              </div>
            </template>
          </Popper>

          <Popper placement="top" arrow hover
                  v-if="data.runeSecondaryStyle">
            <img
              :src="store.getters.LCUAPIPrefix(data.runeSecondaryStyle.iconPath)"
              alt="runes img"
              class="match_details_spells_runeSecondary"
            />
            <template #content>
              <p
                style="max-width: 400px"
                v-html="data.runeSecondaryStyle.description"
              ></p>
            </template>
          </Popper>
        </div>

        <div class="match_details_score">
          <p>
            {{ data.participants[data.poz].stats.kills }} /
            {{ data.participants[data.poz].stats.deaths }} /
            {{ data.participants[data.poz].stats.assists }}
          </p>
          <p>
            KDA:
            {{
              (
                (data.participants[data.poz].stats.kills + data.participants[data.poz].stats.assists)
                  / max(1, data.participants[data.poz].stats.deaths)
              ).toFixed(2)
            }}
          </p>
        </div>

        <div class="match_details_level">
          <p>
            补刀(cs): {{
              data.participants[data.poz].stats.totalMinionsKilled + data.participants[data.poz].stats.neutralMinionsKilled
            }}
          </p>
          <p>{{ data.kp }}% 参团率</p>
        </div>

        <div class="match_items_container">
          <div class="match_items">
            <div
              v-for="i in ['item0Ob', 'item1Ob', 'item2Ob', 'item3Ob', 'item4Ob', 'item5Ob', 'item6Ob']"
              :key="i"
            >
              <div v-if="(data.participants[data.poz].stats as any)[i].id == 0">
                <Popper placement="top" arrow hover>
                  <img
                    src="https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/items/icons2d/gp_ui_placeholder.png"
                    alt="placeholder"
                  />

                  <template #content>
                    <p style="max-width: 400px">No item</p>
                  </template>
                </Popper>
              </div>
              <div v-else>
                <Popper placement="top" arrow hover>
                  <img
                    :src="store.getters.LCUAPIPrefix((data.participants[data.poz].stats as any)[i].iconPath)"
                    alt="item"
                  />
                  <template #content>
                    <div style="text-align: left">
                      <p class="spell_name">
                        {{
                          (data.participants[data.poz].stats as any)[i].name
                        }}
                      </p>
                      <p
                        style="max-width: 400px"
                        v-html="(data.participants[data.poz].stats as any)[i].description"
                      ></p>
                      <p class="spell_name">
                        Cost:
                        {{
                          (data.participants[data.poz].stats as any)[i].priceTotal
                        }}
                      </p>
                    </div>
                  </template>
                </Popper>
              </div>
            </div>
          </div>
          <p class="match_items_vs">
            {{ data.participants[data.poz].stats.visionScore }} 视野得分
          </p>
        </div>

        <div
          class="match_btn"
          :class="data.gameDurationInt<250 ? 'remake_btn' : (data.participants[data.poz].stats.win ? 'win_btn' : 'lose_btn')"
          @click="data.toggleAdvancedDetails = !data.toggleAdvancedDetails"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            aria-hidden="true"
            role="img"
            width="1em"
            height="1em"
            preserveAspectRatio="xMidYMid meet"
            viewBox="0 0 24 24"
          >
            <path
              d="M16.59 8.59L12 13.17L7.41 8.59L6 10l6 6l6-6z"
              fill="currentColor"
            />
          </svg>
        </div>

      </div>

      <div v-if="data.toggleAdvancedDetails">
        <MatchDetail :matchInfo="data"></MatchDetail>
      </div>

    </div>

  </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import MatchDetail from "./MatchDetail.vue";
import Popper from "vue3-popper";

import { useStore } from 'vuex';

import {model} from "../../wailsjs/go/models";
import {max} from "@floating-ui/utils";

const store = useStore();

const props = defineProps<{
  data: model.DisplayMatch,
}>()

</script>


<style scoped lang="scss">
.match {
  padding: 0.5rem 1rem 0.5rem 1rem;
  display: flex;
  text-align: center;
  font-size: 1.2rem;

  img {
    border-radius: 3px;
    object-fit: cover;
  }
  &_details {
    height: 100% !important;
    width: 100% !important;
    display: grid;
    grid-template-columns: 13% min-content min-content max-content max-content min-content 10% 10% min-content;
    align-items: center;
    justify-content: space-between;

    &_time {
      display: flex;
      flex-direction: column;
    }

    &_champ {
      img {
        width: 6rem;
        height: 6rem;
      }
    }

    &_spells {
      display: grid;
      grid-template-columns: 1fr 1fr;
      img {
        width: 2.5rem;
        height: 2.5rem;
        margin-right: 0.5rem;
        margin-top: 0.5rem;
        cursor: pointer;
      }

      &_runePrimary {
        background-color: black;
        border-radius: 50% !important;
      }
    }
    .spell_name {
      color: var(--color-yellow);
      font-weight: bold;
    }
    &_score {
      display: flex;
      flex-direction: column;
    }
  }

  &_items {
    display: grid;
    grid-template-columns: repeat(4, max-content);
    justify-items: center;
    align-items: center;
    img {
      width: 3rem;
      height: 3rem;
      margin-right: 0.3rem;
    }

    &_container {
      display: flex;
      flex-direction: column;
    }
  }

  &_participanti {
    display: flex;
    flex-direction: column;
    width: 100%;

    a {
      text-decoration: none;
      color: inherit;
      overflow: hidden;

      p {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }

    img {
      width: 2rem;
      height: 2rem;
      object-fit: cover;
      margin-right: 0.5rem;
    }
    &_name {
      display: flex;
      align-items: center;
      margin-right: 0.3rem;
      margin-bottom: 0.3rem;
    }
  }

  &_btn {
    height: 100%;
    padding: 0.5rem;
    cursor: pointer;

    display: flex;
    align-items: center;
    justify-content: center;

    svg {
      width: 2rem;
      height: 2rem;
    }
  }
}
.win {
  background-image: linear-gradient(
          to left,
          var(--color-win07),
          var(--color-win01)
  );
  border-left: 1px solid var(--color-win);
  height: auto;
  padding: 0.5rem;
}
.lose {
  background-image: linear-gradient(
          to left,
          var(--color-lose07),
          var(--color-lose01)
  );
  border-left: 1px solid var(--color-lose);
  height: auto;
  padding: 0.5rem;
}
.remake {
  background-image: linear-gradient(
          to left,
          var(--color-remake07),
          var(--color-remake01)
  );
  border-left: 1px solid var(--color-remake);
  height: auto;
  padding: 0.5rem;
}
.remake_btn {
  background-color: var(--color-remake07);
}

.win_btn {
  background-color: var(--color-win07);
}
.lose_btn {
  background-color: var(--color-lose07);
}
.tooltip_style {
  max-width: 400px;
}
</style>
