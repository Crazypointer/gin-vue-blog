<template>
  <el-row>
    <el-col :span="8">
      <div class="grid-content ep-bg-purple">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <el-avatar :size="300" src="https://avatars.githubusercontent.com/u/30772089?v=4" />
            </div>
          </template>
          <div class="text item">注册IP: 127.0.0.1</div>
          <div class="text item">注册时间: 中国 四川</div>
        </el-card>
      </div>
    </el-col>
    <el-col :span="16">
      <div class="grid-content ep-bg-purple-light">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>博客信息</span>
            </div>
          </template>
          <el-row :gutter="12">
            <el-col :span="6">
              <el-card shadow="hover">
                <div class="info">
                  <div>
                    <el-icon size="60"><Document /></el-icon>
                  </div>
                  <div class="detail">
                    <div>文章</div>
                    <div class="number">{{ articleCount }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card shadow="hover">
                <div class="info">
                  <div>
                    <el-icon size="60"><PieChart /></el-icon>
                  </div>
                  <div class="detail">
                    <div>分类</div>
                    <div class="number">{{ 5 }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>

            <el-col :span="6">
              <el-card shadow="hover">
                <div class="info">
                  <div>
                    <el-icon size="60"><PriceTag /></el-icon>
                  </div>
                  <div class="detail">
                    <div>标签</div>
                    <div class="number">{{ tagCount }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>

            <el-col :span="6">
              <el-card shadow="hover">
                <div class="info">
                  <div>
                    <el-icon size="60"><User /></el-icon>
                  </div>
                  <div class="detail">
                    <div>用户</div>
                    <div class="number">{{ userCount }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </el-card>
      </div>
    </el-col>
  </el-row>
</template>
<script setup>
import axios from 'axios';
import { onMounted, computed, ref } from 'vue';
const userCount = ref(0);
const articleCount = ref(0);
const tagCount = ref(0);
//获取文章数
function getArticleCount() {
  axios.get('/api/articles').then(res => {
    articleCount.value = res.data.data.count;
  });
}
//获取标签数
function getTagCount() {
  axios.get('/api/tags').then(res => {
    tagCount.value = res.data.data.count;
  });
}

function getUserCount() {
  //获取当前用户的token
  const token = localStorage.getItem('token');
  axios
    .get('/api/users', {
      //请求头中设置token
      headers: {
        token: token,
      },
    })
    .then(res => {
      console.log(res);
      userCount.value = res.data.data.count;
    });
}
onMounted(() => {
  getUserCount();
  getArticleCount();
  getTagCount();
});
</script>

<style lang="less" scoped>
.info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  .detail {
    font-size: 20px;
    font-weight: bold;
    .number {
      margin-top: 4px;
      text-align: center;
    }
  }
}
</style>
