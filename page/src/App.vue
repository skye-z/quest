<template>
  <n-config-provider :theme="dark" :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <n-watermark v-if="mask" :content="mask" cross fullscreen
    :font-size="16" :line-height="16" :width="300" :height="250" :x-offset="5" :y-offset="60" :rotate="-15"/>
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
    },
    mask: ''
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
        this.mask = res.mask
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
