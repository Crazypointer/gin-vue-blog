<template>
  <div class="user">
    <el-row class="mb-4">
      <el-button type="success" @click="dialogFormVisible = true">添加用户</el-button>
      <el-button type="danger">删除</el-button>
    </el-row>
    <UserList />
    <el-dialog v-model="dialogFormVisible" title="添加用户">
      <el-form :model="form">
        <el-form-item label="用户名" :label-width="formLabelWidth">
          <el-input v-model="form.user_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="昵称" :label-width="formLabelWidth">
          <el-input v-model="form.nick_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="用户权限" :label-width="formLabelWidth">
          <el-select v-model="form.region" placeholder="请选择用户权限">
            <el-option label="admin" value="1" />
            <el-option label="user" value="2" />
            <el-option label="customer" value="3" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="submit"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup>
import axios from 'axios';
import { reactive, ref } from 'vue';

const dialogFormVisible = ref(false);
const formLabelWidth = '130px';
const form = reactive({
  user_name: '',
  nick_name: '',
  password: '',
  role: 2,
});
function submit() {
  const token = localStorage.getItem('token');
  axios
    .post(
      '/api/users',
      {
        user_name: form.user_name,
        nick_name: form.nick_name,
        password: form.password,
        role: form.role,
      },
      {
        //请求头中设置token
        headers: {
          token: token,
        },
      }
    )
    .then(res => {
      console.log(res.data);
    });
  dialogFormVisible.value = false;
}
</script>
<style lang="less" scoped>
.user {
  margin-top: 5px;
}
</style>
