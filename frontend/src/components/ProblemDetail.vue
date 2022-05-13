<template>
  <div id="oj-problem-detail">
    <h1>题目描述</h1>
    <p>{{this.problem_detail.description}}</p>
    <h1>输入描述</h1>
    <p>{{this.problem_detail.input_desc}}</p>
    <h1>输出描述</h1>
    <p>{{this.problem_detail.output_desc}}</p>

    <div v-for="item in this.problem_detail.sample_cases">
      <h1>输入样例</h1>
      <p class="code">{{item.input}}</p>
      <h1>输出样例</h1>
      <p class="code">{{item.output}}</p>
    </div>
    <div style="max-width: 70%">
      <el-select v-model="submission.code_language" placeholder="请选择语言" style="margin:auto; width: 100%">
        <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-input style="font-family: monospace; font-size: 12pt; width: 100%" v-model="submission.code_text"
                type="textarea" :rows="15" placeholder="code here"></el-input>
      <el-button style="width: 200px; margin-top: 20px; margin-bottom: 20px; background-color: #2d8cf0"
                 type="primary" @click="handleSubmit">提交</el-button>
    </div>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "ProblemDetail",
  data() {
    this.loadProblemDetail()

    return {
      "problem_detail": {
        "problem_id": "",
        "problem_no": "",
        "title": "",
        "author_id": "",
        "description": "",
        "input_desc": "",
        "output_desc": "",
        "sample_cases": [{
          "input": "",
          "output": ""
        }]
      },
      "options": [{
        value: 'c',
        label: 'C',
      },{
        value: 'cpp',
        label: "C++",
      },{
        value: "cpp14",
        label: "C++14"
      }],
      "submission": {
        "problem_id": "1234",
        "code_language": "",
        "code_text": ""
      }
    }
  },
  methods : {
    loadProblemDetail() {
      const id = this.$route.params.problem_id;
      Axios({
        method: "post",
        url: "/api/problem/query_detail",
        data: JSON.stringify({"problem_id": id}),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
        }
      }).then((res)=>{
        this.problem_detail = res.data.result
      })
    },
    handleSubmit() {
      this.submission.problem_id = this.problem_detail.problem_id

      if (this.submission.code_language === "") {
        this.$message({type: "warning", message: "请选择语言"})
        return
      }
      // this.$alert()
      Axios({
        method: "post",
        url: "/api/submission/submit",
        data: JSON.stringify({
          "problem_id": this.$route.params.problem_id,
          "code_text": this.submission.code_text,
          "code_language": this.submission.code_language
        }),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
          'token': sessionStorage.getItem("login_token")
        }
      }).then((result) => {
        const r = result.data;
        if (!r.ok) {
          this.$alert(r.err_info)
          return
        }

        this.$message({type: "success", message: "提交成功"})
        this.$router.push("/submissions")
      })
    }
  }
}
</script>

<style scoped>
#oj-problem-detail {
  background: #FDFDFD;
  width: 70%;
  margin: auto;
}

#oj-problem-detail * {
  background: #ffff;
  /*width: 70%;*/
  margin-left: auto; margin-right: auto;
}

#oj-problem-detail h1 {
  margin-left: 15px;
  color: #2d8cf0;
}

#oj-problem-detail p {
  margin-left: 30px;
  padding-bottom: 30px;
}
</style>