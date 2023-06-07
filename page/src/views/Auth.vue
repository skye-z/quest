<template>
    <div id="auth-box" class="no-select">
        <div class="text-center">
            <div id="app-name">{{ app.name }}</div>
            <div id="auth-card">
                <n-form ref="form" size="large" :show-require-mark="false" :model="form" :rules="rules">
                    <n-form-item label="用户名" path="name">
                        <n-input v-model:value="form.name" placeholder="请输入用户名" type="text" />
                    </n-form-item>
                    <n-form-item label="密码" path="pass">
                        <n-input v-model:value="form.pass" placeholder="请输入密码" type="password" />
                    </n-form-item>
                </n-form>
                <div id="auth-tools" class="flex justify-between align-center">
                    <div class="flex align-center text-gray">
                        <n-switch class="mr-5" /> 记住我
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
            this.loading = false;
        },
        submit() {
            this.form.loading = true
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

#auth-tools{
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
  