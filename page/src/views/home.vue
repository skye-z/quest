<template>
    <div class="no-select">
        <loading ref="loading" />
        <head-bar :app="app" :user="user" />
        <div class="app-body">
            <div class="tips text-center">请选择您要进行的操作</div>
            <div class="menu-list">
                <div @click="jump('search')" class="menu-item text-center">
                    <n-icon size="80">
                        <Search24Regular />
                    </n-icon>
                    <div class="menu-name">查题</div>
                    <div class="menu-tips">通过题干或答案查询题目</div>
                </div>
                <div @click="jump('question')" class="menu-item text-center">
                    <n-icon size="80">
                        <BookOpen24Regular />
                    </n-icon>
                    <div class="menu-name">刷题</div>
                    <div class="menu-tips">查看指定科目下的所有题目</div>
                </div>
                <div @click="jump('exam')" class="menu-item text-center">
                    <n-icon size="80">
                        <BookClock24Regular />
                    </n-icon>
                    <div class="menu-name">考试</div>
                    <div class="menu-tips">使用题库生成特定科目模拟考试</div>
                </div>
                <div v-if="user.admin" @click="jump('admin')" class="menu-item text-center">
                    <n-icon size="80">
                        <Settings24Regular />
                    </n-icon>
                    <div class="menu-name">管理</div>
                    <div class="menu-tips">管理科目、题目、考试和用户</div>
                </div>
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";
import { Settings24Regular, Search24Regular, BookOpen24Regular, BookClock24Regular } from '@vicons/fluent'
import { init } from "../plugins/common"

export default {
    name: "Home",
    components: { Loading, HeadBar, FootBar, Settings24Regular, Search24Regular, BookOpen24Regular, BookClock24Regular },
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
    font-size: 24px;
    padding: 10vh 0 50px 0;
}

.menu-list {
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    display: flex;
}

.menu-item {
    transition: all .3s ease-in;
    border-radius: 12px;
    padding: 20px 10px;
    max-width: 230px;
    cursor: pointer;
    width: 100%;
}

.menu-item:hover {
    background-color: #2a3546;
}

.menu-item:active {
    background-color: #151d29;
}

.menu-name {
    font-size: 20px;
}

.menu-tips {
    font-size: 14px;
    color: #8d8686;
}

@media (max-width:768px) {
    .menu-item {
        margin: 10px auto 0 auto;
        max-width: 170px;
    }

    .menu-tips {
        font-size: 12px;
    }
}

@media only screen and (min-width:320px) and (max-width:380px) {
    .menu-list {
        display: block;
    }

    .menu-item {
        max-width: calc(100% - 40px);
    }

    .menu-tips {
        font-size: 12px;
    }
}
</style>
  
