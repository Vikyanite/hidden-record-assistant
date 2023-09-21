<template>
  <div class="breakdown">
    
      <div v-if="stats.mvp" class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.mvp.championOb.id + '/' + stats.mvp.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <h2 style="color: var(--color-yellow); margin-bottom: 0;">MVP</h2>
            <h2>{{stats.mvp.value}} score</h2>
            <div class="breakdown_box_text_name">
              <img :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/'+ stats.mvp.championOb.id + '.png'" alt="lol_img">
              <p>{{stats.mvp.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>

      <div v-if="stats.highestDMG" class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.highestDMG.championOb.id + '/' + stats.highestDMG.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <p>Highest DMG</p>
            <h2>{{stats.highestDMG.value/1000}}</h2>
            <div class="breakdown_box_text_name">
              <img :src="stats.highestDMG.championOb.squarePortraitPath" alt="lol_img">
              <p>{{stats.highestDMG.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>

      <div v-if="stats.highestCS" class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.highestCS.championOb.id + '/' + stats.highestCS.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <p>Highest CS</p>
            <h2>{{stats.highestCS.value}}</h2>
            <div class="breakdown_box_text_name">
              <img :src="stats.highestCS.championOb.squarePortraitPath" alt="lol_img">
              <p>{{stats.highestCS.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>

      <div v-if="stats.bestGold" class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.bestGold.championOb.id + '/' + stats.bestGold.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <p>Best Gold</p>
            <h2>{{stats.bestGold.value/1000}}</h2>
            <div class="breakdown_box_text_name">
              <img :src="stats.bestGold.championOb.squarePortraitPath" alt="lol_img">
              <p>{{stats.bestGold.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>


      <div v-if="stats.bestVS" class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.bestVS.championOb.id + '/' + stats.bestVS.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <p>Best Vision Score</p>
            <h2>{{stats.bestVS.value}}</h2>
            <div class="breakdown_box_text_name">
              <img :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/'+ stats.bestVS.championOb.id + '.png'" alt="lol_img">
              <p>{{stats.bestVS.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>

      <div class="breakdown_box">
        <div class="breakdown_box_bgImg">
          <img class="breakdown_box_bgImg_img" :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-splashes/uncentered/' + stats.highestKDA.championOb.id + '/' + stats.highestKDA.championOb.id + '000.jpg'" alt="lol_img">

          <div class="breakdown_box_overlay">
      
          </div>

          <div class="breakdown_box_text">
            <p>Highest KDA</p>
            <h2>{{stats.highestKDA.value}}</h2>
            <div class="breakdown_box_text_name">
              <img :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/'+ stats.highestKDA.championOb.id + '.png'" alt="lol_img">
              <p>{{stats.highestKDA.summonerName}}</p>
            </div>
          </div>
        </div>
      </div>

  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { model } from "../../wailsjs/go/models";

 const props = defineProps<{
   stats: model.Breakdown,
 }>()

</script>

<style lang="scss" scoped>
.breakdown{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;

  margin-top: 0;
  padding: 1rem;



  &_box {
    &_bgImg {
      position: relative;

      &_img {
        width: 100%;
        height: 17rem;
        border-radius: 5px;
        object-fit: cover;
      }
    }

    &_overlay {
      position: absolute;
      top: 0;
      left: 0;
      background-image: linear-gradient(
        to left,
        rgba(0,0,0,0.2),
        rgba(0,0,0,0.7)
      );
      // opacity: 0.4;
      width: 100%;
      height: 96.5%;
      border-radius: 5px;
      border: 1px solid rgba(0,0,0,0.5);
    }

    &_text {
      position: absolute;
      top: 50%;
      left: 0%;
      transform: translate(10%, -60%);
      text-align: left;

      &>* {
        margin-bottom: .5rem;
      }

      &_name {
        display: flex;
        align-items: center;

        img {
          border-radius: 50%;
          width: 3rem;
          height: 3rem;
          margin-right: .5rem;
        }
      }
    }
  }
}



</style>