import Vue from "vue";
import VueRouter from "vue-router";
import IndexPage from "~/pages/IndexPage";
import ProblemListPage from "~/pages/ProblemListPage";
import Register from "~/components/Register";
import Login from "~/components/Login";
import ProblemDetailPage from "~/pages/ProblemDetailPage";
import ProblemSearch from "~/components/ProblemSearch";
import ProblemSearchPage from "~/pages/ProblemSearchPage";
import SubmissionListPage from "~/pages/SubmissionListPage";
import AddProblemPage from "~/pages/AddProblemPage";
import EditProblemPage from "~/pages/EditProblemPage";

Vue.use(VueRouter)

const routes = [
    {
        path: "/",
        component: IndexPage
    },
    {
        path: '/login',
        component: Login
    },
    {
        path: "/register",
        component: Register
    },
    {
        path: "/problems",
        component: ProblemListPage
    },
    {
        path: "/problem_search/:title",
        component: ProblemSearchPage
    },
    {
        path: "/problem_search",
        component: ProblemSearchPage
    },
    {
        path: "/problem/:problem_id",
        component: ProblemDetailPage
    }, {
        path: "/submissions",
        component: SubmissionListPage
    }, {
        path: "/add_problem",
        component: AddProblemPage
    }, {
        path: "/edit_problem/:problem_id",
        component: EditProblemPage
    }
]

const router = new VueRouter({routes})

export default router