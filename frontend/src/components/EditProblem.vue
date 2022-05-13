<template>
  <el-container style="background-color: #FDFDFD; max-width: 1000px; margin: auto">
    <el-header style="height: 0"></el-header>
    <h2 style="margin-left: 100px">编辑题目</h2>
    <el-form ref="form" :model="problem" label-width="100px" style="width: 80%; margin: auto">
      <el-form-item label="题目编号">
        <el-input v-model="problem.problem_no"></el-input>
      </el-form-item>
      <el-form-item label="标题">
        <el-input v-model="problem.title"></el-input>
      </el-form-item>
      <el-form-item label="问题描述">
        <el-input v-model="problem.description"></el-input>
      </el-form-item>
      <el-form-item label="输入描述">
        <el-input v-model="problem.input_desc"></el-input>
      </el-form-item>
      <el-form-item label="输出描述">
        <el-input v-model="problem.output_desc"></el-input>
      </el-form-item>
      <div v-for="cs in problem.sample_cases">
        <el-form-item label="样例输入">
          <el-input v-model="cs.input"></el-input>
        </el-form-item>
        <el-form-item label="样例输出">
          <el-input v-model="cs.output" style="width: 80%;"></el-input>
          <el-button style="width:19%" type="danger" @click="deleteSample(cs)">删除</el-button>
        </el-form-item>

      </div>
      <el-button style="margin: 0 0 20px 100px" type="primary" @click="addSample">添加样例</el-button>
      <div v-for="cs in problem.test_cases">
        <el-form-item label="测试输入">
          <el-input v-model="cs.input"></el-input>
        </el-form-item>
        <el-form-item label="测试输出">
          <el-input v-model="cs.output" style="width: 80%;"></el-input>
          <el-button style="width:19%" type="danger" @click="deleteTest(cs)">删除</el-button>
        </el-form-item>
      </div>
      <el-button style="margin: 0 0 20px 100px" type="primary" @click="addTest">添加测试</el-button>
    </el-form>
    <el-button style="width: 200px; margin: auto" type="primary" @click="handleSubmit">提交</el-button>
  </el-container>
</template>

<script>
import Axios from "axios";

export default {
  name: "AddProblem",
  data() {
    this.loadProblem()

    return {
      problem: {
        "problem_no": "",
        "title": "",
        "description": "",
        "input_desc": "",
        "output_desc": "",
        "sample_cases": [
          {
            "input": "",
            "output": ""
          }],
        "test_cases": [
          {
            "input": "",
            "output": ""
          }]
      }
    }
  }, methods : {
    addSample() {
      this.problem.sample_cases.push({"input": "", "output": ""})
    },
    addTest() {
      this.problem.test_cases.push({"input": "", "output": ""})
    },
    deleteSample(item) {
      const index = this.problem.sample_cases.indexOf(item)
      if (index !== -1) {
        this.problem.sample_cases.splice(index, 1)
      }
    },
    deleteTest(item) {
      const index = this.problem.test_cases.indexOf(item)
      if (index !== -1) {
        this.problem.test_cases.splice(index, 1)
      }
    },
    handleSubmit() {
      var problem = this.problem;
      problem.problem_id = this.$route.params.problem_id

      Axios({
        method: "post",
        url: "/api/problem/modify",
        data: JSON.stringify(problem),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
          'token': sessionStorage.getItem("login_token")
        }
      }).then((res) => {
        const r = res.data
        if (!r.ok) {
          this.$alert(res)
          return
        }

        this.$message({type:"success", message: "添加成功"})
        this.$router.push("/problems")
      })
    },
    loadProblem() {
      Axios({
        method: "post",
        url: "/api/problem/query",
        data: JSON.stringify({"problem_id": this.$route.params.problem_id}),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
          'token': sessionStorage.getItem("login_token")
        }
      }).then((res)=>{
        this.problem = res.data.result
      })
    }
  }
}
</script>

<style scoped>

</style>