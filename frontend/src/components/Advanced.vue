<template>
<div class="advanced_container">
  <div class="team_btn">
    <button class="blue_team" style="background-color: var(--color-win)">blue team</button>
    <button class="red_team" style="background-color: var(--color-lose)">red team</button>
  </div>

  <p class="team_name" style="color:var(--color-win)">Blue Team:</p>


  <div class="advanced" style="margin-top:2rem;">

    <div class="advanced_name">
     <p></p>
    </div>

    <div class="advanced_champ_list" v-for="(i, index) in matchData.participants" :key="index">
      <div class="advanced_champ_box">
        <img :src="store.getters.LCUAPIPrefix(matchData.participants[index].championOb.squarePortraitPath)" alt="lol stats img">
        <p style="white-space: nowrap;">{{matchData.participants[index].championOb.name}}</p>
      </div>
    </div>
  </div>


  <div class="advanced_title advanced_title_combat">
    <p>Combat</p>
    <svg class="advanced_svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z" fill="currentColor"/></svg>
  </div>
  <div class="advanced advanced_combat advanced2">

    <div class="advanced_name">
     <p>KDA</p>
     <p>Largest Killing Spree</p>
     <p>Largest Multi Kill</p>
     <p>First Blood</p>
    </div>
  
    <div class="advanced_champ_list" v-for="(i,index) in matchData.participants" :key="index">
      <p>{{matchData.participants[index].stats.kills}}/{{matchData.participants[index].stats.deaths}}/{{matchData.participants[index].stats.assists}}</p>
      <p>{{matchData.participants[index].stats.largestKillingSpree}}</p>
      <p>{{matchData.participants[index].stats.largestMultiKill}}</p>
      <p>{{(matchData.participants[index].stats.firstBloodKill)}}</p>
    </div>
  </div>


  <div class="advanced_title advanced_title_dmgDealt">
    <p>Damage Dealt</p>
    <svg class="advanced_svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z" fill="currentColor"/></svg>
  </div>
  <div class="advanced advanced_dmgDealt advanced2">

    <div class="advanced_name">
     <p>Total Damage to Champions</p>
     <p>Psysical Damage to Champions</p>
     <p>Magic Damage to Champions</p>
     <p>True Damage to Champions</p>
     <p>Total Damage Dealt</p>
     <p>Psyhical Damage Dealt</p>
     <p>Magic Damage Dealt</p>
     <p>True Damage Dealt</p>
     <p>Largest Critical Strike</p>
     <p>Total Damage to Objectives</p>
     <p>Total Damage to Turrets</p>
    </div>
  
    <div class="advanced_champ_list" v-for="(i,index) in matchData.participants" :key="index">
      <p>{{(matchData.participants[index].stats.totalDamageDealtToChampions/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.physicalDamageDealtToChampions/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.magicDamageDealtToChampions/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.trueDamageDealtToChampions/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.totalDamageDealt/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.physicalDamageDealt/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.magicDamageDealt/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.trueDamageDealt/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.largestCriticalStrike/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.damageDealtToObjectives/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.damageDealtToTurrets/1000).toFixed(1)}}K</p>
    </div>
  </div>


  <div class="advanced_title advanced_title_dmgTaken">
    <p>Damage Taken and Healed</p>
    <svg class="advanced_svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z" fill="currentColor"/></svg>
  </div>
  <div class="advanced advanced_dmgTaken advanced2">
    <div class="advanced_name">
     <p>Damage Healed</p>
     <p>Damage Taken</p>
     <p>Psysical Damage Taken</p>
     <p>Magic Damage Taken</p>
     <p>True Damage Taken</p>
    </div>
  
    <div class="advanced_champ_list" v-for="(i,index) in matchData.participants" :key="index">
      <p>{{(matchData.participants[index].stats.totalHeal/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.totalDamageTaken/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.physicalDamageTaken/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.magicalDamageTaken/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.trueDamageTaken/1000).toFixed(1)}}K</p>
    </div>

  </div>


  <div class="advanced_title advanced_title_vision">
    <p>Vision</p>
    <svg class="advanced_svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z" fill="currentColor"/></svg>
  </div>
  <div class="advanced advanced_vision advanced2">

    <div class="advanced_name">
    <p>Vision Score</p>
     <p>Wards Placed</p>
     <p>Wards Destroyed</p>
     <p>Stealth Wards Purchased</p>
     <p>Control Wards Purchased</p>
    </div>
  
    <div class="advanced_champ_list" v-for="(i,index) in matchData.participants" :key="index">
      <p>{{matchData.participants[index].stats.visionScore}}</p>
      <p>{{(matchData.participants[index].stats.wardsPlaced)}}</p>
      <p>{{(matchData.participants[index].stats.wardsKilled)}}</p>
      <p>{{(matchData.participants[index].stats.sightWardsBoughtInGame)}}</p>
      <p>{{(matchData.participants[index].stats.visionWardsBoughtInGame)}}</p>
    </div>

  </div>

  <div class="advanced_title advanced_title_income">
    <p>Income</p>
    <svg class="advanced_svg" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 1024 1024"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z" fill="currentColor"/></svg>
  </div>
  <div class="advanced advanced_income advanced2">

    <div class="advanced_name">
     <p>Total Gold Earned</p>
     <p>Total Gold Spent</p>
     <p>Minions Killed</p>
     <p>Neutral Minions Killed</p>
    </div>
  
    <div class="advanced_champ_list" v-for="(i,index) in matchData.participants" :key="index">
      <p>{{(matchData.participants[index].stats.goldEarned/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.goldSpent/1000).toFixed(1)}}K</p>
      <p>{{(matchData.participants[index].stats.totalMinionsKilled)}}</p>
      <p>{{(matchData.participants[index].stats.neutralMinionsKilled)}}</p>
    </div>



  </div>
