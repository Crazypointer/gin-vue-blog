<template>
  <div class="login">
    <h2>登录</h2>
    <p style="font-size: 14px; color: gray">在进入系统前请先输入密码进行登录</p>
    <div style="margin-top: 50px">
      <div style="margin: 0 20px">
        <el-input type="text" placeholder="用户名/邮箱登录" v-model="form.username">
          <template #prefix>
            <el-icon><User /></el-icon>
          </template>
        </el-input>
        <el-input type="password" style="margin-top: 10px" placeholder="请输入密码" v-model="form.password">
          <template #prefix>
            <el-icon><Lock /></el-icon>
          </template>
        </el-input>
        <div style="margin-top: 5px">
          <el-row style="color: gray">
            <el-col style="text-align: left" :span="12"> <el-checkbox label="记住我" size="large" v-model="form.remember" /></el-col>
            <el-col style="text-align: right" :span="12"><el-checkbox label="忘记密码" size="large" /></el-col>
          </el-row>
          <el-button type="success" @click="login()">立即登录</el-button>
          <el-divider>
            <span> 没有账号 </span>
          </el-divider>
          <router-link to="/register">
            <el-button type="warning" @click="register()">注册账号</el-button>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ElMessage } from 'element-plus';
import { reactive } from 'vue';
import router from '@/router/index.js';
import axios from 'axios';
const form = reactive({
  username: '',
  password: '',
});
const login = () => {
  if (!form.username || !form.password) {
    ElMessage.warning('请填写用户名和密码');
  } else {
    // 在组件中调用登录请求
    axios
      .post('/api/email_login', {
        user_name: form.username,
        password: form.password,
      })
      .then(res => {
        console.log(res);
        if (res.status === 200) {
          ElMessage.success('登录成功');
          //获取返回的token
          let token = res.data.data;
          console.log(token);
          // 设置 Token 值到 localStorage
          localStorage.setItem('token', token); // 替换为你的 Token 值
          router.push('/admin');
        } else {
          ElMessage.error('登录失败');
        }
      });
  }
};
const register = () => {
  router.push('/register');
};
</script>
<style scoped>
.login {
  font-size: 28px;
  margin-top: 150px;
}
</style>
