<template>
    <div>
        <div class="card mb-10">
            <div class="flex align-center float-right mt-5 mr-5">
                <n-button size="small" strong secondary circle class="mr-5">
                    <template #icon>
                        <n-icon>
                            <SearchRound />
                        </n-icon>
                    </template>
                </n-button>
                <n-button @click="add" size="small" type="primary" round>
                    <template #icon>
                        <n-icon>
                            <AddCircle24Regular />
                        </n-icon>
                    </template>
                    添加
                </n-button>
            </div>
            <div class="card-name">题目</div>
            <n-spin :show="loading">
                <n-scrollbar style="height: 300px">
                    <div class="question-item flex align-center justify-between" v-for="item in list" @click="edit(item)">
                        <div>
                            <div class="line1">{{ item.question }}</div>
                            <n-tag v-if="item.level == 1 || item.level == 2" :bordered="item.level == 1" size="small"
                                type="success">{{ item.level == 1 ? '极易' : '容易' }}</n-tag>
                            <n-tag v-else-if="item.level == 3" size="small" type="warning">一般</n-tag>
                            <n-tag v-else-if="item.level == 4 || item.level == 5" :bordered="item.level == 4" size="small"
                                type="error">{{ item.level == 4 ? '较难' : '极难' }}</n-tag>
                            <n-tag class="ml-5" :bordered="false" type="info" size="small">{{ options.type[item.type -
                                1].label }}</n-tag>
                            <n-text depth="3" class="ml-10">
                                <n-ellipsis style="max-width: 190px">#{{ item.subject }}</n-ellipsis>
                            </n-text>
                        </div>
                        <n-button @click.stop="remove(item.id, item.question)" strong secondary circle type="error">
                            <template #icon>
                                <n-icon>
                                    <DeleteForeverRound />
                                </n-icon>
                            </template>
                        </n-button>
                    </div>
                </n-scrollbar>
            </n-spin>
        </div>
        <n-modal v-model:show="form.show" preset="card" :style="{ width: '500px' }" :mask-closable="false">
            <template #header>
                {{ form.id ? '编辑题目' : '添加题目' }}
            </template>
            <n-form ref="form" :disabled="form.loading" :model="form" :rules="rules">
                <n-form-item path="question" label="题干">
                    <n-input v-model:value="form.question" maxlength="110" type="textarea"
                        placeholder="为保证显示效果, 限110个字符以内" />
                </n-form-item>
                <div class="flex align-center">
                    <n-form-item class="full-width" path="sid" label="科目">
                        <n-select v-model:value="form.sid" placeholder="请选择科目" :options="subjects" />
                    </n-form-item>
                    <n-form-item class="ml-10" path="type" label="题型" style="min-width: 90px;width: 160px;">
                        <n-select v-model:value="form.type" @update:value="updateType" placeholder="请选择"
                            :options="options.type" />
                    </n-form-item>
                    <n-form-item class="ml-10" path="level" label="难度" style="width: 160px;">
                        <n-select v-model:value="form.level" placeholder="请选择" :options="options.level" />
                    </n-form-item>
                </div>
                <!-- 单选题||多选题 -->
                <template v-if="form.type == 1 || form.type == 2">
                    <div class="mb-10">
                        <n-button @click="addOps" :disabled="form.loading" class="float-right" size="small"
                            quaternary>+</n-button>
                        <div>选项</div>
                    </div>
                    <div class="form-ops mb-10">
                        <div v-if="form.options.length > 0" class="form-ops-item flex align-center mt-10"
                            v-for="(_item, index) in form.options">
                            <n-text class="form-ops-index text-center mr-5" type="success">{{ index + 1 }}</n-text>
                            <n-input v-model:value="form.options[index].label" :disabled="form.loading" type="text"
                                placeholder="请输入选项内容" />
                            <n-button @click="delOps(index)" :disabled="form.loading" class="ml-10" size="small"
                                type="error" secondary strong>-</n-button>
                        </div>
                        <div v-else class="text-center"><n-text depth="3">请点击右侧加号添加选项</n-text></div>
                    </div>
                    <n-form-item path="answer" label="答案">
                        <n-select v-model:value="form.answer" :multiple="form.type == 2" placeholder="请选择答案" clearable
                            :options="form.options" />
                    </n-form-item>
                </template>
                <!-- 判断题 -->
                <template v-if="form.type == 3">
                    <n-form-item path="answer" label="答案">
                        <n-radio-group v-model:value="form.answer" name="answer">
                            <n-radio-button value="1">
                                正确
                            </n-radio-button>
                            <n-radio-button value="0">
                                错误
                            </n-radio-button>
                        </n-radio-group>
                    </n-form-item>
                </template>
                <!-- 填空题 -->
                <template v-if="form.type == 4">
                    <div class="mb-10">
                        <n-button @click="addAnswer" :disabled="form.loading" class="float-right" size="small"
                            quaternary>+</n-button>
                        <div>答案</div>
                    </div>
                    <div class="form-ops mb-10">
                        <div v-if="form.answer.length > 0" class="form-ops-item flex align-center mt-10"
                            v-for="(_item, index) in form.answer">
                            <n-text class="form-ops-index text-center mr-5" type="success">{{ index + 1 }}</n-text>
                            <n-input v-model:value="form.answer[index]" :disabled="form.loading" type="text"
                                :placeholder="'请输入第 ' + (index + 1) + ' 空答案内容'" />
                            <n-button @click="delAnswer(index)" :disabled="form.loading" class="ml-10" size="small"
                                type="error" secondary strong>-</n-button>
                        </div>
                        <div v-else class="text-center"><n-text depth="3">请点击右侧加号添加每空答案</n-text></div>
                    </div>
                </template>
                <!-- 对号题 -->
                <template v-if="form.type == 5">
                    <n-form-item path="options" label="A组(请按正确顺序)">
                        <n-dynamic-tags v-model:value="form.options" />
                    </n-form-item>
                    <n-form-item path="answer" label="B组(请按正确顺序)">
                        <n-dynamic-tags v-model:value="form.answer" />
                    </n-form-item>
                </template>
                <!-- 简答题||论述题 -->
                <template v-if="form.type == 6 || form.type == 7">
                    <n-form-item path="answer" label="关键词">
                        <n-dynamic-tags v-model:value="form.answer" />
                    </n-form-item>
                </template>
                <!-- 作文 -->
                <template v-if="form.type == 8">
                    <n-form-item path="answer" label="最低字数">
                        <n-input-number v-model:value="form.answer" clearable />
                    </n-form-item>
                </template>
            </n-form>
            <div class="flex justify-between align-center">
                <n-text type="error">* 切换题型会清空选项和答案!!!</n-text>
                <n-button :loading="form.loading" type="primary" @click="submit">
                    <template #icon>
                        <n-icon>
                            <CheckRound />
                        </n-icon>
                    </template>
                    完成{{ form.id ? '编辑' : '添加' }}
                </n-button>
            </div>
        </n-modal>
    </div>
