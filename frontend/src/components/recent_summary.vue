<template>
  <div class="recent" v-if="allMatches">
    <div class="recent_title">
      <h4>RECENT SUMMARY</h4> 
      <p>(last {{10}} games):</p>
    </div>

    <div class="recent_summary" style="padding:.5rem">
      <div class="recent_summary_total">
        <p>{{matches.wins}}W {{matches.defeats}}L ({{(matches.wins*100/(matches.wins + matches.defeats)).toFixed(0)}}%)</p>
        <pie-chart :data="[['Wins', matches.wins], ['Defeats', matches.defeats]]" :colors="['rgb(32,178,170)', 'rgb(240, 128, 128)']" width="100px" height="100px" :donut="true" :legend="false"></pie-chart>
        <div>
          <p>{{(matches.kills / (matches.wins + matches.defeats)).toFixed(1)}} / {{(matches.deaths / (matches.wins + matches.defeats)).toFixed(1)}} / {{(matches.assists / (matches.wins + matches.defeats)).toFixed(1)}}</p>
          <p>{{((matches.kills + matches.assists) / matches.deaths).toFixed(2)}} KDA</p>
        </div>
      </div>

      <div class="recent_summary_lane">
        <p>Preferred Position:</p>
        <div class="recent_summary_lane_name">
          <img :src="'../assets/images/lane/' + matches.preferablyLane + '.png'" class="lane-icon" alt="lane">
          <p v-if="matches.preferablyLane == 'utility'">Support ({{(matches.preferablyLaneGames * 100 / (matches.wins + matches.defeats)).toFixed(0)}}%)</p>
          <p v-else>{{matches.preferablyLane}} ({{(matches.preferablyLaneGames * 100 / (matches.wins + matches.defeats)).toFixed(0)}}%)</p>
        </div>
        <p>{{matches.wins_preferablyLane}}W {{matches.defeats_preferablyLane}}L ({{(matches.wins_preferablyLane * 100 / (matches.preferablyLaneGames)).toFixed(1)}}%)</p>
        <p>{{(matches.kills_preferablyLane / matches.preferablyLaneGames).toFixed(1)}} / {{(matches.deaths_preferablyLane / matches.preferablyLaneGames).toFixed(1)}} / {{(matches.assists_preferablyLane / matches.preferablyLaneGames).toFixed(1)}}  ({{((matches.kills_preferablyLane + matches.assists_preferablyLane)/matches.deaths_preferablyLane).toFixed(2)}} KDA)</p>
      </div>

      <div class="recent_summary_champs">
        <p>Preferred champions:</p>
        <div class="recent_summary_champs_1">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            matches.preferablyChamp1 +
            '.png'
          "
          alt="champ img"
          />
          <div>
            <p>{{matches.wins_preferablyChamps[0]}}W {{matches.defeats_preferablyChamps[0]}}L ({{(matches.wins_preferablyChamps[0] * 100 / (matches.wins_preferablyChamps[0] + matches.defeats_preferablyChamps[0])).toFixed(1)}}%)</p>
            <p>{{(matches.kills_preferablyChamps[0] / (matches.wins_preferablyChamps[0] + matches.defeats_preferablyChamps[0])).toFixed(1)}} / {{(matches.deaths_preferablyChamps[0] / (matches.wins_preferablyChamps[0] + matches.defeats_preferablyChamps[0])).toFixed(1)}} / {{(matches.assists_preferablyChamps[0] / (matches.wins_preferablyChamps[0] + matches.defeats_preferablyChamps[0])).toFixed(1)}} ({{((matches.kills_preferablyChamps[0] + matches.assists_preferablyChamps[0]) / matches.deaths_preferablyChamps[0]).toFixed(2)}} KDA)</p>
          </div>
        </div>
        <div class="recent_summary_champs_2">
          <img
          :src="
            'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' +
            matches.preferablyChamp2 +
            '.png'
          "
          alt="champ img"
          />  
          <div>
            <p>{{matches.wins_preferablyChamps[1]}}W {{matches.defeats_preferablyChamps[1]}}L ({{(matches.wins_preferablyChamps[1] * 100 / (matches.wins_preferablyChamps[1] + matches.defeats_preferablyChamps[1])).toFixed(1)}}%)</p>
            <p>{{(matches.kills_preferablyChamps[1] / (matches.wins_preferablyChamps[1] + matches.defeats_preferablyChamps[1])).toFixed(1)}} / {{(matches.deaths_preferablyChamps[1] / (matches.wins_preferablyChamps[1] + matches.defeats_preferablyChamps[1])).toFixed(1)}} / {{(matches.assists_preferablyChamps[1] / (matches.wins_preferablyChamps[1] + matches.defeats_preferablyChamps[1])).toFixed(1)}} ({{((matches.kills_preferablyChamps[1] + matches.assists_preferablyChamps[1]) / matches.deaths_preferablyChamps[1]).toFixed(2)}} KDA)</p>
          </div>       
        </div>
      </div>
    </div>


    <div class="recent_title">
      <h4>PERFORMANCE OVERVIEW:</h4> 
    </div>

    <div class="recent_overview">
      <div>
        <h3 class="name_yellow">{{(matches.firstBloodTimes * 100 / (matches.wins + matches.defeats)).toFixed(1)}}%</h3>
        <p>First Blood %</p>
      </div>
      <div>
        <h3 class="name_yellow">{{(matches.gold / (matches.wins + matches.defeats)).toFixed(2)}}</h3>
        <p>Gold / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{(matches.cs / (matches.wins + matches.defeats)).toFixed(2)}}</h3>
        <p>CS / min</p>
      </div>
      <div>
        <h3 class="name_yellow">{{(matches.control_wards / (matches.wins + matches.defeats)).toFixed(2)}}</h3>
        <p>Control wards / game</p>
      </div>
      <div>
        <h3 class="name_yellow">{{(matches.vision_score / (matches.wins + matches.defeats)).toFixed(2)}}</h3>
        <p>Vision score / game</p>
      </div>

    </div>
  </div>

