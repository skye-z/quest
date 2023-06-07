<template>
    <div id="app-foot" class="text-center full-width text-gray">Powered by <a href="https://github.com/skye-z/quest"
            target="_blank">Quest</a> v{{ app.version }}</div>
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
#app-foot{
    padding: 10px 0;
}
</style>