</template>
<script>
import { DeleteForeverRound, CheckRound, SearchRound } from '@vicons/material'
import { AddCircle24Regular } from '@vicons/fluent'
import { question } from '../../plugins/api'
export default {
    components: { DeleteForeverRound, CheckRound, SearchRound, AddCircle24Regular },
    props: {
        subjects: {
            type: Array,
            default: []
        },
    },
    data: () => ({
        loading: true,
        list: [],
        form: {
            show: false,
            loading: false,
        },
        rules: {
            question: {
                required: true,
                message: '题干不能为空',
                type: 'string',
                trigger: 'blur'
            },
            sid: {
                required: true,
                message: "科目不能为空",
                type: 'number',
                trigger: ["blur", "change"]
            },
            type: {
                required: true,
                message: "题型不能为空",
                type: 'number',
                trigger: ["blur", "change"]
            },
            level: {
                required: true,
                message: "难度不能为空",
                type: 'number',
                trigger: ["blur", "change"]
            },
        },
        options: {
            form: [],
            level: [
                {
                    label: "极易",
                    value: 1
                },
                {
                    label: "容易",
                    value: 2
                },
                {
                    label: "一般",
                    value: 3
                },
                {
                    label: "较难",
                    value: 4
                },
                {
                    label: "极难",
                    value: 5
                }
            ],
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
            ]
        }
    }),
    methods: {
        getList() {
            this.loading = true
            question.getList(null, null, 1, 10).then(res => {
                if (res.list) {
                    for (let i in res.list) {
                        res.list[i].subject = this.getSubject(res.list[i].sid)
                    }
                    setTimeout(() => {
                        this.list = res.list
                        this.loading = false
                    }, 500)
                    console.log(this.list)
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(() => {

            })
        },
        getSubject(sid) {
            for (let i in this.subjects) {
                if (this.subjects[i].value == sid) return this.subjects[i].label
            }
            return '未知科目'
        },
        add() {
            this.form = {
                show: true,
                loading: false,
                sid: null,
                level: null,
                type: null,
                question: '',
                options: null,
                answer: null
            }
        },
        edit(item) {
            if (item.options != null && item.options.indexOf('[') != -1) item.options = JSON.parse(item.options)
            if (item.answer != null) {
                if (item.answer.indexOf('[') != -1) item.answer = JSON.parse(item.answer)
                else if (item.answer.indexOf('1') == 0 && item.answer.length == 13) item.answer = parseInt(item.answer)
            }
            this.form = {
                show: true,
                loading: false,
                ...item
            }
            console.log(item)
        },
        updateType(type) {
            this.form.options = []
            if (type == 4 || type == 5 || type == 6 || type == 7) this.form.answer = []
            else this.form.answer = null
        },
        addOps() {
            this.form.options.push({
                label: '',
                value: new Date().getTime()
            })
        },
        delOps(index) {
            this.form.options.splice(index, 1)
        },
        addAnswer() {
            this.form.answer.push('')
        },
        delAnswer(index) {
            this.form.answer.splice(index, 1)
        },
        checkForm(form) {
            if (form.type == 1 || form.type == 2) {
                if (form.options == null || form.options.length < 2) {
                    window.$message.warning('选项最少2项')
                    return false
                }
                if (form.type == 2 && (form.answer == null || form.answer.length < 2)) {
                    window.$message.warning('多选题答案最少2个')
                    return false
                }
            } else if (form.type == 3) {
                if (form.answer == null) {
                    window.$message.warning('答案不能为空')
                    return false
                }
            } else if (form.type == 4) {
                if (form.answer == null || form.answer.length < 1) {
                    window.$message.warning('答案最少1个空')
                    return false
                }
            } else if (form.type == 5) {
                if (form.options == null || form.options.length < 1) {
                    window.$message.warning('A组不能为空')
                    return false
                }
                if (form.answer == null || form.answer.length < 1) {
                    window.$message.warning('B组不能为空')
                    return false
                }
            } else if (form.type == 6 || form.type == 7) {
                if (form.answer == null || form.answer.length < 1) {
                    window.$message.warning('关键词不能为空')
                    return false
                }
            } else if (form.type == 8) {
                if (form.answer == null) {
                    window.$message.warning('最低字数不能为空')
                    return false
                }
            }
            if (form.options != null && form.options instanceof Array) form.options = JSON.stringify(form.options)
            if (form.answer != null) {
                if (form.answer instanceof Array) form.answer = JSON.stringify(form.answer)
                else if (typeof (form.answer) === 'number') form.answer = '' + form.answer
            }
            return true
        },
        submit() {
            this.$refs.form.validate((errors) => {
                if (!errors) {
                    let form = JSON.parse(JSON.stringify(this.form));
                    if (!this.checkForm(form)) return false;
                    this.form.loading = true
                    if (this.form.id) this.submitEdit(form)
                    else this.submitAdd(form)
                }
            })
        },
        submitAdd(form) {
            question.add(form.sid, form.level, form.type, form.question, form.options, form.answer).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('题目添加成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '题目添加失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        submitEdit(form) {
            question.edit(form.id, form.sid, form.level, form.type, form.question, form.options, form.answer).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('题目编辑成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '题目编辑失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        remove(id, name) {
            window.$dialog.warning({
                title: '操作确认',
                content: '确认要删除”' + name + '“题目?',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    question.remove(id).then(res => {
                        if (res.state) {
                            window.$message.success('删除成功')
                            this.getList()
                        } else window.$message.warning(res.message ? res.message : '删除失败')
                    }).catch(() => {
                        this.form.loading = false
                    })
                }
            })
        },
    },
    mounted() {
        setTimeout(() => {
            this.getList()
        }, 500);
    }
};
</script>
<style scoped>
.question-list {
    height: 300px;
}

.form-ops {
    min-height: 34px;
}

.form-ops-index {
    width: 24px;
}

.question-item {
    border-top: #556276 1px solid;
    cursor: pointer;
    padding: 10px;
}

.question-item:hover {
    background-color: #45546a;
}
</style>