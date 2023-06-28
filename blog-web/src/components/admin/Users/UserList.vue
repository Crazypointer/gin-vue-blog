<template>
  <div class="user-list">
    <el-table :border="true" :data="userList" style="width: 100%">
      <el-table-column type="selection" width="40" />
      <el-table-column prop="id" label="ID" width="40"> </el-table-column>
      <el-table-column prop="user_name" label="UserName" width="120"> </el-table-column>
      <el-table-column prop="nick_name" label="NickName" width="120"> </el-table-column>
      <el-table-column prop="email" label="Email" width="120"> </el-table-column>
      <el-table-column prop="created_at" label="Date" width="150"> </el-table-column>
      <el-table-column prop="avatar" label="Avatar" width="120"> </el-table-column>
      <el-table-column prop="role" label="Role" width="120"> </el-table-column>
      <el-table-column prop="ip" label="IP" width="120"> </el-table-column>
      <el-table-column prop="sign_status" label="Status" width="120"> </el-table-column>
    </el-table>
    <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :total="50" />
  </div>
</template>
<style lang="less" scoped>
.user-list {
  margin-top: 16px;
}
</style>
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
const userList = ref([]);
function getUserList() {
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
      userList.value = res.data.data.list;
      //缩短created_at字段
      userList.value.forEach(item => {
        item.created_at = item.created_at.substr(0, 10);
      });
    });
}
onMounted(() => {
  getUserList();
});
</script>
