<template>
  <div class="match" style="display:flex; flex-direction:column;"> 
    <div v-if="data.poz != null" style="width: 100%" :class="data.gameDurationInt<250 ? 'remake' : (data.participants[data.poz].stats.win ? 'win' : 'lose')">
      <div class="match_details">

        <div class="match_details_time">
          <p>{{ data.queueDescription }}</p>

          <p v-if="data.gameDurationInt<250" style="color: var(--color-remake)">Remake</p>
          <p v-else-if="data.participants[data.poz].stats.win" style="color: var(--color-win)">Win</p>
          <p v-else style="color: var(--color-lose)">Defeat</p>
          <p v-if="data.gameDurationInt">{{ data.gameDuration }}</p>
          <p>{{ data.gameTimeAgo }} ago</p>
        </div>

        <div class="match_details_champ" >
          <p v-if="data.realChampsNames[data.poz]">{{ data.realChampsNames[data.poz] }}</p>
          <img
            :src="
              'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
              data.participants[data.poz].championId +
              '.png'
            "
            alt="champ img"
          />
        </div>

        <div class="match_details_spells">
          <Popper placement="top" arrow hover>
            <img
              :src=" store.getters.LCUAPIPrefix() + data.spellD.iconPath"
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
                :src="store.getters.LCUAPIPrefix() + data.runePrimary.iconPath"
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
                :src=" store.getters.LCUAPIPrefix() + data.spellF.iconPath"
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

              :src="store.getters.LCUAPIPrefix() + data.runeSecondaryStyle.iconPath"
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
                (data.participants[data.poz].stats.kills +
                    data.participants[data.poz].stats.assists) /
                data.participants[data.poz].stats.deaths
              ).toFixed(2)
            }}
          </p>
        </div>

        <div class="match_details_level">
          <p>Level {{ data.participants[data.poz].stats.champLevel }}</p>
          <p>
            {{
              data.participants[data.poz].stats.totalMinionsKilled +
              data.participants[data.poz].stats.neutralMinionsKilled
            }}
            CS
          </p>
          <p>
            {{
              (
                (data.participants[data.poz].stats.totalMinionsKilled +
                    data.participants[data.poz].stats.neutralMinionsKilled)
                / (data.gameDurationInt / 60)
              ).toFixed(1)
            }}
            CS/MIN
          </p>
          <p>{{ data.kp }}% KP</p>
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
                    :src="store.getters.LCUAPIPrefix() + (data.participants[data.poz].stats as any)[i].iconPath"
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
            {{ data.participants[data.poz].stats.visionScore }} vision
            score
          </p>
        </div>

        <div class="match_participanti">
          <div v-for="(i, index) in data.participants" :key="index">
            <div v-if="index < 5" class="match_participanti_name">
              <img
                :src="
                  'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
                  data.participants[index].championId +
                  '.png'
                "
                alt="champ img"
              />
              <a
                :href="
                  '/' +
                  data.participantIdentities[index].player.summonerName
                "
                ><p>
                  {{ data.participantIdentities[index].player.summonerName }}
                </p>
              </a>
            </div>
          </div>
        </div>

        <div class="match_participanti">
          <div v-for="(i, index) in data.participants" :key="index">
            <div v-if="index >= 5" class="match_participanti_name">
              <img
                :src="
                  'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
                  data.participants[index].championId +
                  '.png'
                "
                alt="champ img"
              />
              <a
                :href="
                  '/' +
                  data.participantIdentities[index].player.summonerName
                "
                ><p>{{ data.participantIdentities[index].player.summonerName }}</p>
              </a>
            </div>
          </div>
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

      <div v-if="data.toggleAdvancedDetails && data.realChampsNames">
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

const store = useStore();

const props = defineProps<{
  data: model.DisplayMatch,
}>()

</script>


<style scoped lang="scss">
@import "../styles/match_details.scss"; // search_page styles
</style>
