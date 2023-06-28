<template>
  <div class="advert">
    <el-row class="mb-4">
      <el-button type="success">添加广告</el-button>
      <el-button type="danger">删除</el-button>
    </el-row>
    <div class="advert-list">
      <el-table :border="true" :data="tableData" style="width: 100%">
        <el-table-column type="selection" width="40" />
        <el-table-column label="ID" width="80">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <span style="margin-left: 5px">{{ scope.row.id }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Title" width="120">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <span style="margin-left: 5px">{{ scope.row.title }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Href" width="200">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <span style="margin-left: 5px">{{ scope.row.href }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Images" width="200">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <span style="margin-left: 5px">{{ scope.row.images }}</span>
            </div>
          </template>
        </el-table-column>

        <el-table-column label="是否展示" width="120">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <span style="margin-left: 5px">{{ scope.row.is_show }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Operations">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination style="margin-top: 20px" background layout="prev, pager, next" :total="50" />
    </div>
  </div>
</template>
<style lang="less" scoped>
.advert {
  margin-top: 5px;
  .advert-list {
    margin-top: 16px;
  }
}
</style>
<script lang="ts" setup>
interface Adverts {
  id: number;
  created_at: string;
  title: string;
  href: string;
  images: string;
  is_show: boolean;
}

const handleEdit = (index: number, row: Adverts) => {
  console.log(index, row);
};
const handleDelete = (index: number, row: Adverts) => {
  console.log(index, row);
};

const tableData: Adverts[] = [
  {
    id: 10,
    created_at: '2023-05-28T18:10:33.638+08:00',
    title: 'abc',
    href: 'http://www.baidu.com',
    images: 'http://www.baidu.com',
    is_show: true,
  },
  {
    id: 9,
    created_at: '2023-05-27T20:32:59.962+08:00',
    title: 'abc',
    href: 'http://www.baidu.com',
    images: 'http://www.baidu.com',
    is_show: true,
  },
];
</script>
