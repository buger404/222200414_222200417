<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { computed } from 'vue';

const props = defineProps<{
  content: string,
  to: string
}>();

const route = useRoute();
const router = useRouter();

const isActive = computed(() => {
  return route.path === props.to;
});

function navigate() {
  router.push(props.to);
}
</script>

<template>
  <div
      @click="navigate"
      :style="{ 'border-bottom': isActive ? '#F18787 4px solid' : 'transparent 4px solid' }"
      class="navigator_container"
  >
    <div style="width: 30px; height: 30px;">
      <slot></slot>
    </div>
    <span class="link">{{ content }}</span>
  </div>
</template>

<style scoped>
.link{
  margin-left: 0.8rem;
  color: black;
  text-align: center;
  font-size: 18px;
  font-style: normal;
  font-weight: 500;
  line-height: 20px; /* 100% */
  letter-spacing: 0.1px;
}
.navigator_container {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 100%;
  margin-left: 2.5rem;
  cursor: pointer;
}
</style>