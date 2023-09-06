<template>
  <div class="match" style="display:flex; flex-direction:column;"> 
    <div v-if="mainPlayer.poz != null" style="width: 100%" :class="match.gameDuration<250 ? 'remake' : (match.participants[mainPlayer.poz].stats.win ? 'win' : 'lose')">
      <div class="match_details">

        <div class="match_details_time">
          <p>{{ mainPlayer.queue }}</p>

          <p v-if="match.gameDuration<250" style="color: var(--color-remake)">Remake</p>
          <p v-else-if="match.participants[mainPlayer.poz].stats.win" style="color: var(--color-win)">Win</p>
          <p v-else style="color: var(--color-lose)">Defeat</p>
          <p v-if="mainPlayer.gameDuration">{{ mainPlayer.gameDuration }}</p>
          <p>{{ mainPlayer.gameTimeAgo }} ago</p>
        </div>

        <div class="match_details_champ" >
          <p v-if="mainPlayer.realChampsNames[mainPlayer.poz]">{{ mainPlayer.realChampsNames[mainPlayer.poz] }}</p>
          <img
            :src="
              'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
              match.participants[mainPlayer.poz].championId +
              '.png'
            "
            alt="champ img"
          />
        </div>

        <div class="match_details_spells">
          <Popper placement="top" arrow hover>
            <img
              :src=" store.getters.LCUAPIPrefix() + mainPlayer.spellD.iconPath"
              alt="spell img"
              class="match_details_spells_spellD"
            />

            <template #content>
              <div style="max-width: 400px">
                <p class="spell_name">{{ mainPlayer.spellD.name }}</p>
                {{ mainPlayer.spellD.description }}
              </div>
            </template>
          </Popper>

          <Popper placement="top" arrow hover>
            <img
                :src="store.getters.LCUAPIPrefix() + mainPlayer.runePrimary.iconPath"
              alt="runes img"
              class="match_details_spells_runePrimary"
            />
            <template #content>
              <p class="spell_name">{{ mainPlayer.runePrimary.name }}</p>
              <p
                style="max-width: 400px"
                v-html="mainPlayer.runePrimary.longDesc"
              ></p>
            </template>
          </Popper>

          <Popper placement="top" arrow hover>
            <img
                referrerpolicy="no-referrer"
                :src=" store.getters.LCUAPIPrefix() + mainPlayer.spellF.iconPath"
              alt="spell img"
              class="match_details_spells_spellF"
            />

            <template #content>
              <div style="max-width: 400px">
                <p class="spell_name">{{ mainPlayer.spellF.name }}</p>
                {{ mainPlayer.spellF.description }}
              </div>
            </template>
          </Popper>

          <Popper placement="top" arrow hover
                  v-if="mainPlayer.runeSecondaryStyle">
            <img

              :src="store.getters.LCUAPIPrefix() + styleMap[mainPlayer.runeSecondaryStyle].iconPath"
              alt="runes img"
              class="match_details_spells_runeSecondary"
            />
            <template #content>
              <p
                style="max-width: 400px"
                v-html="styleMap[mainPlayer.runeSecondaryStyle].description"
              ></p>
            </template>
          </Popper>
        </div>

        <div class="match_details_score">
          <p>
            {{ match.participants[mainPlayer.poz].stats.kills }} /
            {{ match.participants[mainPlayer.poz].stats.deaths }} /
            {{ match.participants[mainPlayer.poz].stats.assists }}
          </p>
          <p>
            KDA:
            {{
              (
                (match.participants[mainPlayer.poz].stats.kills +
                  match.participants[mainPlayer.poz].stats.assists) /
                match.participants[mainPlayer.poz].stats.deaths
              ).toFixed(2)
            }}
          </p>
        </div>

        <div class="match_details_level">
          <p>Level {{ match.participants[mainPlayer.poz].stats.champLevel }}</p>
          <p>
            {{
              match.participants[mainPlayer.poz].stats.totalMinionsKilled +
              match.participants[mainPlayer.poz].stats.neutralMinionsKilled
            }}
            CS
          </p>
          <p>
            {{
              (
                (match.participants[mainPlayer.poz].stats.totalMinionsKilled +
                  match.participants[mainPlayer.poz].stats.neutralMinionsKilled)
                / (match.gameDuration / 60)
              ).toFixed(1)
            }}
            CS/MIN
          </p>
          <p>{{ mainPlayer.kp }}% KP</p>
        </div>

        <div class="match_items_container">
          <div class="match_items">
            <div
              v-for="i in [
                'item0',
                'item1',
                'item2',
                'item3',
                'item4',
                'item5',
                'item6',
              ]"
              :key="i"
            >
              <div v-if="(match.participants[mainPlayer.poz].stats as any)[i] == 0">
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

              <div v-else-if="((match.participants[mainPlayer.poz].stats as any)[i]>=7000 )&& ((match.participants[mainPlayer.poz].stats as any)[i]<=7025)">
                <Popper placement="top" arrow hover>
                  <img
                    src="https://raw.communitydragon.org/latest/game/assets/items/itemmodifiers/bordertreatmentornn.png"
                    alt="placeholder ornn item"
                  />

                  <template #content>
                    <p style="max-width: 400px">Ornn upgraded item</p>
                  </template>
                </Popper>
              </div>

              <div v-else>
                <Popper placement="top" arrow hover>
                  <img
                    :src="store.getters.LCUAPIPrefix() + store.state.items[(match.participants[mainPlayer.poz].stats as any)[i]].iconPath
                    "
                    alt="item"
                  />
                  <template #content>
                    <div style="text-align: left">
                      <p class="spell_name">
                        {{
                          store.state.items[(match.participants[mainPlayer.poz].stats as any)[i]].name
                        }}
                      </p>
                      <p
                        style="max-width: 400px"
                        v-html="store.state.items[(match.participants[mainPlayer.poz].stats as any)[i]].description"
                      ></p>
                      <p class="spell_name">
                        Cost:
                        {{
                          store.state.items[
                              (match.participants[mainPlayer.poz].stats as any)[i]
                          ].priceTotal
                        }}
                      </p>
                    </div>
                  </template>
                </Popper>
              </div>
            </div>
          </div>
          <p class="match_items_vs">
            {{ match.participants[mainPlayer.poz].stats.visionScore }} vision
            score
          </p>
        </div>

        <div class="match_participanti">
          <div v-for="(i, index) in match.participants" :key="index">
            <div v-if="index < 5" class="match_participanti_name">
              <img
                :src="
                  'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
                  match.participants[index].championId +
                  '.png'
                "
                alt="champ img"
              />
              <a
                :href="
                  '/' +
                  match.participantIdentities[index].player.summonerName
                "
                ><p>
                  {{ match.participantIdentities[index].player.summonerName }}
                </p>
              </a>
            </div>
          </div>
        </div>

        <div class="match_participanti">
          <div v-for="(i, index) in match.participants" :key="index">
            <div v-if="index >= 5" class="match_participanti_name">
              <img
                :src="
                  'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
                  match.participants[index].championId +
                  '.png'
                "
                alt="champ img"
              />
              <a
                :href="
                  '/' +
                  match.participantIdentities[index].player.summonerName
                "
                ><p>{{ match.participantIdentities[index].player.summonerName }}</p>
              </a>
            </div>
          </div>
        </div>

        <div
          class="match_btn"
          :class="match.gameDuration<250 ? 'remake_btn' : (match.participants[mainPlayer.poz].stats.win ? 'win_btn' : 'lose_btn')"
          @click="mainPlayer.toggleAdvancedDetails = !mainPlayer.toggleAdvancedDetails"
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

