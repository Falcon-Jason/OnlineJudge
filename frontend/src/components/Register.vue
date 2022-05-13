<template>
  <div id="oj-register">
    <el-form class="login-form" label-width="80px">
      <div style="margin: auto; text-align: center; font-size: 30px;">注册</div>
      <el-form-item label="用户名" class="login-form-item">
        <el-input
            v-model="rgst.username"
            placeholder="用户名"
            class="login-input"></el-input>
      </el-form-item>
      <el-form-item label="密码" class="login-form-item">
        <el-input
            v-model="rgst.password"
            placeholder="密码"
            type="password"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="doRegister">注册</el-button>
        <el-button @click="()=>{this.$router.back()}">返回</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "Register",
  data() {
    return {
      "rgst": {
        username: "",
        password: ""
      }
    }
  },
  methods: {
    doRegister() {
      Axios({
        method: "post",
        url: "/api/user/register",
        data: JSON.stringify(this.rgst),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
        },
      }).then((res) => {
        if (!res.data.ok) {
          this.$alert("注册失败").then(()=>{this.$router.back()})
        } else {
          this.$alert("注册成功").then(()=>{this.$router.back()})
        }

      })
    }
  }
}
</script>

<style scoped>
#oj-register {
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