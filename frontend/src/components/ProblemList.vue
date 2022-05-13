<template>
  <div id="oj-problem-list">
  <el-container class="oj-problem-list-main">
    <el-header style="height: 0"></el-header>
    <h2 style="text-align: center;">题目列表</h2>
    <el-table :data="problemList" style="margin: auto; width: 100%">
      <el-table-column prop="problem_no" label="题目编号" width="200"></el-table-column>
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column label="操作" width="240">
        <template slot-scope="scope">
          <el-button size="mini" type="primary" @click="handleSubmit(scope.$index)">提交</el-button>
          <el-button size="mini" @click="handleEdit(scope.$index)"
                     v-if="isOwner(problemList[scope.$index])">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index)"
                     v-if="isOwner(problemList[scope.$index])">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button type="primary" v-if="isTeacher()" @click="handleAddProblem"
               style="width: 100px; margin: 10px 10px 10px 10px">添加题目</el-button>
  </el-container>
  <el-container>
    <el-pagination
        style="margin-top: 30px; margin-bottom: 30px"
        background
        layout="prev, pager, next"
        :total="1000"
        v-model="listRequest.index"
        @current-change="onPageChanged">
    </el-pagination>
  </el-container>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  name: "ProblemList",
  data() {
    this.loadProblemList()

    return {
      search: false,
      problemList: [{
        problem_id: '',
        problem_no: '',
        author_id: '',
        title: '',
      }],
      listRequest: {
        index: 1,
      },
    }
  },
  methods: {
    handleSubmit(index) {
      this.$router.push("/problem/" + this.problemList[index].problem_id)
    },
    handleEdit(index) {
      this.$router.push("/edit_problem/" + this.problemList[index].problem_id)
    },
    handleAddProblem() {
      this.$router.push("/add_problem")
    },
    handleDelete(index) {
      // this.$alert("delete: " + this.problemList[index].problem_id)
      this.$confirm("是否确认删除题目", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        Axios({
          method: "post",
          url: "/api/problem/delete",
          data: JSON.stringify({"problem_id": this.problemList[index].problem_id}),
          headers: {
            'Content-Type': 'application/json;charset=UTF-8',
            'token': sessionStorage.getItem("login_token")
          },
        }).then((res) => {
          if (res.data.ok) {
            this.$message({type:"success", message: "删除成功"})
          } else {
            this.$alert("删除失败: " + res.data.err_info)
            return
          }
          this.loadProblemList()
        }).catch((err) => {
          this.$message({type: "warning", message: err})
        })
      }).catch(() => {
        this.$message({type: 'info', message: "已取消删除"})
      })
    },
    isOwner(problem) {
      // return true
      return problem.author_id === sessionStorage.getItem("user_id")
    },
    isTeacher() {
      return sessionStorage.getItem('is_teacher') === "true"
    },
    loadProblemList() {
      Axios({
        method: "post",
        url: "/api/problem/list",
        data: JSON.stringify(this.listRequest ? this.listRequest : {"index": 1}),
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
    onPageChanged(index) {
      this.listRequest.index = index
      this.loadProblemList()
    }
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