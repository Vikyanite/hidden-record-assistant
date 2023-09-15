<template>
  <div>
    <div class="navbar">
      <button
          v-for="section in toggle"
          :key="section"
          class="navbar_btn"
          @click="setActiveSection(section)"
      >
        {{ section }}
      </button>
    </div>

    <div v-if="activeSection === 'Overview'">
      <overview  :matchData="matchInfo" :data="matchInfo.overview"></overview>
    </div>

    <div v-else-if="activeSection === 'Advanced'">
      <advanced :matchData="matchInfo"></advanced>
    </div>

<!--    <div v-else-if="activeSection === 'Breakdown'">-->
<!--      <breakdown :stats="matchInfo.breakdown"></breakdown>-->
<!--    </div>-->

      <!-- <p style="width:200px;">{{this.matchInfo}}</p>  -->
  </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue";

import overview from "../components/overview.vue";
import advanced from "../components/advanced.vue";
import breakdown from "../components/breakdown.vue";
import {model} from "../../wailsjs/go/models";

  const props = defineProps<{
    matchInfo: model.DisplayMatch,
  }>()

  const toggle = ["Overview",
    "Advanced",
    // "Breakdown"
  ];

  const activeSection = ref("Overview");

  function setActiveSection(section: string) {
    activeSection.value = section;
  }

</script>

<style lang="scss" scoped>
.navbar {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(3,1fr);
  align-items: center;
  gap: .5rem;
  justify-content: space-between;
  margin-top: 1rem;

  &_btn {
    width: 100%;
    background-color: transparent;
    border: 1px solid lightgrey;
    outline: none;
    color: #fff;
    font-size: 1.3rem;
    font-family: inherit;
    cursor: pointer;
    transition: all .2s;


    &:hover {
      transform: translateY(-2px);
    }
    &:active {
      transform: translateY(2px);
    }
  }
}

</style>