<template>
  <div class="index-main">
    <div class="left">
      <ArticleLists />
    </div>
    <div class="right">
      <WebInfo />
    </div>
  </div>
</template>
<script setup></script>
<style lang="less" scoped>
.index-main {
  margin: 0 160px;
  display: flex;
  .left {
    width: 70%;
  }
  .right {
    width: 30%;
  }
}
</style>
