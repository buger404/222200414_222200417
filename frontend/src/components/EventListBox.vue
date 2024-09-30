<script setup lang="ts">
import {ref} from "vue";
import Button from "@/components/controls/Button.vue";
import axios from "@/utils/axios";
import {ElMessage} from "element-plus";
import type {AllEventList, EventGroup} from "@/models/allEventList";
import {ArrowDown} from "@element-plus/icons-vue";
import EventBar from "@/components/controls/EventBar.vue";
import SimpleEventBar from "@/components/controls/SimpleEventBar.vue";

const active = ref(false);
const listLoaded = ref(false);
const eventID = ref('');

const title = ref('');

const groups = ref<EventGroup[]>();
const currentGroup = ref<EventGroup>();

const open = (t: string, event : string) => {
  title.value = t;
  eventID.value = event;
  active.value = true;
  fetchEventList();
}

const close = () => {
  active.value = false;
}

function fetchEventList() {
  listLoaded.value = false;
  axios.get('/api/event/list', {
    params: {
      eventID: eventID.value.slice(0, 22)
    }
  }).then((response) => {
    listLoaded.value = true;
    groups.value = response.data.data.event;
    if (groups.value){
      currentGroup.value = groups.value.find((g) => g.contests.findIndex((c) => c.contest_id == eventID.value) != -1);
      console.log(currentGroup.value);
    }
    console.log(response);
  }).catch((error) => {
    listLoaded.value = true;
    ElMessage({
      message: '网络错误，请重试',
      type: 'error',
      plain: true,
    });
    console.error('请求失败', error);
  })
}

function updateSelection(group : EventGroup) {
  currentGroup.value = group;
}

defineExpose({ open });
</script>

<template>
  <transition name="slide">
    <div class="event_list_box" v-if="active">
      <div class="panel" v-loading="!listLoaded">
        <div class="row" style="width: 100%; align-items: end;">
          <h1>{{ title }} 详细赛程</h1>
          <div class="row" v-if="groups && currentGroup">
            <h3 style="margin-right: 10px; font-size: 18px;">项目组别：</h3>
            <el-dropdown
                class="drop_down"
                max-height="60vh"
            >
              <el-button type="primary">
                {{ currentGroup.group_id }}<el-icon class="el-icon--right"><arrow-down /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu class="drop_down_menu">
                  <el-dropdown-item
                      v-for="item in groups"
                      @click="updateSelection(item)"
                  >
                    {{ item.group_id }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
        <div style="width: 100%; overflow-x: scroll; overflow-y: hidden;">
          <div class="event_bar_box" v-if="currentGroup">
            <SimpleEventBar
                v-for="item in currentGroup.contests"
                :active="(item.contest_id == eventID)"
                :id="(item.contest_id == eventID) ? 'currentEvent' : ''"
                :data="item"/>
          </div>
        </div>
        <div style="height: 100px; width: 100%; display: flex; align-items: center;">
          <Button @click="close()" content="返回每日赛程" style="width: 100%;"/>
        </div>
      </div>
    </div>
  </transition>
</template>

<style scoped>
.event_bar_box {
  width: auto;
  display: flex;
  flex-direction: row;
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
  font-size: 36px;
  font-style: normal;
  font-weight: 700;
  letter-spacing: 0.1px;
}

.event_list_box {
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
  z-index: 10;
}

.panel {
  --el-loading-spinner-size: 90px;
  --el-color-primary: #f33e3e;
  width: 80vw;
  min-height: 400px;
  height: 400px;
  border-radius: 13px;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.90) 0%, #FFF 100%),
  url('../assets/table.jpg') lightgray 50% / cover no-repeat;
  box-shadow: 0px 4px 16px 0px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  vertical-align: center;
  padding: 20px 45px;

  color: #2D2D2D;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  letter-spacing: 0.1px;
}
</style>