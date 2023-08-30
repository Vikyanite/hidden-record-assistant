<script setup lang="ts">
  import {onMounted} from "vue";
  import {CurrentSummoner} from "../../wailsjs/go/service/WailsApp";

  onMounted(() => {
    console.log("Personal Page Mounted");
    CurrentSummoner()
      .then((summoner) => {
        console.log("info:" + summoner);
      })
      .catch((err) => {
        console.log("err:" + err);
      })
  })
</script>

<template>
  <div class="myfont">
    Personal Page
  </div>

</template>

<style scoped>
.myfont {
  font-size: 50px;
  color: #333333;
}
</style>