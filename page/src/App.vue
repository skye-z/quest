<template>
  <n-config-provider :theme="dark" :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <n-dialog-provider>
      <n-message-provider>
        <global-api />
      </n-message-provider>
    </n-dialog-provider>
    <router-view />
    <!-- <NThemeEditor /> -->
  </n-config-provider>
</template>

<script>
import { init } from './plugins/api'
import { darkTheme, zhCN, dateZhCN } from 'naive-ui'
import theme from './theme.json'
import GlobalApi from './components/globalApi.vue'

export default {
  name: "App",
  components: { GlobalApi },
  data: () => ({
    dark: darkTheme,
    theme,
    i18n: {
      main: zhCN,
      date: dateZhCN
    }
  }),
  methods: {
    checkLogin() {
      let path = window.location.pathname;
      let token = localStorage.getItem('user:access:token')
      if (token == '' || token == undefined) this.$router.push('/auth')
      else if(path == '/' || path == '/app' || path == '/app/') this.$router.push('/home')
    },
    initConfig() {
      init().then(res => {
        localStorage.setItem('app:config:name', res.name)
        localStorage.setItem('app:config:version', res.version)
        this.checkLogin()
      })
    }
  },
  mounted() {
    this.initConfig()
  },
  watch: {
    $route: {
      handler(to) {
        if (to.path == '/') this.checkLogin()
      },
      deep: true,
    },
  },
};
</script>

<style scoped></style>
