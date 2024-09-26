<script lang="ts" setup>
defineProps<{
  content: string
}>()
</script>

<template>
  <span class="button">{{ content }}</span>
</template>

<style scoped>
.button {
  border-radius: 109px;
  border: 2px solid #B7B7B7;
  background: rgba(255, 255, 255, 0.50);
  box-shadow: 0px 4px 16px 0px rgba(0, 0, 0, 0.15);
  text-align: center;
  padding: 11px 30px 15px;
  transition: all .5s;
  color: black;
  font-size: 20px;
  font-style: normal;
  font-weight: 400;
  line-height: 20px; /* 83.333% */
  letter-spacing: 0.1px;
}
.button:hover {
  cursor: pointer;
  border-radius: 109px;
  border: 2px solid #000;
  background: #F18787;
  box-shadow: 0px 4px 16px 0px rgba(241, 135, 135, 0.50);
  text-align: center;
  align-items: center;
  color: black;
}
</style>