<template>
    <div id="app-foot" class="text-center full-width text-gray">
        <span>Powered by </span>
        <a href="https://github.com/skye-z/quest" rel="noopener" target="_blank">
            <n-gradient-text type="info">
                Quest
            </n-gradient-text>
        </a>
        <span> v{{ app.version }}</span>
    </div>
</template>
<script>
export default {
    props: {
        app: {
            type: Object,
            default: {
                name: 'Quest云题库',
                version: '1.0.0'
            }
        }
    }
};
</script>
<style scoped>
#app-foot {
    padding: 10px 0;
}
</style>