<template>
    <div id="app-head" class="flex align-center justify-between no-select">
        <div class="flex align-center">
            <n-icon size="40" class="mr-5">
                <school-round />
            </n-icon>
            <div id="app-name">{{ app.name }}</div>
        </div>
        <div class="flex align-center">
            <div id="user-name" class="mr-5">{{ user.nickname }}</div>
            <n-icon size="24">
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
    }
};
</script>
  
<style scoped>
#app-head {
    background-color: #1f2937;
    border-bottom: #3c4a5f 1px solid;
    padding: 10px;
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
</style>
  