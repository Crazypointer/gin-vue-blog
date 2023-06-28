<template>
  <div class="header">
    <div class="nav">
      <div class="left">
        <router-link to="/">首页</router-link>
        <router-link to="/article">文章</router-link>
        <router-link to="/about">关于</router-link>
      </div>
      <div class="right">
        <router-link to="/login">登陆</router-link>
      </div>
    </div>
  </div>
  <Pictures />
</template>

<script setup>
import Pictures from './Pictures.vue';
const url = 'http://www.fengfengzhidao.com/media/site_bg/wallhaven-y8lqo7.jpg';
</script>
<style lang="less" scoped>
.header {
  background-color: black;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .nav {
    width: 100%;
    height: 60px;
    margin: 0 160px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    a {
      font-size: 20px;
      margin-right: 20px;
      color: white;
      cursor: pointer;
    }
    //最后一个a标签取消右边距
    a:last-child {
      margin-right: 0;
    }
  }
}
</style>
