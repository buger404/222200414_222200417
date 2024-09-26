<script setup lang="ts">
import Button from "@/components/controls/Button.vue";
import {useRouter} from "vue-router";
const router = useRouter();

function navigate(path : string) {
  router.push(path);
}
</script>

<template>
  <main class="container">
    <p class="title">2024 巴黎奥运会</p>
    <p class="subtitle">数据浏览站</p>
    <div class="button_area">
      <Button
          @click="navigate('/overview')"
          content="查看各国奖牌榜"
          style="margin-right: 100px;
          width: 272px"
      />
      <Button content="关于我们" style="width: 272px"/>
    </div>
  </main>
</template>

<style scoped>
.container {
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
  url('../assets/home.jpg') lightgray 50% / cover no-repeat;
  width: 100%;
  height: 100%;
  overflow: auto;
}

.button_area{
  width: 100%;
  height: 100px;
  align-items: center;
  justify-content: center;
  display: flex;
  margin-top: 12vh;
}

.title {
  display: flex;
  width: 100%;
  height: 100px;
  flex-direction: column;
  justify-content: center;
  flex-shrink: 0;
  text-align: center;
  font-size: 64px;
  font-style: normal;
  font-weight: 800;
  line-height: 20px; /* 31.25% */
  letter-spacing: 0.1px;
  background: linear-gradient(273deg, #000 -26.29%, #F2BA71 84.49%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-top: 36vh;
}
.subtitle {
  display: flex;
  width: 100%;
  height: 50px;
  flex-direction: column;
  justify-content: center;
  flex-shrink: 0;
  text-align: center;
  font-size: 36px;
  font-style: normal;
  font-weight: 800;
  line-height: 20px; /* 55.556% */
  letter-spacing: 0.1px;
  background: linear-gradient(273deg, rgba(0, 0, 0, 0.50) -26.29%, rgba(242, 186, 113, 0.50) 84.49%);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>