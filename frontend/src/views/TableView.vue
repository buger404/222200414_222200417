<script setup lang="ts">
import {ArrowDown} from "@element-plus/icons-vue";
import type {EventType, EventTypeList} from "@/models/eventTypeList";
import {ref} from "vue";
import axios from "@/utils/axios";
import {ElMessage} from "element-plus";

const typeListLoaded = ref(false);
const types = ref<EventType[]>();
const selectedType = ref<EventType>();

function fetchEventTypes(): void {
  typeListLoaded.value = false;
  types.value = [];
  axios.get('/api/event/typelist')
  .then((response : EventTypeList) => {
    types.value = response.data.types;
    selectedType.value = types.value[0];
    typeListLoaded.value = true;
    console.log(response);
  }).catch((error) => {
    console.error('请求失败', error);
  })
}

fetchEventTypes();

</script>

<template>
  <main class="container">
    <div class="sub_container" v-loading="!typeListLoaded">
      <div class="row" style="position: sticky; top: 0; z-index: 1;">
        <h1>对阵表</h1>
        <el-dropdown
            class="drop_down"
            v-if="typeListLoaded && selectedType"
            max-height="60vh"
        >
          <el-button type="primary">
            比赛项目：{{ selectedType.name }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu class="drop_down_menu">
              <el-dropdown-item
                  v-for="item in types"
                  @click="selectedType = item;"
              >
                {{ item.name }}
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
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
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.80) 0%, #FFF 100%),
  url('../assets/table.jpg') white 100% / cover no-repeat;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.sub_container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 100%;
  height: 100%;
  overflow: auto;
  margin-top: 72px;
  padding: 58px 7vw 120px;

  --el-mask-color: transparent;
  --el-loading-spinner-size: 90px;
  --el-color-primary: #f33e3e;
}
</style>