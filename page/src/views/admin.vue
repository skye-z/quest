<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">
            <system-manage ref="system" />
            <div class="admin-box flex">
                <div class="admin-left full-width">
                    <subject-manage />
                    <user-manage />
                </div>
                <div class="admin-right full-width">
                    <question-manage />
                    <exam-manage />
                </div>
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
<script>
import Loading from "../components/loading.vue"
import HeadBar from "../components/headBar.vue"
import FootBar from "../components/footBar.vue"
import SystemManage from "../components/manage/systemManage.vue"
import SubjectManage from "../components/manage/subjectManage.vue"
import UserManage from "../components/manage/userManage.vue"
import QuestionManage from "../components/manage/questionManage.vue"
import ExamManage from "../components/manage/examManage.vue"
import { init } from "../plugins/common"

export default {
    name: "Admin",
    components: { Loading, HeadBar, FootBar, SystemManage, SubjectManage, UserManage, QuestionManage, ExamManage },
    data: () => ({
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        user: {}
    }),
    methods: {

    },
    mounted() {
        init(this)
    },
    beforeRouteLeave() {
        this.$refs.system.stop()
    }
};
</script>
  
<style scoped>
.admin-left {
    margin-right: 10px;
    max-width: 292px;
    min-width: 292px;
}

.admin-right {}
@media (max-width:670px) {
    .admin-left {
        max-width: none;
    }

    .admin-box {
        display: block;
    }
}
</style>
  
