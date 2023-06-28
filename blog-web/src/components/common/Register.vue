<template>
  <div class="register">
    <h2>注册新用户</h2>
    <p style="font-size: 14px; color: gray">欢迎注册本系统，请在下方填写相关信息</p>
    <el-form :model="form" :rules="rules" @validate="onValidate" ref="formRef">
      <el-form-item prop="username">
        <el-input v-model="form.username" placeholder="用户名">
          <template #prefix>
            <el-icon><User /></el-icon>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" placeholder="密码">
          <template #prefix>
            <el-icon><Lock /></el-icon>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="password_rp">
        <el-input type="password" v-model="form.password_rp" placeholder="重复密码">
          <template #prefix>
            <el-icon><Lock /></el-icon>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="email">
        <el-input type="email" v-model="form.email" placeholder="邮箱">
          <template #prefix>
            <el-icon><Message /></el-icon>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-row>
          <el-col :span="16">
            <el-input type="text" :maxlength="6" v-model="form.emailCode" placeholder="请输入邮箱验证码">
              <template #prefix>
                <el-icon><EditPen /></el-icon>
              </template>
            </el-input>
          </el-col>
          <el-col :span="4" style="margin-left: 9px">
            <el-button type="success" @click="validateEmail">获取邮箱验证码</el-button>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>

    <el-button type="primary" style="margin-top: 10px" @click="register">立即注册</el-button>
    <div style="margin-top: 20px">
      <span style="height: 14px; font-size: 14px; line-height: 14px; color: gray"> 已有账号？</span>
      <el-link type="primary" style="translate: 0 -2px" @click="router.push('/login')">返回登录界面</el-link>
    </div>
  </div>
</template>
<script setup>
import { reactive, ref } from 'vue';
import router from '@/router/index';
import { ElMessage } from 'element-plus';
const form = reactive({
  username: '',
  password: '',
  password_rp: '',
  email: '',
  emailCode: '',
});
const validateUsername = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请输入用户名'));
  } else if (!/^[a-zA-Z0-9\u4e00-\u9fa5]+$/.test(value)) {
    callback(new Error('用户名不能包含特殊字符，只能是中文/英文'));
  } else {
    callback();
  }
};

const validatePassword = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('请再次输入密码'));
  } else if (value !== form.password) {
    callback(new Error('两次输入的密码不一致'));
  } else {
    callback();
  }
};
const rules = {
  username: [
    { validator: validateUsername, trigger: ['blur', 'change'] },
    { min: 2, max: 8, message: '用户名的长度必须在2-8个字符之间', trigger: ['blur', 'change'] },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 16, message: '密码的长度必须在6-16个字符之间', trigger: ['blur', 'change'] },
  ],
  password_rp: [{ validator: validatePassword, trigger: ['blur', 'change'] }],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入合法的电子邮件地址', trigger: ['blur', 'change'] },
  ],
  // code: [{ required: true, message: '请输入获取的验证码', trigger: 'blur' }],
};

const formRef = ref();
const isEmailValid = ref(false);
const onValidate = (prop, isValid) => {
  if (prop === 'email') isEmailValid.value = isValid;
};
const register = () => {
  formRef.value.validate(isValid => {
    if (isValid) {
      post(
        '/api/auth/register',
        {
          username: form.username,
          password: form.password,
          email: form.email,
          code: form.code,
        },
        message => {
          ElMessage.success(message);
          router.push('/');
        }
      );
    } else {
      ElMessage.warning('请完整填写注册表单内容！');
    }
  });
};
const validateEmail = () => {
  post(
    '/api/auth/valid-register-email',
    {
      email: form.email,
    },
    message => {
      ElMessage.success(message);
    },
    message => {
      ElMessage.warning(message);
    }
  );
};
</script>
<style scoped>
.register {
  margin: 0 20px;
  margin-top: 150px;
}
</style>
