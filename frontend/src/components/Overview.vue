<template>
  <div class="team">
    <div v-for="(i, index) in matchData.teams">
      <div :class="['team_overview', index === 0 ? 'team_blue_overview' : 'team_red_overview']">
        <div class="team_overview_result">
          <p v-if="matchData.teams[index].win=='Win'" class="win">Win</p>
          <p v-else class="lose">Defeat</p>
          <p>{{ index==0 ? "(Blue Team)" : "(Red Team)" }}</p>
        </div>

        <div class="team_overview_score">
          <p v-if="index==0">{{data.teamBlue.kills}} / {{data.teamBlue.deaths}} / {{data.teamBlue.assists}}</p>
          <p v-else>{{data.teamRed.kills}} / {{data.teamRed.deaths}} / {{data.teamRed.assists}}</p>
        </div>

        <div class="team_overview_obj">
          <Popper placement="top" arrow hover>
            <div class="team_overview_obj_item">
              <img v-if="index==0" src="../assets/images/other/tower-100.png" alt="tower svg">
              <img v-else src="../assets/images/other/tower-200.png" alt="tower svg">
              <p>{{matchData.teams[index].towerKills}}</p>
            </div>
            <template #content>
              <p>Towers destroyed</p>
            </template>
          </Popper>

          <Popper placement="top" arrow hover>
            <div class="team_overview_obj_item">
              <img v-if="index==0" src="../assets/images/other/inhibitor-100.png" alt="inhibitor svg">
              <img v-else src="../assets/images/other/inhibitor-200.png" alt="inhibitor svg">
              <p>{{matchData.teams[index].inhibitorKills}}</p>
            </div>
            <template #content>
              <p>Inhibitors destroyed</p>
            </template>
          </Popper>
        </div>

        <div class="team_overview_monsters">
          <Popper placement="top" arrow hover>
            <div class="team_overview_obj_item">
              <img v-if="index==0" src="../assets/images/other/baron-100.png" alt="baron svg">
              <img v-else src="../assets/images/other/baron-200.png" alt="baron svg">
              <p>{{matchData.teams[index].baronKills}}</p>
            </div>
            <template #content>
              <p>Barons killed</p>
            </template>
          </Popper>
          <Popper placement="top" arrow hover>
            <div class="team_overview_obj_item">
              <img v-if="index==0" src="../assets/images/other/herald-100.png" alt="herald svg">
              <img v-else src="../assets/images/other/herald-200.png" alt="herald svg">
              <p>{{matchData.teams[index].riftHeraldKills}}</p>
            </div>
            <template #content>
              <p>Heralds killed</p>
            </template>
          </Popper>
          <Popper placement="top" arrow hover>
            <div class="team_overview_obj_item">
              <img v-if="index==0" src="../assets/images/other/dragon-100.png" alt="dragon svg">
              <img v-else src="../assets/images/other/dragon-200.png" alt="dragon svg">

              <p>{{matchData.teams[0].dragonKills}}</p>
            </div>
            <template #content>
              <p>Dragons killed</p>
            </template>
          </Popper>
        </div>

        <Popper placement="top" arrow hover>
          <div class="team_overview_gold">
            <img src="../assets/images/other/icon_gold.png" alt="gold">
            <p v-if="index==0">{{(data.teamBlue.gold/1000).toFixed(1)}}K</p>
            <p v-else>{{(data.teamRed.gold/1000).toFixed(1)}}K</p>
          </div>
          <template #content>
            <p>Total team gold</p>
          </template>
        </Popper>

      </div>

      <div :class="index == 0 ? 'team_blue' : 'team_red'">
        <div v-for="(i, secIndex) in matchData.participants" :key="secIndex">
          <div class="team_player" v-if="index == 0 && secIndex < 5 || index == 1 && secIndex >=5 ">
            <Popper placement="top" arrow hover>
              <div class="team_player_level">
                <img :src="matchData.participants[secIndex].championOb.squarePortraitPath" alt="lol stats img">
                <p>{{matchData.participants[secIndex].stats.champLevel}}</p>
              </div>
              <template #content>
                <div style="max-width: 400px">
                  <p>{{matchData.participants[secIndex].championOb.name}}</p>
                </div>
              </template>
            </Popper>

            <div class="team_player_spells">
              <Popper placement="top" arrow hover>
                <img
                    :src="data.spellDs[secIndex].iconPath"
                    alt="spell img"
                />

                <template #content>
                  <div style="max-width: 400px">
                    <p class="spell_name">{{ data.spellDs[secIndex].name }}</p>
                    {{ data.spellDs[secIndex].description }}
                  </div>
                </template>
              </Popper>

              <Popper placement="top" arrow hover>
                <img
                    :src="data.spellFs[secIndex].iconPath"
                    alt="spell img"
                />

                <template #content>
                  <div style="max-width: 400px">
                    <p class="spell_name">{{ data.spellFs[secIndex].name }}</p>
                    {{ data.spellFs[secIndex].description }}
                  </div>
                </template>
              </Popper>
            </div>

            <div class="team_player_name">
              <a :href="'/' + matchData.participantIdentities[secIndex].player.summonerName" target="_blank">
                <p>{{matchData.participantIdentities[secIndex].player.summonerName}}</p>
              </a>
            </div>

            <div class="team_player_score">
              <p>{{((matchData.participants[secIndex].stats.kills + matchData.participants[secIndex].stats.assists)/matchData.participants[secIndex].stats.deaths).toFixed(1)}} KDA</p>
              <div class="team_player_score_kp">
                <p>{{matchData.participants[secIndex].stats.kills}}/{{matchData.participants[secIndex].stats.deaths}}/{{matchData.participants[secIndex].stats.assists}}</p>
                <p>({{(((matchData.participants[secIndex].stats.kills + matchData.participants[secIndex].stats.assists)/data.teamBlue.kills)*100).toFixed(0)}}% KP)</p>
              </div>
            </div>

            <div class="team_player_cs">
              <p>{{matchData.participants[secIndex].stats.totalMinionsKilled + matchData.participants[secIndex].stats.neutralMinionsKilled}} CS</p>
              <p>{{((matchData.participants[secIndex].stats.totalMinionsKilled + matchData.participants[secIndex].stats.neutralMinionsKilled) / (matchData.gameDurationInt / 60)).toFixed(1)}}/min</p>
            </div>

            <div class="team_player_dmg">
              <p><span>{{((matchData.participants[secIndex].stats.totalDamageDealtToChampions)/1000).toFixed(3)}}</span> <span>DMG</span></p>
            </div>

            <div class="team_player_items">
              <div v-for="j in ['item0Ob','item1Ob','item2Ob','item3Ob','item4Ob','item5Ob','item6Ob',]" :key="j">
                <div v-if="(matchData.participants[secIndex].stats as any)[j].id == 0">
                  <Popper placement="top" arrow hover>
                    <img
                        src="../assets/images/other/gp_ui_placeholder.png"
                        alt="placeholder"
                    />

                    <template #content>
                      <p style="max-width: 400px">No item {{(matchData.participants[secIndex].stats as any)[j].id}}</p>
                    </template>
                  </Popper>
                </div>
                <div v-else>
                  <Popper placement="top" arrow hover>
                    <img
                        :src="(matchData.participants[secIndex].stats as any)[j].iconPath"
                        alt="item"
                    />
                    <template #content>
                      <div style="text-align: left">
                        <p class="spell_name">
                          {{ (matchData.participants[secIndex].stats as any)[j].name }}
                        </p>
                        <p
                            style="max-width: 400px"
                            v-html="(matchData.participants[secIndex].stats as any)[j].description"
                        ></p>
                        <p class="spell_name">
                          Cost: {{ (matchData.participants[secIndex].stats as any)[j].priceTotal }}
                        </p>
                      </div>
                    </template>
                  </Popper>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="team_bans">
        <div>
            <div v-for="(i,index) in matchData.teams[0].bans" :key="index">
                <img style="filter: grayscale(100%);" :src="matchData.teams[0].bans[index].championOb.squarePortraitPath" alt="champ img"/>
            </div>
        </div>

        <p>: Bans :</p>

        <div>
            <div v-for="(i, index) in matchData.teams[1].bans" :key="index">
                <img style="filter: grayscale(100%);" :src="matchData.teams[1].bans[index].championOb.squarePortraitPath" alt="champ img"/>
            </div>
        </div>
    </div>

  </div>
