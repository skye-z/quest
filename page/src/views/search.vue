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
                搜索
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
import { init,getSubjectList } from "../plugins/common"

export default {
    name: "Search",
    components: { Loading, HeadBar, FootBar, Warning24Filled },
    data: () => ({
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        user: {},
        subject: 0,
        subjects: []
    }),
    methods: {
        jump(name) {
            this.$router.push('/' + name)
        },
        selectSubject(value) {
            this.subject = value
        }
    },
    mounted() {
        init(this);
        getSubjectList(this,subject)
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
  
