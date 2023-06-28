<template>
  <el-row>
    <el-col>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>网站信息</span>
          </div>
        </template>
        <div class="text item">网站名：{{ siteInfo.title }}</div>
        <div class="text item">建站时间：{{ siteInfo.created_at }}</div>
        <div class="text item">网站备案号：{{ siteInfo.beian }}</div>
        <div class="text item">当前版本：{{ siteInfo.version }}</div>
      </el-card>
    </el-col>
    <el-col>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Email</span>
          </div>
        </template>
        <div class="text item">邮箱：{{ emailInfo.user }}</div>
        <div class="text item">Host：{{ emailInfo.host }}</div>
        <div class="text item">Port：{{ emailInfo.port }}</div>
        <div class="text item">默认发件人姓名：{{ emailInfo.default_from_email }}</div>
      </el-card>
    </el-col>
    <el-col>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>七牛</span>
          </div>
        </template>
        <div class="text item">是否开启：{{ qiniuInfo.enable }}</div>
        <div class="text item">前缀：{{ qiniuInfo.prefix }}</div>
        <div class="text item">单个图片大小：{{ qiniuInfo.size }}MB</div>
        <div class="text item">地区：{{ qiniuInfo.zone }}</div>
      </el-card>
    </el-col>
    <el-col>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>JWT</span>
          </div>
        </template>
        <div class="text item">密钥：{{ jwtInfo.secret }}</div>
        <div class="text item">token过期时间：{{ jwtInfo.expires }}</div>
        <div class="text item">签发人：{{ jwtInfo.issuer }}</div>
      </el-card>
    </el-col>
  </el-row>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue';
import axios from 'axios';
const siteInfo = ref([]);
const qiniuInfo = ref([]);
const emailInfo = ref([]);
const jwtInfo = ref([]);
// 请求
function getSiteInfo(str, v) {
  const token = localStorage.getItem('token');
  let url = '/api/settings/' + str;
  axios
    .get(url, {
      //请求头中设置token
      headers: {
        token: token,
      },
    })
    .then(res => {
      v.value = res.data.data;
      console.log(v.value);
    });
}
onMounted(() => {
  getSiteInfo('site', siteInfo);
  getSiteInfo('qiniu', qiniuInfo);
  getSiteInfo('email', emailInfo);
  getSiteInfo('jwt', jwtInfo);
});
</script>
