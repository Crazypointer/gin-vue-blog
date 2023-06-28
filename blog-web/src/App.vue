<script setup></script>
<template>
  <router-view :key="key"></router-view>
</template>

<script setup>
import { computed } from 'vue';
const key = computed(() => {
  //获取当前时间戳
  return Date.now();
});
</script>

<style lang="less">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
h1,
h2,
h3,
h4,
h5,
h6 {
  margin: 0;
}
a {
  text-decoration: none;
}
</style>
