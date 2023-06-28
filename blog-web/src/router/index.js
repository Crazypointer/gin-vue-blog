// 导入路由组件
import { createRouter, createWebHistory } from 'vue-router';

// 导入页面组件
import Home from '@/views/web/Index.vue';
import IndexView from '@/views/web/IndexView.vue';
import Article from '@/components/web/Article.vue';
import ArticleView from '@/views/web/ArticleView.vue';
import About from '@/components/web/About.vue';
import LoginView from '@/views/LoginView.vue';
import Login from '@/components/common/Login.vue';
import Register from '@/components/common/Register.vue';
import AdminIndexView from '@/views/admin/AdminIndexView.vue';
import AdminHome from '@/components/admin/Index.vue';
import UserManage from '@/components/admin/Users/UserManage.vue';
import Advert from '@/components/admin/Advert/Advert.vue';
import ArticleList from '@/components/admin/Article/ArticleList.vue';
import EditArticle from '@/components/admin/Article/EditArticle.vue';
import CreateArticle from '@/components/admin/Article/CreateArticle.vue';
import Tag from '@/components/admin/Tag/Tag.vue';
import MenuManage from '@/components/admin/Menu/MenuManage.vue';
import SystemInfo from '@/components/admin/Info/SystemInfo.vue';
import NotFound from '@/views/404.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    children: [
      {
        path: '/',
        name: 'home',
        component: IndexView,
      },
      {
        path: '/article',
        name: 'Article',
        component: Article,
      },
      {
        path: '/article/:id',
        name: 'ArticleContent',
        component: ArticleView,
      },
      {
        path: '/about',
        name: 'about',
        component: About,
      },
    ],
  },
  {
    path: '/login',
    name: 'loginV',
    component: LoginView,
    children: [
      {
        path: '',
        name: 'login',
        component: Login,
      },
      {
        path: '/register',
        name: 'register',
        component: Register,
      },
    ],
  },
  {
    path: '/admin',
    name: 'admin',
    component: AdminIndexView,
    children: [
      {
        path: '',
        name: 'admin-home',
        component: AdminHome,
      },
      {
        path: '/admin/user',
        name: 'admin-user',
        component: UserManage,
      },
      {
        path: '/admin/advert',
        name: 'admin-advert',
        component: Advert,
      },
      // 文章管理
      {
        path: '/admin/article',
        name: 'admin-article',
        children: [
          {
            path: '/admin/article/list',
            name: 'admin-article-list',
            component: ArticleList,
          },
          {
            path: '/admin/article/edit/:id',
            name: 'admin-article-edit',
            component: EditArticle,
          },
          {
            path: '/admin/article/add',
            name: 'admin-article-add',
            component: CreateArticle,
          },
        ],
      },

      {
        path: '/admin/tag',
        name: 'admin-tag',
        component: Tag,
      },
      {
        path: '/admin/menu',
        name: 'admin-menu',
        component: MenuManage,
      },
      {
        path: '/admin/info',
        name: 'admin-info',
        component: SystemInfo,
      },
    ],
  },

  {
    // 404
    path: '/:catchAll(.*)',
    name: 'not-found',
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
