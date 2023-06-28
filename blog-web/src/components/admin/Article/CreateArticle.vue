<template>
  <div class="article">
    <p>文章标题</p>
    <el-input v-model="article.title" placeholder="请输入文章标题" />
    <p>文章关键词</p>
    <el-input v-model="article.keyword" placeholder="请输入文章关键词" />
    <p>文章摘要</p>
    <el-input v-model="article.abstract" placeholder="请输入文章摘要" />
    <div class="m-4">
      <p>请选择文章分类</p>
      <el-cascader v-model="article.category" :options="options" @change="handleChange" />
    </div>
    <div class="editor">
      <p>文章内容</p>
      <v-md-editor v-model="article.content" height="400px"></v-md-editor>
    </div>

    <p>请输入文章来源</p>
    <el-input v-model="article.source" placeholder="请输入文章来源" />

    <p>请输入原文章链接</p>
    <el-input v-model="article.link" placeholder="请输入原文章链接" />
    <div class="tags">
      <p>请选择文章标签</p>
      <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable :disable-transitions="false" @close="handleClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input
        v-if="inputVisible"
        ref="InputRef"
        v-model="inputValue"
        class="ml-1 w-20"
        size="small"
        @keyup.enter="handleInputConfirm"
        @blur="handleInputConfirm"
      />
      <el-button v-else class="button-new-tag ml-1" size="small" @click="showInput"> + New Tag </el-button>
    </div>

    <div class="banner">
      <p>请上传文章封面</p>
      <el-upload
        ref="uploadRef"
        :file-list="fileList"
        list-type="picture-card"
        action="http://127.0.0.1:8080/api/images"
        name="images"
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove"
        :on-success="handleSuccess"
        :auto-upload="false"
      >
        <el-icon><Plus /></el-icon>
      </el-upload>

      <el-dialog v-model="dialogVisible">
        <img w-full :src="dialogImageUrl" alt="Preview Image" />
      </el-dialog>
      <el-button class="ml-3" type="success" @click="submitUpload" style="margin-top: 10px"> 上传到服务器 </el-button>
    </div>

    <el-button type="primary" class="save" @click="save">保存文章</el-button>
  </div>
</template>
<script lang="ts" setup>
import { nextTick, reactive, ref, computed } from 'vue';
import { UploadProps, UploadUserFile, UploadInstance, ElInput, ElMessage } from 'element-plus';
import axios from 'axios';
const article = reactive({
  title: '',
  keyword: '',
  content: '',
  category: '',
  source: '',
  link: '',
  banner_id: 1,
  banner_url: '',
  abstract: '',
  tags: [''],
});
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
//文章分类
const handleChange = value => {
  article.category = value[0];
};

//上传图片
const fileList = ref<UploadUserFile[]>([]);
const dialogImageUrl = ref('');
const dialogVisible = ref(false);

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
  // console.log(uploadFile, uploadFiles);
};

const handlePictureCardPreview: UploadProps['onPreview'] = uploadFile => {
  dialogImageUrl.value = uploadFile.url!;
  dialogVisible.value = true;
};

const handleSuccess = (response, file) => {
  article.banner_url = response.data[0].fileName;
  article.banner_id = response.data[0].id;
};
//手动上传图片
const uploadRef = ref<UploadInstance>();
const submitUpload = () => {
  uploadRef.value!.submit();
};

//上传tag
const inputValue = ref('');
const dynamicTags = ref<string[]>([]);
const inputVisible = ref(false);
const InputRef = ref<InstanceType<typeof ElInput>>();

// 删除标签
const handleClose = (tag: string) => {
  dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1);
};

// 展示输入框
const showInput = () => {
  inputVisible.value = true;
  nextTick(() => {
    InputRef.value!.input!.focus();
  });
};

// 输入新的标签
const handleInputConfirm = () => {
  if (inputValue.value) {
    dynamicTags.value.push(inputValue.value);
  }
  inputVisible.value = false;
  inputValue.value = '';
};
// 将dynamicTags数组绑定到article.tags
const tags = computed(() => {
  return dynamicTags.value;
});

// 保存文章
function save() {
  article.tags = tags.value;
  console.log(article);
  //请求接口http://127.0.0.1:8080/api/articles
  axios.post('/api/articles', article).then(res => {
    console.log(res);
    if (res.data.code === 0) {
      ElMessage.success(res.data.msg);
    } else {
      ElMessage.warning(res.data.msg);
    }
  });
}
</script>
<style lang="less" scoped>
p {
  margin: 10px 0;
}
.banner {
  margin-top: 20px;
}
.save {
  margin-top: 20px;
}
</style>
