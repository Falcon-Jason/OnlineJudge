import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import Header from './components/Header.vue'
import ProblemList from "~/components/ProblemList";
import router from "~/router/index.js";
import ProblemDetail from "~/components/ProblemDetail";
import ProblemSearch from "~/components/ProblemSearch";
import Submit from "~/components/SubmissionList";
import SubmissionList from "~/components/SubmissionList";
import AddProblem from "~/components/AddProblem";
import EditProblem from "~/components/EditProblem";

Vue.use(ElementUI)

Vue.component("oj-header", Header)
Vue.component("oj-problem-list", ProblemList)
Vue.component("oj-problem-detail", ProblemDetail)
Vue.component("oj-problem-search", ProblemSearch)
Vue.component("oj-submit-panel", Submit)
Vue.component("oj-submission-list", SubmissionList)
Vue.component("oj-add-problem", AddProblem)
Vue.component("oj-edit-problem", EditProblem)

new Vue({
  el: "#app",
  router,
  render: h => h(App)
})
