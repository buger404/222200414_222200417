<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import NavigatorItem from "@/components/controls/NavigatorItem.vue";
import HomeIcon from "@/components/icons/IconHome.vue";
</script>

<template>
  <header>
    <svg class="logo"/>
    <NavigatorItem content="首页" to="/">
      <HomeIcon/>
    </NavigatorItem>
    <NavigatorItem content="关于" to="/about">
      <HomeIcon/>
    </NavigatorItem>
    <NavigatorItem content="首页" to="/">
      <HomeIcon/>
    </NavigatorItem>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  display: flex;
  flex-direction: row;
  position: fixed;
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
