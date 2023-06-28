<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px"><SideBar /></el-aside>
      <el-container>
        <el-header><HeadBar /></el-header>
        <el-main><router-view></router-view></el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script setup>
import SideBar from './components/SideBar.vue';
import HeadBar from './components/HeadBar.vue';
</script>
<style lang="less" scoped>
.el-header {
  padding: 0;
}
</style>
