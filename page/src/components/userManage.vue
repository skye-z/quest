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
            <div class="card-name">用户 <span class="text-gray">{{ list.length }}</span></div>
            <n-spin :show="loading">
                <n-scrollbar style="height: 300px">
                    <div class="user-item flex align-center justify-between" v-for="item in list" @click="edit(item)">
                        <div>
                            <div class="line1">
                                <span>{{ item.nickname }}</span> <n-text depth="3">@{{ item.name }}</n-text>
                            </div>
                            <n-tag size="small" :bordered="!item.admin" type="success" class="mr-5">管理权限</n-tag>
                            <n-tag size="small" :bordered="!item.edit" type="success">编辑权限</n-tag>
                        </div>
                        <n-button @click.stop="remove(item.id, item.nickname)" strong secondary circle type="error">
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
                {{ form.id ? '编辑用户' : '添加用户' }}
            </template>
            <n-form ref="form" :disabled="form.loading" :model="form" :rules="rules">
                <n-form-item path="name" label="用户名">
                    <n-input v-model:value="form.name" @keyup.enter="submit" maxlength="18"
                        placeholder="为保证显示效果, 限18个字符以内" />
                </n-form-item>
                <n-form-item path="nickname" label="昵称">
                    <n-input v-model:value="form.nickname" @keyup.enter="submit" maxlength="18" placeholder="留空将默认使用用户名" />
                </n-form-item>
                <n-form-item path="pass" label="密码">
                    <n-input v-model:value="form.pass" @keyup.enter="submit" minlength="8" maxlength="24"
                        :placeholder="'请输入8到24个字符' + (form.id ? ', 留空则不修改' : '')" />
                </n-form-item>
                <div class="flex align-center justify-between mb-20">
                    <div class="flex align-center">
                        <n-switch v-model:value="form.admin" />
                        <div class="ml-5">赋予管理权限</div>
                    </div>
                    <div class="flex align-center">
                        <n-switch v-model:value="form.edit" />
                        <div class="ml-5">赋予编辑权限</div>
                    </div>
                </div>
            </n-form>
            <div class="flex justify-between align-center">
                <n-text type="error">* 用户名不允许重名</n-text>
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
import { user } from '../plugins/api'
export default {
    components: { DeleteForeverRound, CheckRound, SearchRound, AddCircle24Regular },
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
                message: '用户名不能为空',
                type: 'string',
                trigger: 'blur'
            },
        }
    }),
    methods: {
        getList() {
            this.loading = true
            user.getList().then(res => {
                if (res.list) {
                    setTimeout(() => {
                        this.list = res.list
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
                nickname: '',
                admin: false,
                edit: false
            }
            this.rules.pass = {
                required: true,
                message: '密码不符合要求',
                type: 'string',
                min: 8,
                max: 24,
                trigger: 'blur'
            }
        },
        edit(item) {
            this.form = {
                show: true,
                loading: false,
                ...item
            }
            this.rules.pass = {}
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
            user.add(form.name, form.pass, form.nickname, form.admin, form.edit).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('用户添加成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '用户添加失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        submitEdit(form) {
            user.edit(form.id, form.name, form.pass, form.nickname, form.admin, form.edit).then(res => {
                this.form.loading = false
                if (res.state) {
                    window.$message.success('用户编辑成功')
                    this.form.show = false
                    this.getList()
                } else window.$message.warning(res.message ? res.message : '用户编辑失败')
            }).catch(() => {
                this.form.loading = false
            })
        },
        remove(id, name) {
            window.$dialog.warning({
                title: '操作确认',
                content: '确认要删除”' + name + '“用户?',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    user.remove(id).then(res => {
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
.user-item {
    border-top: #556276 1px solid;
    cursor: pointer;
    padding: 10px;
}

.user-item:hover {
    background-color: #45546a;
}
</style>