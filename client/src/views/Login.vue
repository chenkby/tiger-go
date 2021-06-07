<template>
  <div class="page-login">
    <div class="login-form">
      <div class="brand-container">
        <brand></brand>
      </div>
      <el-form :model="loginForm" :rules="rules" ref="loginForm" label-width="0" @submit.native.prevent>
        <el-form-item prop="username" class="username-item">
          <el-input type="text" prefix-icon="el-icon-user" v-model="loginForm.username" placeholder="输入登录帐号">
          </el-input>
        </el-form-item>
        <el-form-item prop="password" class="password-item">
          <el-input type="password" prefix-icon="el-icon-lock" v-model="loginForm.password" autocomplete="off" placeholder="输入登录密码">

          </el-input>
        </el-form-item>
        <div class="other-items">
          <div class="other-items__left">
            <el-checkbox v-model="loginForm.remember" :true-label="1" :false-label="0">自动登录</el-checkbox>
          </div>
          <!-- <div class="other-items__right">
            <el-link type="primary">忘记密码</el-link>
          </div> -->
        </div>
        <el-form-item>
          <el-button type="primary" native-type="submit" class="submit" @click="submitForm('loginForm')">登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>

  </div>
  <div class="page-footer">copyright</div>
</template>

<script>
import { defineComponent, ref } from 'vue'
import Brand from '@/components/Brand.vue'

export default defineComponent({
  components: {
    Brand
  },
  data() {
    return {
      loginForm: {
        username: '',
        password: '',
        remember: 0
      },
      rules: {
        username: [{ required: true, message: '请输入登录帐号' }],
        password: [{ required: true, message: '请输入登录密码' }]
      }
    }
  },
  setup() {
    const name = ref('aaa')
    function test() {
      console.log('这是测试方法')
    }
    return {
      test
    }
  }
})
</script>

<style lang="scss">
body {
  padding: 0;
  height: 100%;
}
.page-login {
  position: relative;
  height: 100%;
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;

  &:after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url(../assets/images/login-bg.jpg) no-repeat #f5f7fa;
    background-size: cover;

    filter: blur(2px);
  }
  .brand-container {
    margin-bottom: 40px;
  }
  .login-form {
    width: 420px;
    margin: 0 auto;
    background-color: #fff;
    padding: 30px;
    box-sizing: border-box;
    z-index: 1;
    border-radius: 2px;
    box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.06),
      0 6px 12px 0 rgba(0, 0, 0, 0.06);

    .username-item,
    .password-item {
      margin-bottom: 30px;
    }

    .other-items {
      text-align: left;
      margin-bottom: 30px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      line-height: 1;
      .el-checkbox__label {
        font-weight: normal;
        color: #666;
      }
      .el-link {
        font-weight: normal;
      }
    }
    .submit {
      width: 100%;
      font-size: 16px;
      padding: 12px 20px;
    }
  }
}
.page-footer {
  color: #f5f5f5;
  z-index: 1;
  position: absolute;
  bottom: 30px;
  text-align: center;
  width: 100%;
  font-size: 14px;
}
// 手机端
body.device-mobile {
  .page-login {
    display: block;
    padding-top: 20%;
    &:after {
      display: none;
    }
    .login-form {
      box-shadow: none;
      width: 100%;
    }
  }
  .page-footer {
    color: #999;
  }
}
</style>