</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue";
import {IMatchData} from "../types/match";


  const props = defineProps({
    allMatches: {
      type: Object as () => IMatchData[],
    }, //toate meciurile pt a le face analiza
    puuid: String,
  })

  onMounted(()=>{
    getData();
  })

interface displayMatches{
  mainPlayerPoz: number[]; // array with all positions of main player in all matches
  kills: number;
  assists: number;
  deaths: number;
  wins: number;
  defeats: number;

  gold: number;
  vision_score: number;
  firstBloodTimes: number;
  cs: number;
  control_wards: number;

  role: string[]; // vector cu toate lane-urile
  preferablyLane: string | null; // cel mai jucat lane
  preferablyLaneGames: number; // nr meciurilor jucate pe cel mai jucat lane
  champs: number[]; // vector cu toți campionii
  preferablyChamp1: number | null; // cel mai jucat campion
  preferablyChamp2: number | null; // al doilea cel mai jucat campion

  preferablyChamp1_games: number;
  preferablyChamp2_games: number;

  kills_preferablyLane: number; // statistici pe cel mai jucat lane
  deaths_preferablyLane: number;
  assists_preferablyLane: number;
  wins_preferablyLane: number;
  defeats_preferablyLane: number;

  kills_preferablyChamps: [number, number]; // statistici pentru cei mai jucati campioni [0: cel mai jucat champ, 1: al 2-lea cel mai jucat campion]
  deaths_preferablyChamps: [number, number];
  assists_preferablyChamps: [number, number];
  wins_preferablyChamps: [number, number];
  defeats_preferablyChamps: [number, number];
}

