<script setup lang="ts">
import Button from "@/components/controls/Button.vue";
import { ref, onMounted } from "vue";
import {ArrowDown, Search} from '@element-plus/icons-vue'
import CountryRankBar from "@/components/controls/CountryRankBar.vue";
import type { MedalList, MedalListData } from "@/models/modelMedalList";
import axios from '@/utils/axios';
import {ElMessage} from "element-plus";

let medalList : MedalList;

const listLoaded = ref(false);
const searchCountry = ref('');
const sortMethod = ref(0);
const displayList = ref<MedalListData[]>([]);

const sortMethods = ["金牌", "总奖牌数", "首字母"]

function fetchMedalList() {
  listLoaded.value = false;
  axios.get('/api/medals/all', {
    params: {
      medalSorts: sortMethod.value.toString()
    }
  }).then((response) => {
    medalList = response.data;
    listLoaded.value = true;
    console.log(medalList);
    refreshDisplayList();
    ElMessage({
      message: '已成功获取奖牌信息',
      type: 'success',
      plain: true,
    });
  }).catch((error) => {
    ElMessage({
      message: '网络错误，请重试',
      type: 'error',
      plain: true,
    });
    console.error('请求失败', error);
  })
}

function refreshDisplayList() {
  displayList.value = medalList.data.details;
}

fetchMedalList();

</script>

<template>
  <main class="container">
    <div class="row">
      <h1 style="width: calc(100% - 240px); ">各国奖牌数总榜</h1>
      <el-input
          class="input_box"
          @input="refreshDisplayList()"
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
          排序方式：{{sortMethods[sortMethod]}}<el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu class="drop_down_menu">
            <el-dropdown-item
                v-for="(item, index) in sortMethods"
                @click="sortMethod = index; fetchMedalList();"
            >
              {{item}}
            </el-dropdown-item>
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
    <div class="rank_box" v-loading="!listLoaded">
      <h2 v-if="listLoaded && displayList.every(item => !(searchCountry === '' || item.countryName.includes(searchCountry)))">没有匹配的国家。</h2>
      <CountryRankBar
          v-for="(item, index) in displayList"
          v-if="listLoaded && (displayList.length > 0)"
          :name="(sortMethod == 2 ? (item.flag + ':') : '') + item.countryName"
          :flag="item.flag"
          :gold="item.list.gold ?? 0"
          :silver="item.list.sliver ?? 0"
          :bronze="item.list.bronze ?? 0"
          :total="item.list.total ?? 0"
          :rank="index + 1"
          :style="(searchCountry === '' || item.countryName.includes(searchCountry)) ? 'display: block;' : 'display: none;'"
      />
    </div>
  </main>
</template>

<style scoped>
.rank_box {
  height: 100%;
  flex-grow: 1;
  overflow: auto;
  padding: 0 7vw 130px;
  --el-mask-color: transparent;
  --el-loading-spinner-size: 90px;
  --el-color-primary: #f33e3e;
}

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
  padding: 0 7vw;
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
  padding: 130px 0 0;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
  url('../assets/overview.jpg') white 100% / cover no-repeat;
  width: 100%;
  height: 100%;
  overflow: auto;
}
</style>