</template>

<script setup lang="ts">
  import Popper from "vue3-popper";
  import {model} from "../../wailsjs/go/models";
  import {useStore} from "vuex";

  const props = defineProps<{
    data: model.Overview,
    matchData: model.DisplayMatch,
  }>()

  const store = useStore();

</script>

<style lang="scss" scoped>
.spell_name {
    color: var(--color-yellow);
}
.win {
    color: var(--color-win);
    font-weight: bold;
    font-size: 1.4rem;
    margin-right: .5rem;
}
.lose {
    color: var(--color-lose);
    font-weight: bold;
    font-size: 1.4rem;
    margin-right: .5rem;
}
.team {
    margin-top: 1rem;
    display: grid;
    grid-template-columns: 1fr;
    gap: 1rem;
    width: 100%;

    // white-space: nowrap;
        // overflow-x: scroll;
        // overflow-y: hidden;

    &_overview {
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-weight: bold;

        &_result {
            display: flex;
            align-items: center;
            font-size: 1.4rem;
        }

        &_obj {
            display: flex;
            align-items: center;

            &_item {
                display: flex;
                align-items: center;
                margin-right: .7rem;
                &>* {
                    margin-right: .3rem;
                }
            }
        }

        &_monsters {
            display: flex;
            align-items: center;
        }
        img {
            width: 2rem !important;
            height: 2rem !important;

        }

        &_gold {
            display: flex;
            align-items: center;

            img {
                width: max-content !important;
                object-fit: cover;
                margin-right: .5rem;
            }
        }
    }

    &_bans {
        display: flex;
        align-items: center;
        justify-content: space-between;
        border-top: 4px solid rgba(0,0,0, 0.2);
        padding-top: 1rem;

        div {
            display: flex;
            &>*:not(:last-child) {
                margin-right: .5rem;
            }
        }
    }


    img {
        width: 3rem;
        height: 3rem;
        display: flex;
        border-radius: .4rem;
        object-fit: cover;
    }

    &_blue,&_red {
        display: flex;
        flex-direction: column;
        width: 100% !important;

    }

    &_blue {
        border-bottom: 4px solid rgba(0,0,0, 0.2);
        // border-top: 2px solid rgba(0,0,0, 0.2);
    }

    &_player {
        width: 100% !important;
        display: grid;
        grid-template-columns: min-content min-content min-content 13% min-content 1fr 1fr min-content;
        align-items: center;
        justify-items: center;
        text-align: left;
        gap: 1rem;
        padding-bottom: .5rem;
        padding-top: .5rem;
        width: 100%;


        transition: all .2s;
        border-bottom: 1px solid rgba(0,0,0,.1);

        &:hover {
            background-color: rgba(0,0,0,.1);
        }

        &_name {
            justify-self: flex-start;
            text-decoration: none;
            cursor: pointer;
            width: 100%;
            a {
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                text-decoration: none;
                color: inherit;    
                p{
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                }                                                           
            }
        }

        &_level {
            position: relative;

            p {
                position: absolute;
                bottom: -.7rem;
                right: -.5rem;
                background-color: gray;
                border-radius: 50%;
                width: 1.5rem;
                height: 1.7rem;
                display: flex;
                align-items: center;
                justify-content: center;
                z-index: 5;
            }
        }

        &_dmg {
            white-space: nowrap;

        }

        &_score {
            display: flex;
            flex-direction: column;
            align-items: center;
            white-space: nowrap;

            &_kp {
                display: flex;
                align-items: center;

                &>*:not(:last-child) {
                    margin-right: .7rem;
                }
            }


        }
        &_cs {
            display: flex;
            flex-direction: column;
            align-items: center;
        }


        &_items {
           display: grid;
        //    grid-template-columns: repeat(4,1fr);
            grid-template-columns: repeat(7,1fr);
           gap: .1rem;
           row-gap: 0;

           img {
               width: 2.7rem !important; 
               height: 2.7rem !important;
           }

        }

        &_spells {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;

            img {
                width: 1.7rem;
                height: 1.7rem;
            }
        }

        &_runes {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 1rem;

            &_items {
                &>* {
                    margin-bottom: 1rem;
                }
            }
        }

        &_rune_box {
            display: flex;
            align-items: flex-start;
            img {
                margin-right: .5rem;
            }
        }
        &_statsRunes {
            display: flex;
        }

        &_runeBtn {
            background-color: transparent;
            border: 1px solid white;
            font-family: inherit;
            color: white;
            padding: .4rem;
            cursor: pointer;

 
        }
    }
}


</style>