<!--      <div v-if="mainPlayer.toggleAdvancedDetails && mainPlayer.realChampsNames">-->
<!--        <match_history_advancedDetails_nav :matchInfo="match" :itemsInfo="itemsJson" :spellsInfo="spellsJson" :runesInfo="runesJson" :regionInfo="region" :continentInfo="continent" :puuid="summonersPuuid" :champsData="mainPlayer.realChampsNames"></match_history_advancedDetails_nav>-->
<!--      </div>-->

    </div>

  </div>
</template>

<script lang="ts" setup>
import {onMounted, reactive} from "vue";
import Popper from "vue3-popper";

import {IDisplayMatch, IMatchData} from "../types/match";
import { useStore } from 'vuex';

const store = useStore();

const props = defineProps({
  match: {
    type: Object as()=> IMatchData,
    required: true,
  },  //datele despre meci (un singur meci)
  summonersPuuid: String, //puuid ul player ului cautat
})

const styleMap: { [key: number]: { description: string; iconPath: string; } } = {
  8000: {
    description: "精密",
    iconPath:  "/lol-game-data/assets/v1/perk-images/Styles/7201_precision.png",
  },
  8100: {
    description: "主宰",
    iconPath: "/lol-game-data/assets/v1/perk-images/Styles/7200_domination.png",
  },
  8200: {
    description: "巫术",
    iconPath: "/lol-game-data/assets/v1/perk-images/Styles/7202_sorcery.png",
  },
  8300: {
    description: "启迪",
    iconPath: "/lol-game-data/assets/v1/perk-images/Styles/7203_whimsy.png",
  },
  8400: {
    description: "坚决",
    iconPath: "/lol-game-data/assets/v1/perk-images/Styles/7204_resolve.png",
  },
};

