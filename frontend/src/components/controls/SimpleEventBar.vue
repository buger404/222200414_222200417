<script setup lang="ts">
import type {DayEvent} from "@/models/dayEventList";
import LargeCountryFlag from "@/components/controls/LargeCountryFlag.vue";
import IconVS from "@/components/icons/IconVS.vue";
import type {AllEventList, Contest} from "@/models/allEventList";
import CountryFlag from "@/components/controls/CountryFlag.vue";

const props = defineProps<{
  data : Contest
  active : boolean
}>();

function formatDate(dateStr : string) {
  const date = new Date(dateStr);
  const month = String(date.getMonth() + 1).padStart(2, '0'); // 获取月份，月份从0开始，所以需要+1
  const day = String(date.getDate()).padStart(2, '0'); // 获取日
  const hours = String(date.getHours()).padStart(2, '0'); // 获取小时
  const minutes = String(date.getMinutes()).padStart(2, '0'); // 获取分钟

  return `${month}月${day}日 ${hours}:${minutes}`;
}
</script>

<template>
  <div :class="active ? 'box2' : 'box'">
    <a class="time_text">{{ formatDate(data.date) }}</a>
    <div style="width: 100%; display: flex; flex-direction: row; justify-content: space-between; margin-top: 0.5rem;">
      <div style="display: flex; flex-direction: row;">
        <CountryFlag :country="data.countries[0].flag" style="margin-right: 0.5rem;"/>
        <a :class="(data.status == 0) ? 'a_win' : ''">{{ data.countries[0].name }}</a>
      </div>
      <a :class="(data.status == 0) ? 'a_win' : ''">{{ data.countries[0].rating }}</a>
    </div>
    <div style="width: 100%; display: flex; flex-direction: row; justify-content: space-between;">
      <div style="display: flex; flex-direction: row;">
        <CountryFlag :country="data.countries[1].flag" style="margin-right: 0.5rem;"/>
        <a :class="(data.status == 1) ? 'a_win' : ''">{{ data.countries[1].name }}</a>
      </div>
      <a :class="(data.status == 1) ? 'a_win' : ''">{{ data.countries[1].rating }}</a>
    </div>
  </div>
</template>

<style scoped>
.a_win{
  color: #F33E3E;
  text-align: right;
  font-size: 20px;
  font-style: normal;
  font-weight: 700;
  line-height: 20px; /* 83.333% */
  letter-spacing: 0.1px;
}

a{
  color: black;
  text-align: right;
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: 20px; /* 83.333% */
  letter-spacing: 0.1px;
  word-break: break-word;
}

.time_text{
  color: black;
  font-size: 20px;
  font-style: normal;
  font-weight: bold;
  line-height: 20px; /* 55.556% */
  letter-spacing: 0.1px;
  text-align: left;
}

.box{
  min-width: 300px;
  width: 300px;
  border-radius: 26px;
  border: 2px solid #D9D9D9;
  background: white;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  align-items: start;
  text-align: left;
  padding: 1rem 1.5rem;
  margin-bottom: 60px;
  margin-right: 25px;
}

.box2{
  min-width: 300px;
  width: 300px;
  border-radius: 26px;
  background: white;
  border: 2px solid #F33E3E;
  box-shadow: 0px 4px 4px 0px rgba(243, 62, 62, 0.5);
  display: flex;
  flex-direction: column;
  align-items: start;
  text-align: left;
  padding: 1rem 1.5rem;
  margin-bottom: 60px;
  margin-right: 25px;
}
</style>