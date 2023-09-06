<template>
  <div class="recent" v-if="allMatches">
    <div class="recent_title">
      <h4>RECENT SUMMARY</h4> 
      <p>(last {{10}} games):</p>
    </div>

    <div class="recent_summary" style="padding:.5rem">
      <div class="recent_summary_total">
        <p>{{ statistic.wins }}W {{ statistic.defeats }}L
          ({{ (statistic.wins * 100 / (statistic.wins + statistic.defeats)).toFixed(0) }}%)</p>
        <pie-chart :data="[['Wins', statistic.wins], ['Defeats', statistic.defeats]]" :colors="['rgb(32,178,170)', 'rgb(240, 128, 128)']" width="100px" height="100px" :donut="true" :legend="false"></pie-chart>
        <div>
          <p>{{ (statistic.kills / (statistic.wins + statistic.defeats)).toFixed(1) }} /
            {{ (statistic.deaths / (statistic.wins + statistic.defeats)).toFixed(1) }} /
            {{ (statistic.assists / (statistic.wins + statistic.defeats)).toFixed(1) }}</p>
          <p>{{ ((statistic.kills + statistic.assists) / statistic.deaths).toFixed(2) }} KDA</p>
        </div>
      </div>

      <div class="recent_summary_lane">
        <p>Preferred Position:</p>
        <div class="recent_summary_lane_name">
          <img :src="store.getters.LocalAssetPrefix() + '/assets/images/lane/' + statistic.preferablyLane + '.png'" class="lane-icon" alt="lane">
          <p v-if="statistic.preferablyLane == 'utility'">Support
            ({{ (statistic.preferablyLaneGames * 100 / (statistic.wins + statistic.defeats)).toFixed(0) }}%)</p>
          <p v-else>{{ statistic.preferablyLane }}
            ({{ (statistic.preferablyLaneGames * 100 / (statistic.wins + statistic.defeats)).toFixed(0) }}%)</p>
        </div>
        <p>{{ statistic.wins_preferablyLane }}W {{ statistic.defeats_preferablyLane }}L
          ({{ (statistic.wins_preferablyLane * 100 / (statistic.preferablyLaneGames)).toFixed(1) }}%)</p>
        <p>{{ (statistic.kills_preferablyLane / statistic.preferablyLaneGames).toFixed(1) }} /
          {{ (statistic.deaths_preferablyLane / statistic.preferablyLaneGames).toFixed(1) }} /
          {{ (statistic.assists_preferablyLane / statistic.preferablyLaneGames).toFixed(1) }}
          ({{ ((statistic.kills_preferablyLane + statistic.assists_preferablyLane) / statistic.deaths_preferablyLane).toFixed(2) }}
          KDA)</p>
      </div>

      <div class="recent_summary_champs">
        <p>Preferred champions:</p>
        <div class="recent_summary_champs_1">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            statistic.preferablyChamp1 +
            '.png'
          "
          alt="champ img"
          />
          <div>
            <p>{{ statistic.wins_preferablyChamps[0] }}W {{ statistic.defeats_preferablyChamps[0] }}L
              ({{ (statistic.wins_preferablyChamps[0] * 100 / (statistic.wins_preferablyChamps[0] + statistic.defeats_preferablyChamps[0])).toFixed(1) }}%)</p>
            <p>{{ (statistic.kills_preferablyChamps[0] / (statistic.wins_preferablyChamps[0] + statistic.defeats_preferablyChamps[0])).toFixed(1) }}
              /
              {{ (statistic.deaths_preferablyChamps[0] / (statistic.wins_preferablyChamps[0] + statistic.defeats_preferablyChamps[0])).toFixed(1) }}
              /
              {{ (statistic.assists_preferablyChamps[0] / (statistic.wins_preferablyChamps[0] + statistic.defeats_preferablyChamps[0])).toFixed(1) }}
              ({{ ((statistic.kills_preferablyChamps[0] + statistic.assists_preferablyChamps[0]) / statistic.deaths_preferablyChamps[0]).toFixed(2) }}
              KDA)</p>
          </div>
        </div>
        <div class="recent_summary_champs_2">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            statistic.preferablyChamp2 +
            '.png'
          "
          alt="champ img"
          />  
          <div>
            <p>{{ statistic.wins_preferablyChamps[1] }}W {{ statistic.defeats_preferablyChamps[1] }}L
              ({{ (statistic.wins_preferablyChamps[1] * 100 / (statistic.wins_preferablyChamps[1] + statistic.defeats_preferablyChamps[1])).toFixed(1) }}%)</p>
            <p>{{ (statistic.kills_preferablyChamps[1] / (statistic.wins_preferablyChamps[1] + statistic.defeats_preferablyChamps[1])).toFixed(1) }}
              /
              {{ (statistic.deaths_preferablyChamps[1] / (statistic.wins_preferablyChamps[1] + statistic.defeats_preferablyChamps[1])).toFixed(1) }}
              /
              {{ (statistic.assists_preferablyChamps[1] / (statistic.wins_preferablyChamps[1] + statistic.defeats_preferablyChamps[1])).toFixed(1) }}
              ({{ ((statistic.kills_preferablyChamps[1] + statistic.assists_preferablyChamps[1]) / statistic.deaths_preferablyChamps[1]).toFixed(2) }}
              KDA)</p>
          </div>       
        </div>
      </div>
    </div>


    <div class="recent_title">
      <h4>PERFORMANCE OVERVIEW:</h4> 
    </div>

    <div class="recent_overview">
      <div>
        <h3 class="name_yellow">{{ (statistic.firstBloodTimes * 100 / (statistic.wins + statistic.defeats)).toFixed(1) }}%</h3>
        <p>First Blood %</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (statistic.gold / (statistic.wins + statistic.defeats)).toFixed(2) }}</h3>
        <p>Gold / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (statistic.cs / (statistic.wins + statistic.defeats)).toFixed(2) }}</h3>
        <p>CS / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (statistic.control_wards / (statistic.wins + statistic.defeats)).toFixed(2) }}</h3>
        <p>Control wards / game</p>
      </div>
      <div>
        <h3 class="name_yellow">{{ (statistic.vision_score / (statistic.wins + statistic.defeats)).toFixed(2) }}</h3>
        <p>Vision score / game</p>
      </div>

    </div>
  </div>

