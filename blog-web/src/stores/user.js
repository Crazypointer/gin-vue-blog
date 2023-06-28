import { ref, computed, reactive } from 'vue';
import { defineStore } from 'pinia';
import axios from 'axios';

export const useUserStore = defineStore('user', () => {
  const userList = ref([]);
  const userNumber = computed(() => {
    return userList.value.length;
  });
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
      });
  }

  return { userList, userNumber, getUserList };
});
