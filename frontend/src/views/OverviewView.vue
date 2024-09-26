<script setup lang="ts">
import Button from "@/components/controls/Button.vue";
import { ref } from "vue";
import { Search } from '@element-plus/icons-vue'

const searchCountry = ref('');
</script>

<template>
  <main class="container">
    <div class="row">
      <h1>各国奖牌数总榜</h1>
      <el-input
          class="input_box"
          v-model="searchCountry"
          style="width: 240px;"
          size="large"
          placeholder="搜索国家..."
          :prefix-icon="Search"
      />
    </div>
  </main>
</template>

<style scoped>
.input_box {
  --el-input-focus-border-color: #F18787;
  --el-input-border-radius: 180px;
  --el-input-height: 45px;
  font-size: 18px;
}

.row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

h1 {
  color: black;
  font-size: 48px;
  font-style: normal;
  font-weight: 700;
  line-height: 20px; /* 41.667% */
  letter-spacing: 0.1px;
}

.container {
  display: flex;
  flex-direction: column;
  padding: 150px 7vw;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
  url('../assets/overview.jpg') white 100% / cover no-repeat;
  width: 100%;
  height: 100%;
  overflow: auto;
}
</style>
