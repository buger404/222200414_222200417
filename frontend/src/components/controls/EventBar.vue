<script setup lang="ts">
import type {DayEvent} from "@/models/dayEventList";
import LargeCountryFlag from "@/components/controls/LargeCountryFlag.vue";
import IconVS from "@/components/icons/IconVS.vue";

const props = defineProps<{
  data : DayEvent
}>();
</script>

<template>
<div class="box">
  <div class="head" style="width: 35%; padding-left: 30px;">
    <div>
      <a class="time_text">{{ data.time }}</a>
      <a class="event_text">{{ data.group }}</a>
    </div>
    <div style="margin-top: 0.9rem">
      <a class="group_text">{{ data.event }}</a>
    </div>
  </div>
  <div style="width: 15%; text-align: right; padding-right: 20px;">
    <a :class="(data.winner == 1) ? 'a_win' : ''">{{ data.countries[0].name }}</a>
  </div>
  <div style="width: 5%;">
    <div class="column">
      <LargeCountryFlag :country="data.countries[0].flag" style="margin-top: 0.5rem; margin-bottom: 0.5rem;"/>
      <a :class="(data.winner == 1) ? 'a_win' : ''">{{ data.countries[0].rating }}</a>
    </div>
  </div>
  <div style="width: 15%; text-align: center;">
    <IconVS style="width: 58px; height: 58px"/>
  </div>
  <div style="width: 5%;">
    <div class="column">
      <LargeCountryFlag :country="data.countries[1].flag" style="margin-top: 0.5rem; margin-bottom: 0.5rem;"/>
      <a :class="(data.winner == 2) ? 'a_win' : ''">{{ data.countries[1].rating }}</a>
    </div>
  </div>
  <div style="width: 15%; text-align: left; padding-left: 20px;">
    <a :class="(data.winner == 2) ? 'a_win' : ''">{{ data.countries[1].name }}</a>
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
}

.column{
  display: flex;
  flex-direction: column;
  align-items: center;
}

.head{
  display: flex;
  flex-direction: column;
}

.group_text {
  color: black;
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: 20px; /* 83.333% */
  letter-spacing: 0.1px;
}

.event_text{
  color: black;
  font-size: 24px;
  font-style: normal;
  font-weight: 700;
  line-height: 20px; /* 62.5% */
  letter-spacing: 0.1px;
  margin-left: 5px;
}

.time_text{
  color: #F33E3E;
  font-size: 24px;
  font-style: normal;
  font-weight: 500;
  line-height: 20px; /* 55.556% */
  letter-spacing: 0.1px;
}

.box{
  border-radius: 26px;
  border: 2px solid #F18787;
  background: white;
  box-shadow: 0px 4px 16px 0px rgba(243, 62, 62, 0.25);
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
  padding: 1.5rem 0;
  margin-bottom: 20px;
}
</style>