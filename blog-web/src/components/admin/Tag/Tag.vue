<template>
  <div class="tag">
    <el-row class="mb-4">
      <el-button type="success" @click="dialogFormVisible = true">添加标签</el-button>
      <el-button type="danger" @click="deleteTag">删除</el-button>
    </el-row>
    <div class="tag-list">
      <el-table :border="true" :data="tagList" style="width: 100%" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="40" />
        <el-table-column prop="id" label="ID" width="80"> </el-table-column>
        <el-table-column prop="created_at" label="创建日期" width="200"> </el-table-column>
        <el-table-column prop="title" label="名称" width="120"> </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" title="添加新标签">
      <el-form :model="form">
        <el-form-item label="标签名">
          <el-input v-model="form.title" autocomplete="off" />
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
import { ref, reactive, onMounted } from 'vue';
import axios from 'axios';
const dialogFormVisible = ref(false);
// 获取标签列表
const tagList = ref([]);
function getTagList() {
  //获取当前用户的token
  const token = localStorage.getItem('token');
  axios
    .get('/api/tags', {
      //请求头中设置token
      headers: {
        token: token,
      },
    })
    .then(res => {
      tagList.value = res.data.data.list;
      //缩短created_at字段
      tagList.value.forEach(item => {
        item.created_at = item.created_at.substr(0, 10);
      });
    });
}

const selectedRows = ref([]); // 保存被选中的列数据
const id_list = ref([]); // 保存被选中的列数据的id

// 表单数据
const form = reactive({
  title: '',
});
// 获取被选中的列数据
function handleSelectionChange(val) {
  selectedRows.value = val;
  getIdList();
}
// 处理选中的数据
function getIdList() {
  id_list.value = selectedRows.value.map(item => item.id);
  console.log(id_list.value);
}

// 删除标签
function deleteTag() {
  const token = localStorage.getItem('token');
  axios
    .delete('/api/tags', {
      data: {
        id_list: id_list.value,
      },
      headers: {
        token: token,
      },
    })
    .then(res => {
      console.log(res.data);
    });
  tagList.value = tagList.value.filter(item => !id_list.value.includes(item.id));
  //缩短created_at字段
  tagList.value.forEach(item => {
    item.created_at = item.created_at.substr(0, 10);
  });
}

// 添加标签
function submit() {
  const token = localStorage.getItem('token');
  axios
    .post(
      '/api/tags',
      {
        title: form.title,
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
  //添加成功后重新获取标签列表
  getTagList();
}

onMounted(() => {
  getTagList();
});
</script>

<style lang="less" scoped>
.tag {
  margin-top: 5px;
  .tag-list {
    margin-top: 16px;
  }
}
</style>
