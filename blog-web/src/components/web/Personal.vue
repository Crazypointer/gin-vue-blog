<template>
  <el-card class="person">
    <div class="card">
      <div class="head">
        <el-image style="width: 200px; height: 200px; border-radius: 50%" src="http://www.fengfengzhidao.com/media/site_bg/wallhaven-l3geq2.jpg" />
      </div>
      <div class="info">
        <div>Name: Cyrder</div>
        <div>Job: 全栈开发</div>
        <div>Addr: 四川成都</div>
        <div>Web: www.baidu.com</div>
      </div>
    </div>
  </el-card>
</template>

<style lang="less" scoped>
.person {
  margin-top: 20px;
  .card {
    width: 100%;
    display: flex;
    flex-direction: column;
    .head {
      width: 100%;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }
    .info {
      margin-top: 20px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      font-size: 18px;
      font-weight: 450;
    }
  }
}
</style>
