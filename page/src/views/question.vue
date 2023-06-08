<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" :subjects="subjects" @update="selectSubject" />
        <div class="app-body">
            <div class="tips text-center" v-if="subject == 0">
                <n-icon class="mr-5" size="100">
                    <Warning24Filled />
                </n-icon>
                <div>请先在左上方选择科目</div>
            </div>
            <div v-else>
                刷题
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";
import { Warning24Filled } from '@vicons/fluent'
import { subject } from '../plugins/api'

export default {
    name: "Question",
    components: { Loading, HeadBar, FootBar, Warning24Filled },
    data: () => ({
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        user: {
            id: 0,
            nickname: '',
            name: '',
            admin: false,
            edit: false
        },
        subject: 0,
        subjects: []
    }),
    methods: {
        init() {
            this.app = {
                name: localStorage.getItem('app:config:name'),
                version: localStorage.getItem('app:config:version')
            }
            let user = localStorage.getItem('user:info')
            if (user) this.user = JSON.parse(user)
            this.getSubjectList()
            setTimeout(() => {
                this.$refs.loading.hide()
            }, 500);
        },
        jump(name) {
            this.$router.push('/' + name)
        },
        getSubjectList() {
            subject.getList().then(res => {
                if (res.list) {
                    let list = []
                    for (let i in res.list) {
                        let item = res.list[i]
                        list[i] = {
                            label: item.name,
                            value: item.id
                        }
                    }
                    this.subjects = list
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(() => {

            })
        },
        selectSubject(value) {
            this.subject = value
        }
    },
    mounted() {
        this.init()
    }
};
</script>
  
<style scoped>
.tips {
    padding-top: 30vh;
    font-size: 24px;
    width: 100%;
}
</style>
  
