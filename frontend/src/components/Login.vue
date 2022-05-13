<template>
  <div id="oj-login">
    <el-form class="login-form" label-width="80px">
      <div style="margin: auto; text-align: center; font-size: 30px;">登录</div>
      <el-form-item label="用户名" class="login-form-item">
        <el-input
            v-model="login.username"
            placeholder="用户名"
            class="login-input"></el-input>
      </el-form-item>
      <el-form-item label="密码" class="login-form-item">
        <el-input
            v-model="login.password"
            placeholder="密码"
            type="password"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="doLogin">登录</el-button>
        <el-button @click="()=>{this.$router.back()}">返回</el-button>
        <el-button @click="()=>{this.$router.push('/register')}">注册</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "login",
  data() {
    return {
      login: {
        "username": "",
        "password": ""
      }
    }
  },
  methods : {
    doLogin() {
      Axios({
        method: "post",
        url: "/api/user/login",
        data: JSON.stringify(this.login),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
        },
      }).then((res) => {
        if (!res.data.ok) {
          this.$alert("登录失败")
          return
        }
        sessionStorage.setItem("username", res.data.result.username);
        sessionStorage.setItem("user_id", res.data.result.user_id);
        sessionStorage.setItem("is_teacher", res.data.result.is_teacher);
        sessionStorage.setItem("login_token", res.data.result.login_token);
        this.$router.back()
      })
    }
  }
}
</script>

<style scoped>
#oj-login {
  min-height: 1000px;
  background-color: #FDFDFD;
}

.login-form {
  width: min-content;
  margin: auto;
  border-color: #aaa;
  border-radius: 10px;
  border-style: solid;
  border-width: 2px;
}

.login-form-item {
  height: 60px;
  margin: 10px 30px 10px 10px;
}

.login-input {
  width: 300px;
}

</style>