<script setup lang="ts">
import type {Table} from "@/models/tableList";
import CountryFlag from "@/components/controls/CountryFlag.vue";

const props = defineProps<{
  data : Table
}>();
</script>

<template>
  <div class="bar" style="width: 100%;">
    <div style="flex-grow: 1;">
      <div class="rating_bar">
        <div class="row">
          <CountryFlag :country="data.countries[0].flag"/>
          <h3 class="country_name" :style="'color: ' + (data.winner == 1 ? '#F33E3E' : 'black')"> {{ data.countries[0].name }} </h3>
        </div>
        <h3 class="rating" :style="'color: ' + (data.winner == 1 ? '#F33E3E' : 'black')"> {{ data.countries[0].rating }} </h3>
      </div>
      <div class="rating_bar" style="margin-top: 10px;">
        <div class="row">
          <CountryFlag :country="data.countries[1].flag"/>
          <h3 class="country_name" :style="'color: ' + (data.winner == 2 ? '#F33E3E' : 'black')"> {{ data.countries[1].name }} </h3>
        </div>
        <h3 class="rating" :style="'color: ' + (data.winner == 2 ? '#F33E3E' : 'black')"> {{ data.countries[1].rating }} </h3>
      </div>
    </div>
    <div v-if="data.special">
      <h3> {{ data.special }}</h3>
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

.bar{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
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
}
</style>