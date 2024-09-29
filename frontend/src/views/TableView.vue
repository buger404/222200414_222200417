<script setup lang="ts">
import {ArrowDown} from "@element-plus/icons-vue";
import type {EventType, EventTypeItem, EventTypeList} from "@/models/eventTypeList";
import {ref} from "vue";
import axios from "@/utils/axios";
import {ElMessage} from "element-plus";
import type {Table, TableList} from "@/models/tableList";
import VSTable from "@/components/controls/VSTable.vue";
import VSLine from "@/components/controls/VSLine.vue";

const typeListLoaded = ref(false);
const tableLoaded = ref(false);

const types = ref<EventType[]>();
const selectedType = ref<EventType>();
const selectedItem = ref<EventTypeItem>();

const tables = ref<Table[]>();
const periodTable = ref<Array<Table[]>>();

function fetchEventTypes(): void {
  typeListLoaded.value = false;
  types.value = [];
  axios.get('/api/event/typelist')
  .then((response : EventTypeList) => {
    types.value = response.data.list;
    console.log(response);
    selectedType.value = types.value?.find((item : EventType) : boolean => { return item.name == "乒乓球"; })
    selectedItem.value = selectedType.value?.types[0];
    fetchTableData();
    typeListLoaded.value = true;
  }).catch((error) => {
    ElMessage({
      message: '网络错误，请重试',
      type: 'error',
      plain: true,
    });
    console.error('请求失败', error);
  })
}

function fetchTableData(): void {
  tableLoaded.value = false;
  tables.value = [];
  periodTable.value = [[], [], []];
  axios.get('/api/event/table',{
    params: {
      eventID : selectedItem.value?.id
    }
  }).then((response : TableList) => {
    console.log(response);
    tables.value = response.data.tables;
    if (tables.value){
      periodTable.value[0] = tables.value.filter((table) => table.title.includes('1/4决赛'));
      periodTable.value[1] = tables.value.filter((table) => table.title.includes('半决赛'));
      periodTable.value[2] = tables.value.filter(
          (table) => (table.title.includes('金牌赛') || table.title.includes('银牌赛') ||
                     table.title.includes('铜牌赛') || table.title.includes('决赛'))
                      && !periodTable.value[0].includes(table)
                      && !periodTable.value[1].includes(table))
          .sort((a, b) => {
            const order = ['金牌赛', '银牌赛', '铜牌赛'];
            const indexA = order.findIndex(keyword => a.title.includes(keyword));
            const indexB = order.findIndex(keyword => b.title.includes(keyword));
            return indexA - indexB;
          });
    }
    tableLoaded.value = true;
  }).catch((error) => {
    ElMessage({
      message: '网络错误，请重试',
      type: 'error',
      plain: true,
    });
    console.error('请求失败 或 对阵表不存在', error);
  })
}

function updateSelection(type : EventType) {
  selectedType.value = type;
  selectedItem.value = selectedType.value?.types[0];
  fetchTableData();
}

fetchEventTypes();

</script>

<template>
  <main class="container">
    <div class="sub_container" v-loading="!typeListLoaded || !tableLoaded">
      <div class="head_bar">
        <div class="row">
          <h1>对阵表</h1>
          <div class="row"
               v-if="typeListLoaded && selectedType && tableLoaded"
          >
            <h3 style="margin-right: 10px;">比赛项目：</h3>
            <el-dropdown
                class="drop_down"
                max-height="60vh"
            >
              <el-button type="primary">
                {{ selectedType.name }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu class="drop_down_menu">
                  <el-dropdown-item
                      v-for="item in types"
                      @click="updateSelection(item)"
                  >
                    {{ item.name }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>

            <el-dropdown
                class="drop_down"
                style="margin-left: 20px;"
                max-height="60vh"
            >
              <el-button type="primary">
                {{ selectedItem.name }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu class="drop_down_menu">
                  <el-dropdown-item
                      v-for="item in selectedType.types"
                      @click="selectedItem = item; fetchTableData();"
                  >
                    {{ item.name }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </div>

      <div style="padding: 0 7vw; flex-grow: 1;" v-if="tableLoaded && !tables">
        <h2>当前比赛不是晋级赛，不支持查看对阵表。</h2>
      </div>

      <div style="width: 100%; flex-grow: 1; overflow: auto;">
        <div class="table_box" v-if="tableLoaded && tables">
          <div class="period_header">
            <div
                class="title"
                style="margin-right: 80px; border: 2px solid #D9D9D9; background: white;"
                v-if="periodTable[0].length > 0"
            >
              <h2 style="width: 100%;">1/4 决赛</h2>
            </div>
            <div
                class="title"
                style="margin-right: 80px; border: 2px solid black; background: #F2BA71;"
                v-if="periodTable[1].length > 0"
            >
              <h2 style="width: 100%;">半决赛</h2>
            </div>
            <div
                class="title"
                style="border: 2px solid black; background: #F18787;"
            >
              <h2 style="width: 100%;">决赛</h2>
            </div>
          </div>

          <div style="display: flex; flex-direction: row; padding: 0 7vw;">
            <div class="table_item" v-if="periodTable[0].length > 0">
              <VSTable
                  v-for="item in periodTable[0]"
                  :data="item"
                  style="margin: calc(30px) 0"
              />
            </div>

            <div style="width: 80px" v-if="periodTable[0].length > 0">
              <VSLine style="width: 100%; height: calc(30px * 2 + 143px); margin: calc(30px + 143px / 2) 0;"/>
              <VSLine style="width: 100%; height: calc(30px * 2 + 143px); margin-top: calc((30px + 143px / 2) * 2);"/>
            </div>

            <div class="table_item" v-if="periodTable[1].length > 0">
              <VSTable
                  v-for="item in periodTable[1]"
                  :data="item"
                  style="margin: calc(30px * 2 + 143px / 2) 0"
              />
            </div>

            <div style="width: 80px" v-if="periodTable[1].length > 0">
              <VSLine style="width: 100%; height: calc((30px * 2 + 143px / 2) * 2 + 143px);
              margin: calc((30px * 2 + 143px / 2) + 143px / 2) 0;"/>
            </div>

            <div class="table_item">
              <VSTable
                  v-for="(item, index) in periodTable[2]"
                  :data="item"
                  :style="index > 0 ? 'margin-top: 80px' : 'margin-top: calc(30px * 4 + 143px * 1.5)'"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.period_header {
  display: flex;
  flex-direction: row;
  margin-bottom: 20px;
  padding: 0 7vw;
  position: sticky;
  top: 0;
  z-index: 1;
}

.head_bar {
  position: sticky;
  top: 0px;
  z-index: 1;
  padding: 0 7vw;
}

.table_item {
  display: flex;
  flex-direction: column;
  width: clamp(308.65px, calc(28.66667vw - 53px), calc(28.66667vw - 53px));
}

.title {
  display: flex;
  flex-direction: row;
  border-radius: 76px;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.15);
  width: clamp(308.65px, calc(28.66667vw - 53px), calc(28.66667vw - 53px));
  height: 45px;
  text-align: center;
  overflow: visible;
}

.table_box {
  width: clamp(1085.95px + 14vw, 100%, 100%);
  display: flex;
  flex-direction: column;
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
  width: 100%;
  height: 100%;
  overflow: auto;
  margin-top: 72px;
  padding: 58px 0 72px;

  --el-mask-color: transparent;
  --el-loading-spinner-size: 90px;
  --el-color-primary: #f33e3e;
}
</style>