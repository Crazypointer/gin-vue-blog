<template>
  <div class="box">
    <div class="left">
      <el-image style="width: 100%; height: 100%" src="./src/assets/photo.jpg" fit="cover" />
    </div>
    <div class="right">
      <router-view></router-view>
    </div>
  </div>
</template>

<style lang="less" scoped>
.box {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  display: flex;
  .left {
    flex: 1;
  }
  .right {
    width: 400px;
    background-color: #fff;
    text-align: center;
  }
}
</style>
