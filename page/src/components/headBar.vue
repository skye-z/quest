<template>
    <div id="app-head" class="flex align-center justify-between no-select">
        <div class="flex align-center app-title" @click="goHome">
            <n-icon size="40" class="mr-5">
                <school-round />
            </n-icon>
            <div id="app-name">{{ app.name }}</div>
        </div>
        <div class="flex align-center">
            <div id="user-name" class="mr-5">{{ user.nickname }}</div>
            <n-icon class="exit" @click="exit" size="24">
                <power-settings-new-round />
            </n-icon>
        </div>
    </div>
</template>
<script>
import { SchoolRound, PowerSettingsNewRound } from '@vicons/material'

export default {
    components: { SchoolRound, PowerSettingsNewRound },
    props: {
        subjects: {
            type: Array,
            default: []
        },
        app: {
            type: Object,
            default: {
                name: 'Quest云题库',
                version: '1.0.0'
            }
        },
        user: {
            type: Object,
            default: {
                id: 0,
                nickname: '',
                name: '',
                admin: false,
                edit: false
            }
        }
    },
    methods: {
        goHome() {
            this.$router.push('/home')
        },
        exit() {
            window.$dialog.warning({
                title: '操作确认',
                content: '您正在尝试退出登录, 确认要继续吗?',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    localStorage.removeItem('user:info');
                    localStorage.removeItem('user:access:token');
                    window.$message.success('登录已注销, Bye!');
                    this.$router.push('/auth')
                }
            })
        }
    }
};
</script>
  
<style scoped>
#app-head {
    background-color: #1f2937;
    border-bottom: #3c4a5f 1px solid;
    padding: 10px;
}

.app-title{
    cursor: pointer;
}

#app-name {
    line-height: 20px;
    font-weight: 600;
    font-size: 20px;
}

#user-name {
    line-height: 16px;
    font-size: 16px;
}

.exit {
    transition: all .3s ease-in;
    cursor: pointer;
    color: #aaaaaa;
}

.exit:hover {
    color: #fff;
}

.exit:active {
    color: #ffbcbc;
}
</style>
  