const matches:displayMatches = reactive({
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

  role:[], //vector cu toate lane uri
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
            matches.mainPlayerPoz.push(j);
          }
        }
      }

      //get all kills, asssists and deaths

      for(let i = 0; i < allMatches.length; i++){
        matches.kills += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.kills;
        matches.deaths += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.deaths;
        matches.assists += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.assists;
        //get wins, defeats
        allMatches[i].participants[matches.mainPlayerPoz[i]].stats.win ? matches.wins++ : matches.defeats++;

        //get lanes
        if(allMatches[i].participants[matches.mainPlayerPoz[i]].timeline.role){
          matches.role.push(allMatches[i].participants[matches.mainPlayerPoz[i]].timeline.role);
        }

        //get champs
        matches.champs.push(allMatches[i].participants[matches.mainPlayerPoz[i]].championId);

        //get gold
        matches.gold += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.goldEarned / (allMatches[i].gameDuration/60);

        //get vision score
        matches.vision_score += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.visionScore;

        //first blood kill or assists
        if(allMatches[i].participants[matches.mainPlayerPoz[i]].stats.firstBloodAssist || allMatches[i].participants[matches.mainPlayerPoz[i]].stats.firstBloodKill){
          matches.firstBloodTimes++;
        }

        //get cs
        matches.cs += (allMatches[i].participants[matches.mainPlayerPoz[i]].stats.neutralMinionsKilled + allMatches[i].participants[matches.mainPlayerPoz[i]].stats.totalMinionsKilled) / (allMatches[i].gameDuration/60);

        //get control wards
        matches.control_wards += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.visionWardsBoughtInGame;

      }
      // console.log(matches.lane);
      // console.log(matches.champs);


      //get most played lane
      let nrmax=-100;
      for(let i = 0; i < matches.role.length; i++){
        var nr = 0, lane;
        for(let j = 0; j < matches.role.length; j++){
          if(matches.role[j] == matches.role[i]){
            nr++;
          }
        }
        if(nr>nrmax){
          nrmax = nr;
          matches.preferablyLane = matches.role[i].toLowerCase();
        }
        
      }
      matches.preferablyLaneGames = nrmax;


      //get most 2 played champs
      let max1=-100, max2=-100;
      for(let i = 0; i < matches.champs.length; i++){
        let nr = 0;
        for(let j = 0; j < matches.champs.length; j++){
          if(matches.champs[j] == matches.champs[i]){
            nr++;
          }
        }
        if(nr > max1){
          max1 = nr;
          matches.preferablyChamp1 = matches.champs[i];
        }
      }

      for(let i = 0; i < matches.champs.length; i++){ 
        if(matches.champs[i] == matches.preferablyChamp1){
          matches.champs.splice(i, 1);
        }
      }

      // console.log(matches.champs);

      for(let i = 0; i < matches.champs.length; i++){
        let nr = 0;
        for(let j = 0; j < matches.champs.length; j++){
          if(matches.champs[j] == matches.champs[i]){
            nr++;
          }
        }
        if(nr > max2){
          max2 = nr;
          matches.preferablyChamp2 = matches.champs[i];
        }
      }

      matches.preferablyChamp1_games = max1;
      matches.preferablyChamp2_games = max2;


      console.log(max1, matches.preferablyChamp1);
      console.log(max2, matches.preferablyChamp2);

      //get stats for most played lane and most played champs games
      for(let i = 0; i < allMatches.length; i++){
        if(allMatches[i].participants[matches.mainPlayerPoz[i]].timeline.role == matches.preferablyLane?.toUpperCase()){
          matches.kills_preferablyLane += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.kills;
          matches.deaths_preferablyLane += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.deaths;
          matches.assists_preferablyLane += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.assists;
          if(allMatches[i].participants[matches.mainPlayerPoz[i]].stats.win){
            matches.wins_preferablyLane++;
          } else {
            matches.defeats_preferablyLane++;
          }
        }
        if(allMatches[i].participants[matches.mainPlayerPoz[i]].championId == matches.preferablyChamp1){
          matches.kills_preferablyChamps[0] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.kills;
          matches.deaths_preferablyChamps[0] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.deaths;
          matches.assists_preferablyChamps[0] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.assists;

          if(allMatches[i].participants[matches.mainPlayerPoz[i]].stats.win){
            matches.wins_preferablyChamps[0]++;
          } else {
            matches.defeats_preferablyChamps[0]++;
          }
        }

        if(allMatches[i].participants[matches.mainPlayerPoz[i]].championId == matches.preferablyChamp2){
          matches.kills_preferablyChamps[1] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.kills;
          matches.deaths_preferablyChamps[1] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.deaths;
          matches.assists_preferablyChamps[1] += allMatches[i].participants[matches.mainPlayerPoz[i]].stats.assists;

          if(allMatches[i].participants[matches.mainPlayerPoz[i]].stats.win){
            matches.wins_preferablyChamps[1]++;
          } else {
            matches.defeats_preferablyChamps[1]++;
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
