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
import { init } from "../plugins/common"

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
        jump(name) {
            this.$router.push('/' + name)
        }
    },
    mounted() {
        init(this)
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
  
