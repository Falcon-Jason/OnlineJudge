<template>
  <div id="oj-problem-list">
    <el-container class="oj-problem-list-main">
      <el-header style="height: 0"></el-header>
      <h2 style="text-align: center;">搜索结果</h2>
      <el-table :data="problemList" style="margin: auto; width: 100%">
        <el-table-column prop="problem_no" label="题目编号" width="200"></el-table-column>
        <el-table-column prop="title" label="标题"></el-table-column>
        <el-table-column label="操作" width="240">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleSubmit(scope.$index)">提交</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-container>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "ProblemSearch",
  data() {
    this.loadProblemList()

    return {
      search: false,
      problemList: [{
        problem_id: '',
        problem_no: '',
        title: '',
      }],
    }
  },
  methods: {
    handleSubmit(index) {
      this.$router.push("/problem/"+this.problemList[index].problem_id)
    },
    loadProblemList() {
      const title = this.$route.params.title
      Axios({
        method: "post",
        url: "/api/problem/search",
        data: JSON.stringify({"title": title}),
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
        },
      }).then((res)=>{
        const d = res.data
        if (!d.ok) {
          this.$alert(d)
          return
        }
        console.log(d)
        this.problemList = d.result
      })
    },
  }
}
</script>

<style scoped>
#oj-problem-list {
  margin: auto;
  max-width: 1000px;
}

.oj-problem-list-main {
  background-color: #FDFDFD;
}
</style>