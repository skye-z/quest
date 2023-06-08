<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">
            <n-tabs type="bar" animated>
                <n-tab-pane name="subject" tab="科目">
                    <div class="admin-pane">
                        Wonderwall
                    </div>
                </n-tab-pane>
                <n-tab-pane name="question" tab="题目">
                    <div class="admin-pane">
                        Hey Jude
                    </div>
                </n-tab-pane>
                <n-tab-pane name="exam" tab="考试">
                    <div class="admin-pane">
                        七里香
                    </div>
                </n-tab-pane>
                <n-tab-pane name="user" tab="用户">
                    <div class="admin-pane">
                        稻香
                    </div>
                </n-tab-pane>
                <n-tab-pane name="system" tab="系统">
                    <div class="admin-pane">
                        唯一
                    </div>
                </n-tab-pane>
            </n-tabs>
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
        user: {}
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

.app-body:deep(.n-tab-pane), 
.app-body:deep(.n-tabs-tab) {
    padding-top: 0;
}

.app-body:deep(.v-x-scroll) {
    padding-left: 10px;
    width: calc(100% - 20px);
}

.admin-pane {
    background-color: #3c4a5f;
    border-radius: 8px;
    padding: 10px;
}
</style>
  
