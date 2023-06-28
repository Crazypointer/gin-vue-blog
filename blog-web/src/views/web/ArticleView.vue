<template>
  <div class="main" v-if="article">
    <div class="content">
      <div class="article">
        <div class="article-title">
          <h2 class="name">{{ article.title }}</h2>
          <p class="info">
            <span>
              发布时间：<i>{{ article.created_at }}</i>
            </span>
            <span>
              作者：<i>{{ article.user_nick_name }}</i>
            </span>
            <span>
              来源：<i>{{ article.source }}</i>
            </span>
          </p>
          <div class="tag">
            <el-icon><CollectionTag /></el-icon>
            <i v-for="tag in article.tags">{{ tag }}</i>
          </div>
        </div>
        <div class="article-content">
          <v-md-preview :text="article.content"></v-md-preview>
        </div>
        <div class="comment">评论区</div>
      </div>
    </div>
    <div class="good">点赞按钮区</div>
  </div>
</template>
<script setup>
import { ref, onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios';

//获取当前路由中的参数
const route = useRoute();
//获取文章
const article = ref();
function getArticleById() {
  let { id } = route.params;
  //请求后端接口
  axios.get(`/api/articles/${id}`).then(res => {
    console.log(res.data.data);
    article.value = res.data.data;
  });
}
onBeforeMount(() => {
  getArticleById();
});
</script>
<style lang="less" scoped>
.main {
  width: 1550px;
  margin: 0 160px;
  margin-top: 20px;
  display: flex;
  .article {
    background-color: white;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .article-title {
      height: 120px;
      padding: 10px;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      .info {
        padding: 4px 0;
      }
      .tag {
        display: flex;
        text-align: center;
        align-items: center;
        justify-content: center;
        i {
          margin: 0 5px;
          color: #409eff;
        }
        i:last-child {
          margin-right: 0;
        }
      }
    }
    .article-content {
      margin-top: 1px;
      width: 100%;
      min-height: 200px;
      padding: 20px;
    }
    .comment {
      height: 150px;
      margin-top: 1px;
      width: 100%;
      padding: 20px;
    }
  }
  .content {
    width: 90%;
    margin-right: 20px;
  }
  .good {
    width: 10%;
    height: 100px;
    margin-left: 20px;
    background-color: white;
  }
}
</style>
