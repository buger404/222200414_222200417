<script setup lang="ts">
import {ref} from "vue";
import axios from "@/utils/axios";
import {ElMessage} from "element-plus";
import type {DayEvent, DayEventList} from "@/models/dayEventList";
import EventBar from "@/components/controls/EventBar.vue";
import EventListBox from "@/components/EventListBox.vue";

const date = ref(new Date("2024-07-24"))

const startDate = new Date("2024-07-23");
const endDate = new Date("2024-08-12");

function isDateInvalid(date: Date): boolean {
  return !(date.getTime() >= startDate.getTime() && date.getTime() <= endDate.getTime());
}

let eventList : DayEventList;

const listLoaded = ref(false);
const displayList = ref<DayEvent[]>([]);

function fetchEventList() {
  listLoaded.value = false;
  displayList.value = [];
  axios.get('/api/event/daily', {
    params: {
      date: date.value.toLocaleDateString('en-US', { month: '2-digit', day: '2-digit' }).replace(',', '').replace('/', '-')
    }
  }).then((response) => {
    eventList = response.data;
    listLoaded.value = true;
    console.log(eventList);
    refreshDisplayList();
    ElMessage({
      message: '已成功获取每日赛程信息',
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
  displayList.value = eventList.data.events;
}

fetchEventList();

const eventListBox = ref();

function showEventListBox(title : string, event : string) {
  eventListBox.value.open(title, event);
}
</script>

<template>
  <main class="container">
    <div class="sub_container">
      <div class="row">
        <h1>每日赛程</h1>
        <div class="date_picker">
          <el-date-picker
              @change="fetchEventList()"
              v-model="date"
              type="date"
              placeholder="选择日期..."
              format="MM-DD"
              :disabled-date="isDateInvalid"
          />
        </div>
      </div>
      <div>
        <h2 style="margin-top: -15px; margin-bottom: 15px;">注：赛事时间为巴黎时间(UTC+2)。</h2>
      </div>
      <div class="event_box" v-loading="!listLoaded">
        <h2 v-if="listLoaded && !displayList">╰(￣ω￣ｏ) 该日无赛事类型的活动哦~</h2>
        <EventBar
            @click="showEventListBox(event.group, event.id)"
            v-for="(event) in displayList"
            :data="event"
        />
      </div>
    </div>
    <EventListBox ref="eventListBox"/>
  </main>
</template>

<style scoped>
.event_box {
  padding-bottom: 60px;
  flex-grow: 1;
  --el-mask-color: transparent;
  --el-loading-spinner-size: 90px;
  --el-color-primary: #f33e3e;
}

.date_picker {
  --el-border-radius-base: 180px;
  --el-color-primary: #F18787;
  scale: 1.3;
  margin-right: 30px;
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
  url('../assets/universal.jpg') white 100% / cover no-repeat;
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
  padding: 58px 7vw 120px;
}
</style>