<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">

        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";

export default {
    name: "Exam",
    components: { Loading, HeadBar, FootBar },
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
  
<style scoped></style>
  
