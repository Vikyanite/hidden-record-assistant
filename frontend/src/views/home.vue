<script setup lang="ts">

import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {Close} from "@element-plus/icons-vue";
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
      <el-aside>
        <el-menu
            default-active="/home/personal"
            mode="vertical"
            router
            :collapse="isCollapse"
        >
          <el-menu-item @click="isCollapse=!isCollapse" class="menu-collapse-icon">
            <el-icon v-show="isCollapse"><Expand /></el-icon>
            <el-icon v-show="!isCollapse"><Fold /></el-icon>
          </el-menu-item>
          <el-divider></el-divider>
          <el-menu-item index="/home/personal">
            <el-icon><User /></el-icon>
            <span class="menu-item-text">个人信息</span>
          </el-menu-item>
          <el-menu-item index="/home/test1">
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

.el-menu {
  height:100vh;
  overflow: hidden;
  transition: 0.2s;
}

.el-aside {
  width: auto;
}

.el-menu--collapse {
  width: calc(var(--el-menu-icon-width) + var(--el-menu-base-level-padding) * 2 + var(--el-menu-item-margin) * 2);
}

.el-divider--horizontal{
  margin: 0;
}

.el-menu-item {
  border-radius: 5px;
  margin: 5px;
}

.menu-collapse-icon {
  width: 35px;
}

.menu-item-text {
  margin-left: 5px;
  margin-right: 20px;
}

</style>