<template>
    <div class="loading flex align-center justify-center" :class="{ hide: !isShow }" v-if="loading">
        <n-spin :size="30" :stroke-width="30" stroke="#6d83a1" />
    </div>
</template>
<script>
export default {
    data: () => ({
        loading: true,
        isShow: true
    }),
    methods: {
        show() {
            this.loading = true;
            this.isShow = true;
        },
        hide() {
            this.isShow = false;
            setTimeout(() => {
                this.loading = false;
            }, 350);
        }
    }
};
</script>
<style scoped>
.loading {
    background-color: #1f2937;
    transition: all .3s ease-in;
    position: fixed;
    z-index: 9999;
    height: 100vh;
    width: 100vw;
    left: 0;
    top: 0;
}

.loading.hide {
    opacity: 0;
}
</style>