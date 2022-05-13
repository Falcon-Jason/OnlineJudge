<template>
  <div id="oj-header">
        <div class="grid-content-logo">在线编程评测系统</div>
        <router-link
            class="grid-content-list"
            to="/problems"
        >题目</router-link>
        <router-link
            class="grid-content-list"
            to="/submissions"
        >提交</router-link>
        <el-input
            placeholder="搜索题目"
            style="float: left; width: 200px"
            v-model="search"
            @change="doSearch"
        ></el-input>
        <router-link
            class="grid-content-user"
            to="/login"
            v-if="!isLoggedIn()"
        >登录</router-link>
        <router-link
            class="grid-content-user"
            to="/register"
            v-if="!isLoggedIn()"
        >注册</router-link>
        <a
            class="grid-content-user"
            href="#"
            @click="doUnregister"
            v-if="isLoggedIn()"
        >{{this.username}}</a>
  </div>
</template>

<script>
export default {
  name: "header",
  data() {
    return {
      search: "",
      user_id: sessionStorage.getItem("user_id"),
      username: sessionStorage.getItem("username"),
      is_teacher: sessionStorage.getItem("is_teacher"),
      login_token: sessionStorage.getItem("login_token")
    }
  },
  methods: {
    doRefresh() {
      this.user_id = sessionStorage.getItem("user_id")
      this.username = sessionStorage.getItem("username")
      this.is_teacher = sessionStorage.getItem("is_teacher")
      this.login_token = sessionStorage.getItem("login_token")
    },
    doUnregister() {
      sessionStorage.removeItem("username")
      sessionStorage.removeItem("user_id")
      sessionStorage.removeItem("is_teacher")
      sessionStorage.removeItem("login_token")
      this.doRefresh()
    },
    isLoggedIn() {
      return this.login_token
    },
    doSearch(search) {
      this.$router.push("/problem_search/"+search)
    }
  }
}
</script>

<style scoped>
#oj-header {
  background-color: #FDFDFD;
  color: #333;
  text-align: left;
  line-height: 60px;
}

.grid-content-logo {
  font-size: 20px;
  border-radius: 4px;
  min-height: 36px;
  text-align: left;
  text-decoration: none;
  float: left;
  margin-left: 20px;
  margin-right: 20px;
}

.grid-content-list {
  font-size: 16px;
  border-radius: 4px;
  min-height: 36px;
  text-align: left;
  color: #2d8cf0;
  text-decoration: none;
  margin-left: 20px;
}

.grid-content-user {
  font-size: 20px;
  border-radius: 4px;
  min-height: 36px;
  text-align: left;
  text-decoration: none;
  color: #333333;
  float: right;
  margin-right: 30px;
}
</style>