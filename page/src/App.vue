<template>
  <n-config-provider :theme="dark" :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <router-view v-slot="{ Component }">
      <transition name="fade">
        <component :is="Component" />
      </transition>
    </router-view>
    <NThemeEditor />
  </n-config-provider>
</template>

<script>
import { init } from './plugins/api'
import { darkTheme, zhCN, dateZhCN, NThemeEditor } from 'naive-ui'
import theme from './theme.json'

export default {
  name: "App",
  components: { NThemeEditor },
  data: () => ({
    dark: darkTheme,
    theme,
    i18n:{
      main: zhCN,
      date: dateZhCN
    }
  }),
  methods: {
    checkLogin() {
      let token = localStorage.getItem('user:access:token')
      if (token == '' || token == undefined) this.$router.push('/auth')
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
