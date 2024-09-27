<script setup lang="ts">
import {ref} from "vue";
import Button from "@/components/controls/Button.vue";
const props = defineProps<{
  title: string
}>();

const active = ref(false);

const open = () => {
  active.value = true;
}

const close = () => {
  active.value = false;
}

defineExpose({ open });
</script>

<template>
  <transition name="slide">
    <div v-if="active" class="message_box">
      <div class="panel">
        <div style="height: 20%; display: flex; align-items: center;">
          <h1>{{title}}</h1>
        </div>
        <div style="height: 60%; text-align: left; width: 100%; overflow: auto; scrollbar-width: none;">
          <slot></slot>
        </div>
        <div style="height: 20%; width: 100%; display: flex; align-items: center;">
          <Button @click="close()" content="确定" style="width: 100%;"/>
        </div>
      </div>
    </div>
  </transition>
</template>

<style scoped>
h1{
  color: black;
  text-align: center;
  font-size: 36px;
  font-style: normal;
  font-weight: 600;
  letter-spacing: 0.1px;
}

.message_box {
  transition: 0.25s all;
  background: rgba(255, 255, 255, 0.50);
  backdrop-filter: blur(25px);
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999999999;
}
.panel {
  width: clamp(60vw, 935px, 935px);
  height: clamp(40vw, 551px, 551px);
  border-radius: 26px;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
              url('../assets/universal.jpg') lightgray 50% / cover no-repeat;
  box-shadow: 0px 4px 16px 0px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  vertical-align: center;
  padding: 20px 50px;

  color: #2D2D2D;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  letter-spacing: 0.1px;
}
</style>