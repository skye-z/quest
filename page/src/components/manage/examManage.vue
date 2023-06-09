<template>
    <div>
        <div class="card mb-10">
                <n-button class="float-right mt-5 mr-5" size="small" strong secondary circle>
                    <template #icon>
                        <n-icon>
                            <SearchRound />
                        </n-icon>
                    </template>
                </n-button>
            <div class="card-name">考试</div>
            <div class="exam-list"></div>
        </div>
    </div>
</template>
<script>
import { DeleteForeverRound, CheckRound,SearchRound } from '@vicons/material'
import { AddCircle24Regular } from '@vicons/fluent'
import { exam } from '../../plugins/api'
export default {
    components: { DeleteForeverRound, CheckRound,SearchRound, AddCircle24Regular },
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
.exam-list {
    height: 300px;
}
</style>