</div>


</template>

<script setup lang="ts">
import {onMounted, reactive} from "vue";
import {model} from "../../wailsjs/go/models";
import {useStore} from "vuex";
onMounted(() => {
  hide_unhide();
  switchTeamsResponsive();
})

  const store = useStore();
  const props = defineProps<{
    matchData: model.DisplayMatch,
  }>()

    function hide_unhide() {
      const titles = document.querySelectorAll('.advanced_title');
      const data = document.querySelectorAll('.advanced2');
      const svg = document.querySelectorAll('.advanced_svg');

      for(let i = 0; i < titles.length; i++) {
        titles[i].addEventListener('click',function() {
          data[i].classList.toggle('hidden');
          svg[i].classList.toggle('rotate');
        });
      }
    }

    function switchTeamsResponsive() {
      const blue_btn = document.querySelector('.blue_team')!;
      const red_btn = document.querySelector('.red_team')!;

      const players: HTMLElement[][] = [];

      for (let i = 2; i <= 11; i++) {
        const player = Array.from(document.querySelectorAll(`.advanced>*:nth-child(${i})`) as NodeListOf<HTMLElement>);
        players.push(player);
      }


      blue_btn.addEventListener('click', function() {
        for(let i = 0; i < 5; i++) {
          players[i].forEach((player) => {
            player.style.display="grid";
          });
        }
        for (let i = 5; i < 10; i++) {
          players[i].forEach((player) => {
            player.style.display="none";
          });
        }

        document.querySelector('.team_name')!.innerHTML = '<p class="team_name" style="color:var(--color-win)">Blue Team:</p>'
      });

      red_btn.addEventListener('click', function() {
        for(let i = 0; i < 5; i++) {
          players[i].forEach((player) => {
            player.style.display="none";
          });
        }
        for (let i = 5; i < 10; i++) {
          players[i].forEach((player) => {
            player.style.display="grid";
          });
        }
        document.querySelector('.team_name')!.innerHTML = '<p class="team_name" style="color:var(--color-lose)">Red Team:</p>'
      });
    }



</script>

<style lang="scss" scoped>
.team_btn {
  display: none;

  @media screen and (max-width: 650px) {
    display:flex;
    align-items: center;
    justify-content: center;
    margin-top: 2rem;

    &>*:not(:last-child) {
      margin-right: 1rem;
    }

    &>* {
      border: 1px solid rgba(0,0,0,.1);
      background-color: lightgray;
      font-family: inherit;
      text-transform: uppercase;
      padding: .5rem;
      cursor: pointer;
      border-radius: 4px;
    }
  }
}
.team_name {
  margin-top: 1rem;
  font-size: 1.6rem;

  display: none;

  @media screen and (max-width: 650px) {
    display: block;
  }
}
.rotate{
  transform: rotate(-90deg);
}
.hidden {
  display: none !important;
}
.advanced2{
  transition: all .5s;
}
.advanced_svg {
  transition: all .2s;
}
.advanced {
  display: grid;
  grid-template-columns: repeat(11,9.09%);
  justify-items: stretch;
  align-items: stretch;
  border-top: solid rgba(0,0,0,.1);

  @media screen and (max-width: 650px) {
    grid-template-columns: repeat(6,16.66%);
    &>*:nth-child(7),&>*:nth-child(8),&>*:nth-child(9),&>*:nth-child(10),&>*:nth-child(11) {
      display: none;
    }
  }

  &_container {
    
    @media screen and (max-width: 650px) {
      -webkit-touch-callout: none; /* iOS Safari */
      -webkit-user-select: none; /* Safari */
      -khtml-user-select: none; /* Konqueror HTML */
      -moz-user-select: none; /* Old versions of Firefox */
      -ms-user-select: none; /* Internet Explorer/Edge */
      user-select: none;

      font-size: 1rem;
    }
  }



  &_title {
    text-transform: uppercase;
    height: 5rem;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    width: 100% !important;
    text-align: center;
    cursor: pointer;
    p{
      margin-right: .5rem;
    }
  }

  &_name {
    display: flex;
    flex-direction: column;
    font-size: 1rem;
    color: lightgray;

    &>* {
      height: 5.5rem;
      display: flex;
      align-items: center;
      justify-content: center;
      border-bottom: 1px solid rgba(0,0,0,.1);
      border-right: 1px solid rgba(0,0,0,.1);
      border-left: 1px solid rgba(0,0,0,.1);
      padding: .3rem;
    }
  }

  &_champ{

    &_box {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      padding: 1rem;
      img {
        width: 3rem !important;
        height: 3rem !important;
        margin-top: .7rem;
      }
      p {
        //margin-bottom: 1rem;
      }
    }

    &_list {
      display: flex;
      flex-direction: column;
      &>* {
        height: 5.5rem;
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px solid rgba(0,0,0,.1);
        border-right: 1px solid rgba(0,0,0,.1);
      }

      transition: all .2s;

      &:hover {
        background-color: rgba(0,0,0,.1);
      }
    }
  }
}

</style>