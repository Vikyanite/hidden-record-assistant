<template>
  <div class="recent">
    <div class="recent_title">
      <h4>近期对局情况</h4>
      <p>(最近 {{10}} 场游戏):</p>
    </div>

    <div class="recent_summary" style="padding:.5rem">
      <div class="recent_summary_total">
        <p>{{ data.wins }}胜 {{ data.defeats }}负
          ({{ (data.wins * 100 / (data.wins + data.defeats)).toFixed(0) }}%)</p>
        <pie-chart :data="[['Wins', data.wins], ['Defeats', data.defeats]]" :colors="['rgb(32,178,170)', 'rgb(240, 128, 128)']" width="100px" height="100px" :donut="true" :legend="false"></pie-chart>
        <div>
          <p>{{ (data.kills / (data.wins + data.defeats)).toFixed(1) }} /
            {{ (data.deaths / (data.wins + data.defeats)).toFixed(1) }} /
            {{ (data.assists / (data.wins + data.defeats)).toFixed(1) }}</p>
          <p>{{ ((data.kills + data.assists) / data.deaths).toFixed(2) }} KDA</p>
        </div>
      </div>

      <div class="recent_summary_lane">
        <p>近期常用位置:</p>
        <div class="recent_summary_lane_name">
          <img :src="store.getters.LocalAssetPrefix() + '/assets/images/lane/' + data.preferablyLane + '.png'" class="lane-icon" alt="lane">
          <p v-if="data.preferablyLane == 'utility'">{{data.preferablyLane}}
            ({{ (data.preferablyLaneGames * 100 / (data.wins + data.defeats)).toFixed(0) }}%)</p>
          <p v-else>{{ data.preferablyLane }}
            ({{ (data.preferablyLaneGames * 100 / (data.wins + data.defeats)).toFixed(0) }}%)</p>
        </div>
        <p>{{ data.wins_preferablyLane }}胜 {{ data.defeats_preferablyLane }}负
          ({{ (data.wins_preferablyLane * 100 / (data.preferablyLaneGames)).toFixed(1) }}%)</p>
        <p>{{ (data.kills_preferablyLane / data.preferablyLaneGames).toFixed(1) }} /
          {{ (data.deaths_preferablyLane / data.preferablyLaneGames).toFixed(1) }} /
          {{ (data.assists_preferablyLane / data.preferablyLaneGames).toFixed(1) }}
          ({{ ((data.kills_preferablyLane + data.assists_preferablyLane) / data.deaths_preferablyLane).toFixed(2) }}
          KDA)</p>
      </div>

      <div class="recent_summary_champs">
        <p>近期常用英雄:</p>
        <div class="recent_summary_champs_1">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            data.preferablyChamp1 +
            '.png'
          "
          alt="champ img"
          />
          <div>
            <p>{{ data.wins_preferablyChamps[0] }}胜 {{ data.defeats_preferablyChamps[0] }}负
              ({{ (data.wins_preferablyChamps[0] * 100 / (data.wins_preferablyChamps[0] + data.defeats_preferablyChamps[0])).toFixed(1) }}%)</p>
            <p>{{ (data.kills_preferablyChamps[0] / (data.wins_preferablyChamps[0] + data.defeats_preferablyChamps[0])).toFixed(1) }}
              /
              {{ (data.deaths_preferablyChamps[0] / (data.wins_preferablyChamps[0] + data.defeats_preferablyChamps[0])).toFixed(1) }}
              /
              {{ (data.assists_preferablyChamps[0] / (data.wins_preferablyChamps[0] + data.defeats_preferablyChamps[0])).toFixed(1) }}
              ({{ ((data.kills_preferablyChamps[0] + data.assists_preferablyChamps[0]) / data.deaths_preferablyChamps[0]).toFixed(2) }}
              KDA)</p>
          </div>
        </div>
        <div class="recent_summary_champs_2">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            data.preferablyChamp2 +
            '.png'
          "
          alt="champ img"
          />  
          <div>
            <p>{{ data.wins_preferablyChamps[1] }}胜 {{ data.defeats_preferablyChamps[1] }}负
              ({{ (data.wins_preferablyChamps[1] * 100 / (data.wins_preferablyChamps[1] + data.defeats_preferablyChamps[1])).toFixed(1) }}%)</p>
            <p>{{ (data.kills_preferablyChamps[1] / (data.wins_preferablyChamps[1] + data.defeats_preferablyChamps[1])).toFixed(1) }}
              /
              {{ (data.deaths_preferablyChamps[1] / (data.wins_preferablyChamps[1] + data.defeats_preferablyChamps[1])).toFixed(1) }}
              /
              {{ (data.assists_preferablyChamps[1] / (data.wins_preferablyChamps[1] + data.defeats_preferablyChamps[1])).toFixed(1) }}
              ({{ ((data.kills_preferablyChamps[1] + data.assists_preferablyChamps[1]) / data.deaths_preferablyChamps[1]).toFixed(2) }}
              KDA)</p>
          </div>       
        </div>
      </div>
    </div>


    <div class="recent_title">
      <h4>近期表现:</h4>
    </div>

    <div class="recent_overview">
      <div>
        <h3 class="name_yellow">{{ (data.firstBloodTimes * 100 / (data.wins + data.defeats)).toFixed(1) }}%</h3>
        <p>一血参与率 %</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (data.gold / (data.wins + data.defeats)).toFixed(2) }}</h3>
        <p>经济 / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (data.cs / (data.wins + data.defeats)).toFixed(2) }}</h3>
        <p>补刀 / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (data.control_wards / (data.wins + data.defeats)).toFixed(2) }}</h3>
        <p>视野控制 / game</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (data.vision_score / (data.wins + data.defeats)).toFixed(2) }}</h3>
        <p>视野得分 / game</p>
      </div>

    </div>
  </div>

