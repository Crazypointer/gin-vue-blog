<template>
  <div class="header-container">
    <div class="l-content">
      <el-button :icon="Menu" size="small" />
      <span>首页</span>
    </div>
    <div class="r-content">
      <el-dropdown>
        <span class="el-dropdown-link">
          <el-avatar size="default" src="https://avatars.githubusercontent.com/u/30772089?v=4" />
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>个人中心</el-dropdown-item>
            <el-dropdown-item @click="logout()">退出系统</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<style lang="less" scoped>
.header-container {
  background-color: #333;
  height: 60px;
  display: flex;
  padding: 0 20px;
  justify-content: space-between;
  align-items: center;
  .l-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    span {
      margin: 0 20px;
      color: #fff;
      font-size: 16px;
    }
  }
}
</style>
<script lang="ts" setup>
import { Menu } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import router from '../../../router/index';
import axios from 'axios';
const logout = () => {
  let token = localStorage.getItem('token');
  axios
    .post('/api/logout', {
      headers: {
        token: token,
      },
    })
    .then(res => {
      if (res.status === 200) {
        ElMessage.success('退出成功');
        // 清除 localStorage 中的 Token
        localStorage.removeItem('token');
        router.push('/');
      } else {
        ElMessage.error('退出失败');
      }
    });
};
</script>
