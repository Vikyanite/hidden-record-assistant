<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {InitConnector} from "../../wailsjs/go/main/App";


const router = useRouter()
const text = ref("loadingPage")

onMounted(() => {
  InitConnector()
      .then(() => {
        text.value = "success"
        setTimeout(() => {
          router.push("/home")
        }, 1000)
        //router.push("/home")

      })
      .catch((err: any) => {text.value = err.toString()})
})

</script>

<template>
  <div>
    <div class="myfont">
      {{ text }}
    </div>
  </div>
</template>

<style scoped>
.myfont {
  font-size: 50px;
  color: #333333;
}
</style>