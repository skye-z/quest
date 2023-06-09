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
                <n-button size="small" type="primary" round>
                    <template #icon>
                        <n-icon>
                            <AddCircle24Regular />
                        </n-icon>
                    </template>
                    添加
                </n-button>
            </div>
            <div class="card-name">题目</div>
            <div class="question-list"></div>
        </div>
    </div>
</template>
<script>
import { DeleteForeverRound, CheckRound, SearchRound } from '@vicons/material'
import { AddCircle24Regular } from '@vicons/fluent'
import { question } from '../plugins/api'
export default {
    components: { DeleteForeverRound, CheckRound, SearchRound, AddCircle24Regular },
    data: () => ({
        loading: true,
        list: [],
        form: {
            show: false,
            loading: false,
        },
        rules: {}
    }),
    methods: {
        getList() {
            this.loading = true
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
</style>