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
                <div class="card mb-10">
                    <n-input-group>
                        <n-input v-model:value="keyword" @keyup.enter="getQuestionNumber" type="text" placeholder="请输入关键词, 支持模糊查询题干与选项">
                            <template #prefix>
                                <n-icon>
                                    <Search24Regular />
                                </n-icon>
                            </template>
                        </n-input>
                        <n-button type="primary" @click="getQuestionNumber">
                            搜索
                        </n-button>
                    </n-input-group>
                </div>
                <div class="card search-list mb-10">
                    <n-scrollbar v-if="number > 0" style="max-height: calc(100vh - 210px)">
                        <n-table style="min-width: 700px;" :bordered="false" :single-line="false">
                            <thead>
                                <tr>
                                    <th scope="col" style="width: 70px;">编号</th>
                                    <th scope="col" style="width: 50px;text-align: center;">题型</th>
                                    <th scope="col">题干</th>
                                    <th scope="col" style="width: 30%;">选项</th>
                                    <th scope="col" style="width: 100px;">答案</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="item in questions">
                                    <td>#{{ item.id }}</td>
                                    <td class="text-center">{{ type[item.type - 1].label }}</td>
                                    <td>{{ item.question }}</td>
                                    <td>
                                        <template v-if="item.type == 1 || item.type == 2">
                                            <div v-for="(sub, index) in item.options">
                                                {{ word[index] }}. {{ sub.label }}
                                            </div>
                                        </template>
                                        <template v-else-if="item.type == 5">
                                            <div>
                                                <div>
                                                    <div>A组</div>
                                                    <div v-for="(sub, index) in item.options">
                                                        {{ word[index] }}. {{ sub.label }}
                                                    </div>
                                                </div>
                                                <div>
                                                    <div>B组</div>
                                                    <div v-for="(sub, index) in item.answer">
                                                        {{ index + 1 }}. {{ sub }}
                                                    </div>
                                                </div>
                                            </div>
                                        </template>
                                        <span class="text-gray" v-else-if="item.type != 5">
                                            N/A
                                        </span>
                                    </td>
                                    <td>
                                        <template v-if="item.type == 1">
                                            <div v-for="(opt, index) in item.options">
                                                <div v-if="opt.value == item.answer">{{ word[index] }}</div>
                                            </div>
                                        </template>
                                        <template v-else-if="item.type == 2">
                                            <span v-for="(opt, x) in item.options">
                                                <span v-for="sub in item.answer">
                                                    <span class="mr-5" v-if="opt.value == sub">{{ word[x] }}</span>
                                                </span>
                                            </span>
                                        </template>
                                        <template v-else-if="item.type == 3">
                                            {{ item.answer == '1' ? '正确' : '错误' }}
                                        </template>
                                        <template v-else-if="item.type == 4">
                                            <div v-for="(sub, index) in item.answer">
                                                <div>{{ index + 1 }}. {{ sub }}</div>
                                            </div>
                                        </template>
                                        <template v-else-if="item.type == 5">见选项顺序</template>
                                        <template v-else-if="item.type == 6 || item.type == 7">
                                            <span v-for="(sub, index) in item.answer">
                                                <span v-if="index != 0">、</span>{{ sub }}
                                            </span>
                                        </template>
                                        <span class="text-gray" v-if="item.type == 8">
                                            N/A
                                        </span>
                                    </td>
                                </tr>
                            </tbody>
                        </n-table>
                    </n-scrollbar>
                    <div v-else class="tips">
                    <n-result status="500" title="没找到" description="没有找到符合的试题啊">
                        <template #footer>
                            <n-button @click="getQuestionNumber">重新获取</n-button>
                        </template>
                    </n-result>
                    </div>
                </div>
                <n-pagination v-if="number > 0" class="justify-center" v-model:page="page" :item-count="number" />
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";
import { Search24Regular, Warning24Filled } from '@vicons/fluent'
import { subject, question } from '../plugins/api'
import { init, getSubjectList } from "../plugins/common"

export default {
    name: "Search",
    components: { Loading, HeadBar, FootBar, Search24Regular, Warning24Filled },
    data: () => ({
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        type: [
            {
                label: "单选题",
                value: 1
            },
            {
                label: "多选题",
                value: 2
            },
            {
                label: "判断题",
                value: 3
            },
            {
                label: "填空题",
                value: 4
            },
            {
                label: "对号题",
                value: 5
            },
            {
                label: "简答题",
                value: 6
            },
            {
                label: "论述题",
                value: 7
            },
            {
                label: "作文题",
                value: 8
            }
        ],
        word: ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'],
        keyword: '',
        user: {},
        subject: 0,
        subjects: [],
        questions: [],
        page: 0,
        number: 1
    }),
    methods: {
        jump(name) {
            this.$router.push('/' + name)
        },
        selectSubject(value) {
            this.subject = value
            this.getQuestionNumber();
        },
        getQuestionNumber() {
            question.getNumber(this.subject, this.keyword).then(res => {
                this.number = res.num
                this.checkState = {}
                this.questions = []
                this.form = []
                if (res.num > 0) {
                    this.page = 1
                    this.getQuestionList()
                } else window.$message.warning('未找到可用试题')
            }).catch(() => {

            })
        },
        getQuestionList() {
            question.getList(this.subject, this.keyword, this.page, 20).then(res => {
                for (let i in res.list) {
                    let item = res.list[i]
                    if (item.options.indexOf('[') != -1) item.options = JSON.parse(item.options)
                    if (item.answer.indexOf('[') != -1) item.answer = JSON.parse(item.answer)
                    if (item.type == 4 || item.type == 5) {
                        let ops = []
                        for (let x in item.options) {
                            ops.push({
                                label: item.options[x],
                                value: x
                            })
                        }
                        item.options = ops
                        this.form[item.id] = []
                    }
                }
                this.questions = res.list
            }).catch(() => {

            })
        },
    },
    mounted() {
        init(this);
        getSubjectList(this, subject)
    }
};
</script>
  
<style scoped>
.tips {
    padding-top: 30vh;
    font-size: 24px;
    width: 100%;
}

.search-list {
    height: calc(100vh - 210px);
    width: 100%;
}

.search-id {
    min-width: 60px;
    width: 60px;
}

.search-title {
    width: 200px;
}
</style>
  
