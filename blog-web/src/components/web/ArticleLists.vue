<template>
  <div class="title">博客文章</div>
  <div class="card" v-for="article in articleList" :key="article.id">
    <div class="left">
      <el-image style="width: 300px; height: 200px; margin-left: 5px; border-radius: 5px" :src="getImageUrl(article.banner_url)" fit="cover" />
    </div>
    <div class="right">
      <router-link :to="{ name: 'ArticleContent', params: { id: article.id } }">
        <h2>{{ article.title }}</h2>
      </router-link>
      <div class="content">
        {{ article.content }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
const articleList = ref([]);
function getArticleList() {
  axios.get('/api/articles').then(res => {
    articleList.value = res.data.data.list;
    for (let i = 0; i < articleList.value.length; i++) {
      //判断content字段的长度是否大于200
      if (articleList.value[i].content.length > 400) {
        //如果大于100，截取前200个字符
        articleList.value[i].content = articleList.value[i].content.substring(0, 400) + '...';
      }
    }
  });
}
function getImageUrl(path) {
  // 替换成你的服务器基本 URL 地址
  const baseURL = 'http://127.0.0.1:8080';
  return `${baseURL}/${path}`;
}
onMounted(() => {
  getArticleList();
});
</script>
<style lang="less" scoped>
.title {
  background-color: #fff;
  font-weight: bold;
  font-size: 20px;
  margin-top: 20px;
  margin-right: 20px;
  padding: 20px 20px 10px 20px;
  border-radius: 5px 5px 0 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.card {
  height: 210px;
  margin: 10px 0;
  background-color: #fff;
  box-shadow: 0 0 10px #ccc;
  margin-right: 20px;
  border-radius: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .left {
    width: 30%;
  }
  .right {
    width: 70%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 5px;
    h2 {
      font-size: 28px;
      padding: 10px 0;
    }
    .content {
      font-size: 16px;
      height: 150px;
      padding: 5px;
    }
  }
}

a {
  text-decoration: none;
  color: black;
}
</style>
