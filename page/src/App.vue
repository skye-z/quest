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
export default {
  name: "App",
  data: () => ({

  }),
  methods: {

  }
};
</script>

<style scoped>
</style>
