<template>
  <div>
    <el-container class="oj-problem-list-main">
      <el-header style="height: 0"></el-header>
      <h2 style="text-align: center;">提交</h2>
      <el-table :data="submissionList" style="margin: auto; width: 100%">
        <el-table-column prop="submission_id" label="提交编号" width="200"></el-table-column>
        <el-table-column prop="problem_id" label="题目编号"></el-table-column>
        <el-table-column prop="code_language" label="使用语言"></el-table-column>
        <el-table-column prop="status" label="状态"></el-table-column>
      </el-table>
    </el-container>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "Submit",
  data() {
    this.loadSubmissionList()

    return {
      submissionList: [{
        submission_id: "",
        author_id: "",
        problem_id: "",
        code_language: "",
        status: ""
      }]
    }
  },
  methods: {
    loadSubmissionList() {
      Axios({
        method: "post",
        url: "/api/submission/search_submission",
        data: JSON.stringify({}),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
          'token': sessionStorage.getItem("login_token")
        },
      }).then((res)=>{
        const d = res.data
        if (!d.ok) {
          this.$alert(d)
          return
        }
        console.log(d)
        this.submissionList = d.result
      })
    },
  }
}
</script>

<style scoped>

</style>