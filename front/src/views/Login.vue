<template>
<div class="login">
 <div class="content">
   <a-form
       :model="formState"
       name="basic"
       :label-col="{ span: 4 }"
       :wrapper-col="{ span: 16 }"
       autocomplete="off"
       @finish="onFinish"
   >
     <a-form-item
         label="账号"
         name="username"
         :rules="[{ required: true, message: '请输入账号!' }]"
     >
       <a-input v-model:value="formState.username" />
     </a-form-item>

     <a-form-item
         label="密码"
         name="password"
         :rules="[{ required: true, message: '请输入密码!' }]"
     >
       <a-input-password v-model:value="formState.password" />
     </a-form-item>

     <a-form-item name="remember" :wrapper-col="{ offset: 4, span: 19 }">
       <a-checkbox v-model:checked="formState.remember">记住我</a-checkbox>
     </a-form-item>

     <a-form-item :wrapper-col="{ offset: 4, span: 16 }">
       <a-button type="primary" html-type="submit">登录</a-button>
     </a-form-item>
   </a-form>
 </div>
</div>
</template>

<script setup>
// 导入组合式api
import {reactive} from "vue";
// 导入请求api
import {$Login} from "@/api/login";
// 表单数据
const formState = reactive({
  username: '',
  password: '',
  remember: false,
});

// 登录方法
const onFinish = values => {
  let{username,password} = values
  $Login({username,password})
};

</script>

<style lang="scss" scoped>
.login {
  width: 100vw;
  height: 100vh;
  background: linear-gradient(to bottom, #8686e5, #afafde);
  display: flex;
  justify-content: center;
  align-items: center;
  .content {
    padding: 20px 20px 0;
    width: 450px;
    height: 250px;
    border: 1px solid peachpuff;
    border-radius: 6px;
  }
}

</style>