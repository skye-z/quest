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
                <div class="question-item card mb-10" v-for="item in questions">
                    <div class="float-right flex align-center">
                        <n-tag v-if="item.level == 1 || item.level == 2" :bordered="item.level == 1" size="small"
                            type="success">{{ item.level == 1 ? '极易' : '容易' }}</n-tag>
                        <n-tag v-else-if="item.level == 3" size="small" type="warning">一般</n-tag>
                        <n-tag v-else-if="item.level == 4 || item.level == 5" :bordered="item.level == 4" size="small"
                            type="error">{{ item.level == 4 ? '较难' : '极难' }}</n-tag>
                        <n-tag class="ml-5" :bordered="false" type="info" size="small">{{ type[item.type -
                            1].label }}</n-tag>
                        <n-button size="small" quaternary circle class="ml-10">
                            <template #icon>
                                <n-icon><Eye24Filled /></n-icon>
                            </template>
                        </n-button>
                    </div>
                    <div class="question-title">{{ item.question }}</div>
                    <template v-if="item.type == 1">
                        <n-radio-group @update:value="value => check(item,value)" class="mt-10" v-model:value="form[item.id]" :name="'options-' + item.id">
                            <n-radio class="full-width" v-for="opt in item.options" :key="opt.value" :value="opt.value">
                                {{ opt.label }}
                            </n-radio>
                        </n-radio-group>
                    </template>
                    <template v-else-if="item.type == 2">
                        <n-checkbox-group @update:value="value => check(item,value)" class="mt-10" v-model:value="form[item.id]" :name="'options-' + item.id">
                            <n-checkbox class="full-width" v-for="opt in item.options" :key="opt.value" :value="opt.value"
                                :label="opt.label" />
                        </n-checkbox-group>
                    </template>
                    <template v-else-if="item.type == 3">
                        <n-radio-group @update:value="value => check(item,value)" class="mt-10" v-model:value="form[item.id]" :name="'options-' + item.id">
                            <n-radio-button value="1">
                                正确
                            </n-radio-button>
                            <n-radio-button value="0">
                                错误
                            </n-radio-button>
                        </n-radio-group>
                    </template>
                    <template v-else-if="item.type == 4">
                        <div class="question-ops">
                            <n-input v-for="(sub, index) in item.answer" @update:value="value => check(sub,value)" v-model:value="form[item.id][index]"
                                :placeholder="'请输入第 ' + (index + 1) + ' 空'" class="mr-10 mt-10" style="width: 200px;" />
                        </div>
                    </template>
                    <template v-else-if="item.type == 5">
                        <div class="flex align-center" style="flex-wrap: wrap;">
                            <div class="flex align-center mr-10 mt-10" v-for="(opt, index) in item.answer">
                                <div class="mr-10">{{ opt }}:</div>
                                <n-select @update:value="value => check(opt,value)" v-model:value="form[item.id][index]" :options="item.options"
                                    style="width: 200px;" />
                            </div>
                        </div>
                    </template>
                    <template v-else-if="item.type == 6 || item.type == 7 || item.type == 8">
                        <n-input class="mt-10" v-model:value="form[item.id]" @update:value="value => check(item,value)" type="textarea" label-field=""
                            placeholder="请在此处输入答案" />
                    </template>
                </div>
                <n-pagination v-if="number > 0" v-model:page="page" :item-count="number" />
            </div>
        </div>
        <foot-bar :app="app" />
    </div>
</template>
  
<script>
import Loading from "../components/loading.vue";
import HeadBar from "../components/headBar.vue";
import FootBar from "../components/footBar.vue";
import { Warning24Filled, Eye24Filled } from '@vicons/fluent'
import { subject, question } from '../plugins/api'
import { init, getSubjectList } from "../plugins/common"

export default {
    name: "Question",
    components: { Loading, HeadBar, FootBar, Warning24Filled, Eye24Filled },
    data: () => ({
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        user: {},
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
        form: {

        },
        subject: 0,
        subjects: [],
        questions: [],
        page: 0,
        number: 0
    }),
    methods: {
        jump(name) {
            this.$router.push('/' + name)
        },
        selectSubject(value) {
            this.subject = value;
            this.getQuestionNumber();
        },
        getQuestionNumber() {
            question.getNumber(this.subject).then(res => {
                this.number = res.num
                this.questions = []
                this.form = []
                if (res.num > 0) {
                    this.page = 1
                    this.getQuestionList()
                }
            }).catch(() => {

            })
        },
        getQuestionList() {
            question.getSubList(this.subject, this.page, 20).then(res => {
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
        check(item,value){
            // 检查答案是否正确
            console.log(item,value)
        }
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

.question-item {
    padding: 10px;
}

.question-title {
    font-weight: 600;
    font-size: 16px;
}

.question-ops {
    align-items: center;
    flex-wrap: wrap;
    display: flex;
}
</style>
  
