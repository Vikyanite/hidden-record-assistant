<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
const isCollapse = ref(true)

const router = useRouter()
onMounted(()=> {
  // 在页面加载时手动触发路由导航到 "/home/personal"
  router.push("/home/personal");
})

</script>

<template>
  <div>
    <el-container>
      <el-aside class="el-aside">
        <el-menu
            class="el-menu"
            @mouseenter="isCollapse = false"
            @mouseleave="isCollapse = true"
            default-active="/home/personal"
            mode="vertical"
            router
            :collapse="isCollapse"
        >
          <el-menu-item index="/home/personal">
            <el-icon><User /></el-icon>
            <span>个人信息</span>
          </el-menu-item>
          <el-menu-item index="/home/test1">
            <el-icon><Search /></el-icon>
            <span>搜索召唤师</span>
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
.el-menu{
  height:100vh
}
.el-aside {
  width:200px;
}

</style>