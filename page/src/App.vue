<template>
  <n-config-provider :theme-overrides="{ common: { fontWeightStrong: '600' } }">
    <router-view v-slot="{ Component }">
      <transition name="fade">
        <component :is="Component" />
      </transition>
    </router-view>
  </n-config-provider>
</template>

<script>
import { init } from './plugins/api'

export default {
  name: "App",
  data: () => ({
  }),
  methods: {
    checkLogin() {
      let token = localStorage.getItem('user:access:token')
      if (token == '' || token == undefined) this.$router.push('/auth')
    },
    initConfig() {
      init().then(res => {
        localStorage.setItem('app:config:name',res.name)
        localStorage.setItem('app:config:version',res.version)
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
