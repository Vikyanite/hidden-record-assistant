<script setup lang="ts">

import routes from "../routes/router";
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<template>
  <div>
    <el-container>
      <el-aside width="200px">
        <el-menu
            default-active="/home/personal"
            class="el-menu-vertical-demo"
            mode="vertical"
            router
        >
          <el-menu-item index="/home/personal">
            <span>个人信息</span>
          </el-menu-item>
          <el-menu-item index="/home/test1">
            <span>test1</span>
          </el-menu-item>
          <el-menu-item index="/home/test2">
            <span>test2</span>
          </el-menu-item>

        </el-menu>
      </el-aside>
      <el-container>
        <!--      <el-header>Header</el-header>-->
        <el-main>
          <router-view v-slot="{ Component }">
            <transition appear>
              <keep-alive>
                <component :is="Component"/>
              </keep-alive>
            </transition>
          </router-view>
        </el-main>
        <!--      <el-footer>Footer</el-footer>-->
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.myfont {
  font-size: 50px;
  color: #333333;
}
</style>