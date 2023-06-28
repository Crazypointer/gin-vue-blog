<template>
  <div class="article-list" v-show="dialogVisible == false">
    <el-table :border="true" :data="articleList" style="width: 100%">
      <el-table-column prop="id" label="ID" width="39"> </el-table-column>
      <el-table-column prop="title" label="标题" width="130"> </el-table-column>
      <el-table-column prop="created_at" label="创建日期" width="130"> </el-table-column>
      <el-table-column prop="keyword" label="关键词" width="120"> </el-table-column>
      <el-table-column prop="user_nick_name" label="作者昵称" width="100"> </el-table-column>
      <el-table-column prop="category" label="分类" width="100"></el-table-column>
      <el-table-column prop="source" label="来源" width="120"> </el-table-column>
      <el-table-column prop="link" label="来源链接" width="120"> </el-table-column>
      <el-table-column prop="banner_id" label="封面id" width="80"> </el-table-column>
      <el-table-column prop="banner_url" label="封面图链接" width="150"> </el-table-column>
      <el-table-column prop="tags" label="标签" width="120"> </el-table-column>
      <el-table-column label="Operations">
        <template #default="scope">
          <el-button type="primary" @click="editArticle(scope.row)">编辑</el-button>
          <el-popconfirm title="你确定删除这篇文章吗？" confirm-button-text="确认" cancel-button-text="取消" @confirm="handleDelete(scope.row)">
            <template #reference>
              <el-button type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :total="50" />
  </div>
  <!-- <EditArticle v-if="dialogVisible" :article="article"></EditArticle> -->
</template>
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import router from '../../../router';
const articleList = ref([]);
function getArticleList() {
  //获取当前用户的token
  const token = localStorage.getItem('token');
  axios
    .get('/api/articles', {
      //请求头中设置token
      headers: {
        token: token,
      },
    })
    .then(res => {
      articleList.value = res.data.data.list;
      //缩短created_at字段
      articleList.value.forEach(item => {
        item.created_at = item.created_at.substr(0, 10);
      });
    });
}
//编辑文章
const dialogVisible = ref(false);
//文章信息
const article = ref({});
function editArticle(row) {
  // console.log(row);
  // dialogVisible.value = true;
  // article.value = row;
  router.push('/admin/article/edit/' + row.id);
}
const handleDelete = row => {
  console.log(row);
};

const options = ref([
  {
    value: 'story',
    label: '故事',
  },
  {
    value: 'note',
    label: '笔记',
  },
]);
onMounted(() => {
  getArticleList();
});
</script>
<style lang="less">
.article-list {
  margin-top: 16px;
}
</style>