</template>

<script lang="ts" setup>

import {useStore} from "vuex";
import {model} from "../../wailsjs/go/models";

const props = defineProps<{
  data: model.RecentMatchStatistic,
}>()

const store = useStore()

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.lane-icon {
  width: 4rem;
  height: 4rem;
  margin-right: .5rem;
}
.recent {
  height: auto;
  padding: .5rem;

  &_title{
    display: flex;
    margin-bottom: 1rem;

    h4{
      margin-right: .5rem;
    }
  }

  &_summary{
    text-align: center;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    justify-items: center;
    margin-bottom: 2rem;

    @media screen and (max-width: 650px){
      grid-template-columns: 1fr 1fr;
    }

    &>*{
      background-color: var(--color-remake01);
      border-radius: 5px;
      padding: .5rem;
      width: 95%;
    -webkit-box-shadow: 2px 2px 29px 6px rgba(0,0,0,0.25); 
      box-shadow: 2px 2px 29px 6px rgba(0,0,0,0.25);
    }

    &>*:not(:last-child) {
      margin-right: 1rem;
      @media screen and (max-width: 650px){
        margin-right: 0;
        margin-bottom: 2rem;
      }
    }

    &_total{
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
    }

    &_lane{

      &_name{
        text-transform: uppercase;
      }
      &>*:not(:last-child){
        margin-bottom: .5rem;
      }
    }

    &_champs{
      font-size: 1.3rem;
      //text-align: left;
      @media screen and (max-width: 650px){
        grid-column: 1/-1;
      }
      &_1, &_2{
        display: flex;
        align-items: center;
        justify-content: center;
        text-align: left;
        @media screen and (max-width: 650px){
          padding: 1rem;
        }
        


        img{
          width: 4rem;
          height: 4rem;
          margin-right: 1rem;
        }
      }
      &_1{
        margin-top: 1rem;
        margin-bottom: 1rem;
        @media screen and (max-width: 650px){
          float: left;
          margin: 0;
        }
        @media screen and (max-width: 450px){
          float: none;
        }
      }
      &_2{
        @media screen and (max-width: 650px){
          float: right;
        }
        @media screen and (max-width: 450px){
          float: none;
        }
      }
    }
  }

  &_overview{
    display: grid;
    grid-template-columns: 1fr  1fr 1fr max-content max-content;
    align-items: stretch;
    
    gap: 1rem;
    text-align: center;
  
    @media screen and (max-width: 600px){
      grid-template-columns: 1fr 1fr 1fr;
    }

    @media screen and (max-width: 500px){
      grid-template-columns: 1fr 1fr;
    }
    @media screen and (max-width: 350px){
      grid-template-columns: 1fr;
    }

    &>*{
      background-color: var(--color-remake01);
      padding: .5rem;
      border-radius: 5px;
      width: 100%;
     -webkit-box-shadow: 2px 2px 29px 6px rgba(0,0,0,0.25); 
      box-shadow: 2px 2px 29px 6px rgba(0,0,0,0.25);

      display: flex;
      align-items: center;
      flex-direction: column;
      justify-content: center;

    }
    &>*:not(:last-child) {
      margin-right: .5rem;
    }
  }
}
.name_yellow{
  color: var(--color-yellow);
}
</style>
