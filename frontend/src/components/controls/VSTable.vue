<script setup lang="ts">
import type {Table} from "@/models/tableList";
import CountryFlag from "@/components/controls/CountryFlag.vue";
import {useTableStore} from "@/stores/counter";

const store = useTableStore();

const props = defineProps<{
  data : Table
}>();

function setActiveCountry(index : number) {
  store.country = props.data.countries[index].name;
}

function clearActiveCountry(index : number) {
  if (!isActiveCountry(index)) {
    return;
  }
  store.country = '';
}

function isActiveCountry(index : number) {
  return store.country && store.country == props.data.countries[index].name;
}
</script>

<template>
  <div class="bar" style="width: 100%;">
    <div class="sub_bar" v-if="data.period">
      <h3 style="font-weight: bold; margin-bottom: 10px; margin-top: -5px;"> {{ data.period }}</h3>
    </div>
    <div class="sub_bar">
      <div style="flex-grow: 1;">
        <div
            :class="isActiveCountry(0) ? 'rating_bar_active' : 'rating_bar'"
            @mouseenter="setActiveCountry(0)"
            @mouseleave="clearActiveCountry(0)"
        >
          <div class="row">
            <CountryFlag :country="data.countries[0].flag"/>
            <h3 class="country_name" :style="'color: ' + (data.winner == 1 ? '#F33E3E' : 'black')"> {{ data.countries[0].name }} </h3>
          </div>
          <h3 class="rating" :style="'color: ' + (data.winner == 1 ? '#F33E3E' : 'black')"> {{ data.countries[0].rating }} </h3>
        </div>
        <div
            :class="isActiveCountry(1) ? 'rating_bar_active' : 'rating_bar'"
            style="margin-top: 10px;"
            @mouseenter="setActiveCountry(1)"
            @mouseleave="clearActiveCountry(1)"
        >
          <div class="row">
            <CountryFlag :country="data.countries[1].flag"/>
            <h3 class="country_name" :style="'color: ' + (data.winner == 2 ? '#F33E3E' : 'black')"> {{ data.countries[1].name }} </h3>
          </div>
          <h3 class="rating" :style="'color: ' + (data.winner == 2 ? '#F33E3E' : 'black')"> {{ data.countries[1].rating }} </h3>
        </div>
      </div>
      <div v-if="data.special"
           style="text-align: center; display: flex; justify-content: center; flex-direction: column; margin-left: 15px; margin-right: -5px;">
        <h2 style="font-weight: bold;"> {{ data.special }}</h2>
      </div>
    </div>
  </div>
</template>

<style scoped>
.country_name {
  margin-left: 10px;
}

h3 {
  line-height: 24px;
}

.rating {
  color: black;
  text-align: right;
  font-size: 24px;
  font-style: italic;
  font-weight: 900;
  line-height: 24px; /* 83.333% */
  letter-spacing: 0.1px;
}

.sub_bar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.bar{
  display: flex;
  flex-direction: column;
  padding: 20px 25px;
  border-radius: 26px;
  border: 2px solid #D9D9D9;
  background: white;
  box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.15);
  transition: 0.5s all;
}

.bar:hover{
  border: 2px solid #F33E3E;
  box-shadow: 0px 4px 4px 0px rgba(243, 62, 62, 0.5);
}

.row {
  display: flex;
  flex-direction: row;
}

.rating_bar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  background: white;
  transition: 0.5s all;
}

.rating_bar_active {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  scale: 1.1;
  transition: 0.5s all;
}
</style>