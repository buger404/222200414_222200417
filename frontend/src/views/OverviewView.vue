<script setup lang="ts">
import Button from "@/components/controls/Button.vue";
import { ref } from "vue";
import {ArrowDown, Search} from '@element-plus/icons-vue'
import CountryRankBar from "@/components/controls/CountryRankBar.vue";

const searchCountry = ref('');
</script>

<template>
  <main class="container">
    <div class="row">
      <h1 style="width: calc(100% - 240px); ">各国奖牌数总榜</h1>
      <el-input
          class="input_box"
          v-model="searchCountry"
          style="width: 240px;"
          size="large"
          placeholder="搜索国家..."
          :prefix-icon="Search"
      />
    </div>
    <div class="row">
      <el-dropdown class="drop_down" style="width: 35%">
        <el-button type="primary">
          排序方式<el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu class="drop_down_menu">
            <el-dropdown-item>金牌</el-dropdown-item>
            <el-dropdown-item>总奖牌数</el-dropdown-item>
            <el-dropdown-item>首字母</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <div style="width: 15%; text-align: center;">
        <img style="width: 40px; height: 40px;" src="../assets/gold.png" alt="金牌" />
      </div>
      <div style="width: 15%; text-align: center;">
        <img style="width: 40px; height: 40px;" src="../assets/silver.png" alt="银牌" />
      </div>
      <div style="width: 15%; text-align: center;">
        <img style="width: 40px; height: 40px;" src="../assets/bronze.png" alt="铜牌" />
      </div>
      <div style="width: 15%; text-align: center;">
        <h2>合计</h2>
      </div>
      <div style="width: 5%;"/>
    </div>
    <div style="height: 100%; overflow: auto; scrollbar-width: none;">
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
      <CountryRankBar
          name="美利坚合众国"
          flag="USA"
          gold="40"
          silver="44"
          bronze="42"
          total="126"
          rank=1
      />
    </div>
  </main>
</template>

<style scoped>
.drop_down {
  --el-font-size-base: 18px;
  --el-border-radius-base: 180px;
}

.drop_down_menu {
  --el-dropdown-menuItem-hover-fill: #f33e3e;
  --el-dropdown-menuItem-hover-color: white;
  --el-menu-active-color: #f33e3e;
  --el-menu-hover-text-active-color: #f33e3e;
  --el-menu-hover-text-color: #f33e3e;
  --el-font-size-base: 18px;
}

.drop_down button {
  --el-button-bg-color: #F18787;
  --el-button-hover-bg-color: #f33e3e;
  --el-button-active-bg-color: #f33e3e;
  --el-button-outline-color: transparent;
  --el-button-border-color: #eaaaaa;
  --el-button-hover-border-color: #f33e3e;
  --el-button-active-border-color: #f33e3e;
  padding: 21px 21px 21px 25px;
}

.drop_down button span i {
  margin-left: 15px;
}

.input_box {
  --el-input-focus-border-color: #F18787;
  --el-input-border-radius: 180px;
  --el-input-height: 45px;
  height: 45px;
  font-size: 18px;
}

.row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 20px;
  align-items: center;
}

h1 {
  color: black;
  font-size: clamp(24px, 48px, 48px);
  font-style: normal;
  font-weight: 700;
  letter-spacing: 0.1px;
}

.container {
  display: flex;
  flex-direction: column;
  padding: 130px 7vw 0;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
  url('../assets/overview.jpg') white 100% / cover no-repeat;
  width: 100%;
  height: 100%;
  overflow: auto;
}
</style>