const mainPlayer:IDisplayMatch = reactive({
  poz: 0, //pozitia in array a player ului cautat
  queue: null, //queue ul fiecarui meci
  gameDuration: "",
  gameTimeAgo: "",
  kp: "",
  spellD: {
    id: 0,
    name: "",
    description: "",
    summonerLevel: 0,
    cooldown: 0,
    iconPath: "",
    gameModes: [""],
  },
  spellF: {
    id: 0,
    name: "",
    description: "",
    summonerLevel: 0,
    cooldown: 0,
    iconPath: "",
    gameModes: [""],
  },
  runePrimary: {
    id: 0,
    name: "",
    majorChangePatchVersion: "",
    tooltip: "",
    shortDesc: "",
    longDesc: "",
    recommendationDescriptor: "",
    iconPath: "",
    endOfGameStatDescs: [""],
    recommendationDescriptorAttributes: {},
  }, // pe pozitia 0 am runa in sine, pe poz 1 am descrierea, pe poz 2 am numele
  runeSecondaryStyle: 0,
  toggleAdvancedDetails: false,
  realChampsNames: [],  //un array cu toate numele reale ale campionilor, pt ca unii nu corespund ex: MonkeyKing - Wukong
})

onMounted(()=> {
  getMainPlayer();
  getTime();
  msToTime(new Date().getTime() - props.match.gameCreation);
  kp();
  getRealChampsName();
})

function getRealChampsName(){
  for(let i = 0; i < props.match.participants.length; i++){
    let name = store.state.champions[props.match.participants[i].championId].name
    mainPlayer.realChampsNames.push(name);
  }
}

function getMainPlayer() {
  for (let i = 0; i < props.match.participantIdentities.length; i++) {
    //pozitia in array ul cu meci al player ului cautat
    if (props.summonersPuuid == props.match.participantIdentities[i].player.puuid) {
      mainPlayer.poz = i;
      break
    }
  }

  mainPlayer.queue = store.state.queues[props.match.queueId].description;

  mainPlayer.spellD = store.state.spells[props.match.participants[mainPlayer.poz].spell1Id]
  mainPlayer.spellF = store.state.spells[props.match.participants[mainPlayer.poz].spell2Id]

  mainPlayer.runePrimary = store.state.perks[props.match.participants[mainPlayer.poz].stats.perk0]
  mainPlayer.runeSecondaryStyle = props.match.participants[mainPlayer.poz].stats.perkSubStyle
}

function getTime() {
  //sec to min:sec

  const sec = props.match.gameDuration
  let minutes = Math.floor(sec / 60); // get minutes
  let seconds = sec - minutes * 60; //  get seconds



  mainPlayer.gameDuration = `${zeroPad(minutes, 2)}min ${zeroPad(seconds, 2)}s`;
}

function zeroPad(num: number, width: number): string {
  const numStr = num.toString();
  const diff = width - numStr.length;

  if (diff <= 0) {
    return numStr;
  }

  return '0'.repeat(diff) + numStr;
}

function msToTime(duration:number) {
  // ms to ANYTIME
  let minutes = Math.floor((duration / (1000 * 60)) % 60)
  let hours = Math.floor((duration / (1000 * 60 * 60)) % 24)
  let days = Math.floor(duration / (24 * 60 * 60 * 1000))

  if (duration < 3600000) {
    mainPlayer.gameTimeAgo = zeroPad(minutes, 2) + " minutes";
  } else if (duration <= 86399999) {
    mainPlayer.gameTimeAgo =
      hours + " hours" /* + minutes + ":" + seconds*/;
  } else {
    if (duration > 86399999 && duration <= 172799999) {
      mainPlayer.gameTimeAgo = days + " day";
    } else {
      mainPlayer.gameTimeAgo = days + " days";
    }
  }
}

function kp() {
  let teamkills = 0;
  for (let i = 0; i < props.match.participants.length; i++) {
    if (
      props.match.participants[mainPlayer.poz].teamId ==
      props.match.participants[i].teamId
    ) {
      teamkills += props.match.participants[i].stats.kills;
    }
  }
  mainPlayer.kp = (
    ((props.match.participants[mainPlayer.poz].stats.kills +
      props.match.participants[mainPlayer.poz].stats.assists) /
      teamkills) *
    100
  ).toFixed(0);
}

</script>


<style scoped lang="scss">
@import "../styles/match_details.scss"; // search_page styles
</style>
