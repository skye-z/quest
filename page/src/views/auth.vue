<template>
    <div id="auth-box" class="no-select">
        <div class="text-center">
            <div id="app-name">{{ app.name }}</div>
            <div id="auth-card">
                <n-form ref="form" size="large" :show-require-mark="false" :model="form" :rules="rules">
                    <n-form-item label="用户名" path="name">
                        <n-input v-model:value="form.name" :disabled="form.loading" placeholder="请输入用户名" type="text" />
                    </n-form-item>
                    <n-form-item label="密码" path="pass">
                        <n-input v-model:value="form.pass" @keyup.enter="submit" :disabled="form.loading"
                            placeholder="请输入密码" type="password" />
                    </n-form-item>
                </n-form>
                <div id="auth-tools" class="flex justify-between align-center">
                    <div class="flex align-center text-gray">
                        <n-switch v-model:value="form.remember" :disabled="form.login" class="mr-5" /> 记住我
                    </div>
                    <n-button :loading="form.loading" @click="submit" text-color="#1f2937" id="auth-btn" type="primary">
                        <template #icon>
                            <n-icon>
                                <AirplaneTakeOff16Filled />
                            </n-icon>
                        </template>
                        登录
                    </n-button>
                </div>
                <div id="auth-tips" class="text-center text-gray">忘记密码或注册账户请联系管理员</div>
            </div>
        </div>
        <div id="auth-foot" class="text-center text-gray">Powered by <a href="https://github.com/skye-z/quest"
                target="_blank">Quest</a> v{{ app.version }}</div>
    </div>
</template>
  
<script>
import { AirplaneTakeOff16Filled } from '@vicons/fluent'
import { user } from '../plugins/api'

export default {
    name: "Auth",
    components: { AirplaneTakeOff16Filled },
    data: () => ({
        loading: true,
        app: {
            name: 'Quest云题库',
            version: '1.0.0'
        },
        form: {
            name: '',
            pass: '',
            remember: false,
            loading: false
        },
        rules: {
            name: {
                required: true,
                message: '请输入用户名',
                trigger: 'blur'
            },
            pass: {
                required: true,
                message: '请输入密码',
                trigger: 'blur'
            },
        }
    }),
    methods: {
        init() {
            this.app = {
                name: localStorage.getItem('app:config:name'),
                version: localStorage.getItem('app:config:version')
            }
            let last = localStorage.getItem('user:last:login')
            if (last) {
                this.form.name = last
                this.form.remember = true
            }
            this.loading = false;
        },
        submit() {
            this.form.loading = true
            this.$refs.form.validate((errors) => {
                if (!errors) this.login()
            })
        },
        login() {
            user.login(this.form.name, this.form.pass).then(res => {
                if (res.token && res.user) {
                    window.$message.success('欢迎回来, ' + res.user.nickname + (res.user.admin ? '老师' : '同学') + ', 正在跳转中');
                    localStorage.setItem('user:access:token', res.token)
                    localStorage.setItem('user:info', JSON.stringify(res.user))
                    if (this.form.remember) localStorage.setItem('user:last:login', this.form.name)
                    else localStorage.removeItem('user:last:login')
                    setTimeout(() => {
                        this.$router.push('/home')
                    }, 2000)
                } else {
                    setTimeout(() => {
                        window.$message.warning(res.message ? res.message : '发生意料之外的错误')
                        this.form.loading = false
                    }, 1000)
                }
            }).catch(() => {
                window.$message.error('网络异常')
                this.form.loading = false
            })
        }
    },
    mounted() {
        this.init()
    }
};
</script>
  
<style scoped>
#auth-box {
    justify-content: center;
    align-items: center;
    position: relative;
    display: flex;
    height: 100vh;
}

#app-name {
    margin-bottom: 1.25rem;
    line-height: 2rem;
    font-size: 1.5rem;
    font-weight: 700;
}

#auth-card {
    background-color: #303c4e;
    margin-bottom: 100px;
    border-radius: 12px;
    text-align: left;
    width: 310px;
}

#auth-card form {
    padding: 20px 20px 0 20px;
}

#auth-tools {
    margin: 0 20px 25px 20px;
}

#auth-btn {
    width: 150px;
}

#auth-tips {
    border-top: #63676f 1px solid;
    line-height: 1.25rem;
    font-size: 0.875rem;
    padding: 20px 0;
}

#auth-foot {
    line-height: 1.25rem;
    font-size: 0.875rem;
    position: absolute;
    bottom: 10px;
    width: 100vw;
}
</style>
  
