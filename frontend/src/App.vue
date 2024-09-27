<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import NavigatorItem from "@/components/controls/NavigatorItem.vue";
import HomeIcon from "@/components/icons/IconHome.vue";
import OverviewIcon from "@/components/icons/IconOverview.vue";
import IconDay from "@/components/icons/IconDay.vue";
</script>

<template>
  <header>
    <svg class="logo"/>
    <NavigatorItem content="首页" to="/">
      <HomeIcon/>
    </NavigatorItem>
    <NavigatorItem content="概览" to="/overview">
      <OverviewIcon/>
    </NavigatorItem>
    <NavigatorItem content="每日赛程" to="/day">
      <IconDay/>
    </NavigatorItem>
  </header>

  <RouterView v-slot="{ Component }">
    <transition name="slide" mode="out-in">
      <component :is="Component"/>
    </transition>
  </RouterView>
</template>

<style scoped>
header {
  display: flex;
  flex-direction: row;
  position: fixed;
  z-index: 10;
  top: 0;
  background: rgba(255, 255, 255, 0.80);
  box-shadow: 0px 4px 16px 0px rgba(0, 0, 0, 0.10);
  backdrop-filter: blur(8.75px);
  width: 100%;
  height: 72px;
}

.logo {
  border: none;
  margin-left: 20px;
  margin-top: 14px;
  background: url('assets/logo.png') transparent 100% / cover no-repeat;
  width: 40px;
  height: 40px;
}
</style>