</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue";
import {IMatchStatistic, IMatchData} from "../types/match";
import {useStore} from "vuex";

const props = defineProps({
  allMatches: Object as () => IMatchData[],
  puuid: String,
})

onMounted(()=>{
  getData();
})

const store = useStore()

const statistic:IMatchStatistic = reactive({
  mainPlayerPoz:[], //array with all postion of main player in all matches
  kills:0,
  assists:0,
  deaths: 0,
  wins:0,
  defeats:0,

  gold:0,
  vision_score:0,
  firstBloodTimes:0,
  cs:0,
  control_wards:0,

  lane:[], //vector cu toate lane uri
  preferablyLane:null, //cel mai jucat lane
  preferablyLaneGames: 0, // nr meciurilor jucate pe cel mai jucat lane
  champs:[], //vector cu toti campionii
  preferablyChamp1:null, //cel mai jucat campion
  preferablyChamp2:null, //al doilea cel mai jucat campion,

  preferablyChamp1_games:0,
  preferablyChamp2_games:0,

  kills_preferablyLane:0, //statistici pe cel mai jucat lane
  deaths_preferablyLane:0,
  assists_preferablyLane:0,
  wins_preferablyLane: 0,
  defeats_preferablyLane:0,

  kills_preferablyChamps:[0,0], //statistici pt cei mai jucati campioni pe poz 0: cel mai jucat champ, poz 2: al 2 lea cel mai jucat campion
  deaths_preferablyChamps:[0,0],
  assists_preferablyChamps:[0,0],
  wins_preferablyChamps:[0,0],
  defeats_preferablyChamps:[0,0],
});

