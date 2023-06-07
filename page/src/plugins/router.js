import {
    createRouter,
    createWebHistory
} from 'vue-router'

const Home = () => import('../views/Home.vue')
const Auth = () => import('../views/Auth.vue')
const Admin = () => import('../views/Admin.vue')
const Search = () => import('../views/Search.vue')
const Question = () => import('../views/Question.vue')
const Exam = () => import('../views/Exam.vue')

const router = createRouter({
    history: createWebHistory(import.meta.env.VITE_APP_PATH),
    routes: [
        {
            name: 'Home',
            path: '/home',
            component: Home,
            meta: {
                title: "首页",
                auth: true
            }
        },
        {
            name: 'Auth',
            path: '/auth',
            component: Auth,
            meta: {
                title: "登录",
                auth: false
            }
        },
        {
            name: 'Admin',
            path: '/admin',
            component: Admin,
            meta: {
                title: "管理",
                auth: true
            }
        },
        {
            name: 'Search',
            path: '/search',
            component: Search,
            meta: {
                title: "搜题",
                auth: true
            }
        },
        {
            name: 'Question',
            path: '/question',
            component: Question,
            meta: {
                title: "刷题",
                auth: true
            }
        },
        {
            name: 'Exam',
            path: '/exam',
            component: Exam,
            meta: {
                title: "考试",
                auth: true
            }
        },
    ],
})

router.beforeEach((to, _from, next) => {
    let name = localStorage.getItem('app:config:name')
    if (name == undefined) name = 'Quest云题库'
    if (!to.meta.auth || Boolean(localStorage.getItem("user:access:token")) || to.name == "Auth") {
        document.title = to.meta.title === undefined ? '' : to.meta.title
        document.title = (to.meta.title === undefined ? '未知页面 - ' : to.meta.title + ' - ') + name
        next()
    } else {
        next('/auth')
    }
})

export default router