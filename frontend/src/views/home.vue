<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {Expand, Fold, User} from "@element-plus/icons-vue";
import {useStore} from "vuex";

const isCollapse = ref(true)
const router = useRouter()
const store = useStore()

onMounted(()=> {
  // 在页面加载时手动触发路由导航到 "/home/personal"
  router.push("/home/personal");
})

</script>

<template>
  <div>
    <el-container>
      <el-aside>
        <el-menu
            default-active="/home/personal"
            mode="vertical"
            router
            :collapse="isCollapse"
        >
          <el-menu-item @click="isCollapse=!isCollapse" class="menu-collapse-icon">
            <el-icon v-if="isCollapse"><Expand /></el-icon>
            <el-icon v-else><Fold /></el-icon>
          </el-menu-item>
          <el-divider></el-divider>
          <el-menu-item index="/home/personal">
            <el-icon><User /></el-icon>
            <span class="menu-item-text">个人信息</span>
          </el-menu-item>
          <el-menu-item index="/home/search">
            <el-icon><Search /></el-icon>
            <span class="menu-item-text">搜索召唤师</span>
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
.el-container {
  height: 100%;
}

.el-main {

}

.el-menu {
  height:100vh;
  transition: 0.15s;
  overflow: hidden;
}

.el-aside {
  width: auto;
  overflow: visible;
}

.el-divider--horizontal{
  margin: 0;
}

.el-menu-item {
  border-radius: 5px;
  margin: 0;

}

.menu-collapse-icon {
  width: 100%;
}

.menu-item-text {
  margin-left: 5px;
  margin-right: 20px;
}

</style>