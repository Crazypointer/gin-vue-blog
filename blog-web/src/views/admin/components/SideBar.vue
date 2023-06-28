<template>
  <el-menu
    default-active="1"
    class="el-menu-vertical-demo aside-menu"
    active-text-color="#ffd04b"
    background-color="#545c64"
    text-color="#fff"
    style="border: 0"
  >
    <el-menu-item index="1">
      <h3>Cyrder Blog</h3>
    </el-menu-item>
    <router-link to="/admin">
      <el-menu-item index="2">
        <el-icon><HomeFilled /></el-icon>
        <span>首页</span>
      </el-menu-item>
    </router-link>
    <router-link to="/admin/user">
      <el-menu-item index="3">
        <template #title>
          <el-icon><UserFilled /></el-icon>
          <span>用户管理</span>
        </template>
      </el-menu-item>
    </router-link>
    <el-sub-menu index="4">
      <template #title>
        <el-icon><UserFilled /></el-icon>
        <span>文章管理</span>
      </template>
      <router-link to="/admin/article/list"><el-menu-item index="4-1">文章列表</el-menu-item></router-link>
      <router-link to="/admin/article/add"><el-menu-item index="4-2">添加文章</el-menu-item></router-link>
    </el-sub-menu>
    <router-link to="/admin/advert">
      <el-menu-item index="5">
        <template #title>
          <el-icon><Promotion /></el-icon>
          <span>广告管理</span>
        </template>
      </el-menu-item>
    </router-link>
    <router-link to="/admin/tag">
      <el-menu-item index="6">
        <template #title>
          <el-icon><CollectionTag /></el-icon>
          <span>标签管理</span>
        </template>
      </el-menu-item>
    </router-link>
    <router-link to="/admin/menu">
      <el-menu-item index="7">
        <template #title>
          <el-icon><Menu /></el-icon>
          <span>菜单</span>
        </template>
      </el-menu-item>
    </router-link>
    <router-link to="/admin/info">
      <el-menu-item index="8">
        <template #title>
          <el-icon><Tools /></el-icon>
          <span>网站管理</span>
        </template>
      </el-menu-item>
    </router-link>
  </el-menu>
</template>
<style lang="less" scoped>
.aside-menu {
  height: 100vh;
  h3 {
    color: #fff;
    text-align: center;
    line-height: 48px;
    font-size: 16px;
    font-weight: 400;
  }
}
</style>
