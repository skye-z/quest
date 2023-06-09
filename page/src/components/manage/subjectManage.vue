<template>
    <div>
        <div class="card mb-10">
            <n-button @click="add" class="float-right mt-5 mr-5" size="small" type="primary" round>
                <template #icon>
                    <n-icon>
                        <AddCircle24Regular />
                    </n-icon>
                </template>
                添加
            </n-button>
            <div class="card-name">科目 <span class="text-gray">{{ list.length }}</span></div>
            <n-spin :show="loading">
                <n-scrollbar style="height: 300px">
                    <div class="subject-item flex align-center justify-between" v-for="item in list" @click="edit(item)">
                        <div>
                            <div class="line1">{{ item.name }}</div>
                            <div v-if="item.tag.length > 0">
                                <n-tag class="mr-5" v-for="tag in item.tags" size="small" type="success"
                                    :bordered="false">{{
                                        tag }}</n-tag>
                            </div>
                        </div>
                        <n-button @click.stop="remove(item.id, item.name)" strong secondary circle type="error">
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
        <n-modal v-model:show="form.show" preset="card" :style="{ width: '320px' }" :mask-closable="false">
            <template #header>
                {{ form.id ? '编辑科目' : '添加科目' }}
            </template>
            <n-form ref="form" :disabled="form.loading" :model="form" :rules="rules">
                <n-form-item path="name" label="科目名">
                    <n-input v-model:value="form.name" @keyup.enter="submit" maxlength="30"
                        placeholder="为保证显示效果, 限30个字符以内" />
                </n-form-item>
                <n-form-item path="tag" label="标签">
                    <n-input v-model:value="form.tag" @keyup.enter="submit" placeholder="请使用逗号分隔, 切勿中英文逗号混写" />
                </n-form-item>
            </n-form>
            <div class="flex justify-between align-center">
                <n-text type="error">* 科目不允许重名</n-text>
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
import { DeleteForeverRound, CheckRound } from '@vicons/material'
import { AddCircle24Regular } from '@vicons/fluent'
import { subject } from '../../plugins/api'
export default {
    components: { DeleteForeverRound, CheckRound, AddCircle24Regular },
    data: () => ({
        loading: true,
        list: [],
        form: {
            show: false,
            loading: false,
        },
        rules: {
            name: {
                required: true,
                message: '科目名称不能为空',
                trigger: 'blur'
            },
        }
    }),
    methods: {
        getList() {
            this.loading = true
            subject.getList().then(res => {
                if (res.list) {
                    let list = [];
                    for (let i in res.list) {
                        let item = res.list[i]
                        let tags = [];
                        if (item.tag.indexOf(',') > -1) tags = item.tag.split(',')
                        else if (item.tag.indexOf('，') > -1) tags = item.tag.split(',')
                        else if (item.tag.indexOf(' ') > -1) tags = item.tag.split(' ')
                        else if (item.tag) tags.push(item.tag)
                        item.tags = tags
                        list.push(item)
                    }
                    setTimeout(() => {
                        this.list = list
                        this.loading = false
                    }, 500)
                } else {
                    window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                }
            }).catch(() => {

            })
        },
        add() {
            this.form = {
                show: true,
                loading: false,
                name: '',
                tag: ''
            }
        },
        edit(item) {
            this.form = {
                show: true,
                loading: false,
                ...item
            }
        },
        submit() {
            this.$refs.form.validate((errors) => {
                if (!errors) {
                    let form = this.form;
                    this.form.loading = true
                    if (this.form.id) this.submitEdit(form)
                    else this.submitAdd(form)
                }
            })
        },
        submitAdd(form) {
            subject.add(form.name, form.tag).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('科目添加成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '科目添加失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        submitEdit(form) {
            subject.edit(form.id, form.name, form.tag).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('科目编辑成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '科目编辑失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        remove(id, name) {
            window.$dialog.warning({
                title: '操作确认',
                content: '确认要删除《' + name + '》科目?',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    subject.remove(id).then(res => {
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
.subject-item {
    border-top: #556276 1px solid;
    cursor: pointer;
    padding: 10px;
}

.subject-item:hover {
    background-color: #45546a;
}
</style>