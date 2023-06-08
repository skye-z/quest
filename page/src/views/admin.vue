<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">
            <div>
                管理
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

export default {
    name: "Admin",
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
        }
    }),
    methods: {
        init() {
            this.app = {
                name: localStorage.getItem('app:config:name'),
                version: localStorage.getItem('app:config:version')
            }
            let user = localStorage.getItem('user:info')
            if (user) this.user = JSON.parse(user)
            setTimeout(() => {
                this.$refs.loading.hide()
            }, 500);
        },
        jump(name) {
            this.$router.push('/' + name)
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
  