function getData(){
      let allMatches = props.allMatches || [];
      let puuid = props.puuid;
      
      //get main player poz array
      for(let i = 0; i<allMatches.length; i++){
        for(let j = 0; j < allMatches[i].participantIdentities.length; j++){
          if(puuid == allMatches[i].participantIdentities[j].player.puuid){
            statistic.mainPlayerPoz.push(j);
          }
        }
      }

      //get all kills, asssists and deaths

      for(let i = 0; i < allMatches.length; i++){
        statistic.kills += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.kills;
        statistic.deaths += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.deaths;
        statistic.assists += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.assists;
        //get wins, defeats
        allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.win ? statistic.wins++ : statistic.defeats++;

        //get lanes
        if(allMatches[i].participants[statistic.mainPlayerPoz[i]].timeline.lane){
          statistic.lane.push(allMatches[i].participants[statistic.mainPlayerPoz[i]].timeline.lane);
        }

        //get champs
        statistic.champs.push(allMatches[i].participants[statistic.mainPlayerPoz[i]].championId);

        //get gold
        statistic.gold += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.goldEarned / (allMatches[i].gameDuration/60);

        //get vision score
        statistic.vision_score += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.visionScore;

        //first blood kill or assists
        if(allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.firstBloodAssist || allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.firstBloodKill){
          statistic.firstBloodTimes++;
        }

        //get cs
        statistic.cs += (allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.neutralMinionsKilled + allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.totalMinionsKilled) / (allMatches[i].gameDuration/60);

        //get control wards
        statistic.control_wards += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.visionWardsBoughtInGame;

      }
      // console.log(matches.lane);
      // console.log(matches.champs);


      //get most played lane
      let nrmax=-100;
      for(let i = 0; i < statistic.lane.length; i++){
        var nr = 0, lane;
        for(let j = 0; j < statistic.lane.length; j++){
          if(statistic.lane[j] == statistic.lane[i]){
            nr++;
          }
        }
        if(nr>nrmax){
          nrmax = nr;
          statistic.preferablyLane = statistic.lane[i].toLowerCase();
        }
        
      }
      statistic.preferablyLaneGames = nrmax;


      //get most 2 played champs
      let max1=-100, max2=-100;
      for(let i = 0; i < statistic.champs.length; i++){
        let nr = 0;
        for(let j = 0; j < statistic.champs.length; j++){
          if(statistic.champs[j] == statistic.champs[i]){
            nr++;
          }
        }
        if(nr > max1){
          max1 = nr;
          statistic.preferablyChamp1 = statistic.champs[i];
        }
      }

      for(let i = 0; i < statistic.champs.length; i++){
        if(statistic.champs[i] == statistic.preferablyChamp1){
          statistic.champs.splice(i, 1);
        }
      }

      // console.log(matches.champs);

      for(let i = 0; i < statistic.champs.length; i++){
        let nr = 0;
        for(let j = 0; j < statistic.champs.length; j++){
          if(statistic.champs[j] == statistic.champs[i]){
            nr++;
          }
        }
        if(nr > max2){
          max2 = nr;
          statistic.preferablyChamp2 = statistic.champs[i];
        }
      }

      statistic.preferablyChamp1_games = max1;
      statistic.preferablyChamp2_games = max2;


      console.log(max1, statistic.preferablyChamp1);
      console.log(max2, statistic.preferablyChamp2);

      //get stats for most played lane and most played champs games
      for(let i = 0; i < allMatches.length; i++){
        if(allMatches[i].participants[statistic.mainPlayerPoz[i]].timeline.lane == statistic.preferablyLane?.toUpperCase()){
          statistic.kills_preferablyLane += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.kills;
          statistic.deaths_preferablyLane += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.deaths;
          statistic.assists_preferablyLane += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.assists;
          if(allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.win){
            statistic.wins_preferablyLane++;
          } else {
            statistic.defeats_preferablyLane++;
          }
        }
        if(allMatches[i].participants[statistic.mainPlayerPoz[i]].championId == statistic.preferablyChamp1){
          statistic.kills_preferablyChamps[0] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.kills;
          statistic.deaths_preferablyChamps[0] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.deaths;
          statistic.assists_preferablyChamps[0] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.assists;

          if(allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.win){
            statistic.wins_preferablyChamps[0]++;
          } else {
            statistic.defeats_preferablyChamps[0]++;
          }
        }

        if(allMatches[i].participants[statistic.mainPlayerPoz[i]].championId == statistic.preferablyChamp2){
          statistic.kills_preferablyChamps[1] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.kills;
          statistic.deaths_preferablyChamps[1] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.deaths;
          statistic.assists_preferablyChamps[1] += allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.assists;

          if(allMatches[i].participants[statistic.mainPlayerPoz[i]].stats.win){
            statistic.wins_preferablyChamps[1]++;
          } else {
            statistic.defeats_preferablyChamps[1]++;
          }
        }
      }

    }

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
