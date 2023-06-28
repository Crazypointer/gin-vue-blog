<template>
  <div class="index">
    <HeaderBar />
    <router-view :key="key"></router-view>
  </div>
</template>
<script setup>
import { computed } from 'vue';
const key = computed(() => {
  //获取当前时间戳
  return Date.now();
});
</script>
<style lang="less">
.index {
  background-color: #e2e2e2;
}
</style>
