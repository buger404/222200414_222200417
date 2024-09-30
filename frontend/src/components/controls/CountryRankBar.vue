<script setup lang="ts">
import CountryFlag from "@/components/controls/CountryFlag.vue";
const props = defineProps<{
  flag : string
  rank : number
  gold : number
  silver : number
  bronze : number
  total : number
  name : string
}>();
</script>

<template>
  <div style="position: relative; width: 100%">
    <h2 class="rank">{{rank}}</h2>
    <div class="box">
      <div style="width: 80px; text-align: center;">
        <CountryFlag :country="flag"/>
      </div>
      <div style="width: calc(35% - 80px);">
        <h2>{{name}}</h2>
      </div>
      <div style="width: 15%; text-align: center;">
        <h2 class="gold">{{gold}}</h2>
      </div>
      <div style="width: 15%; text-align: center;">
        <h2 class="silver">{{silver}}</h2>
      </div>
      <div style="width: 15%; text-align: center;">
        <h2 class="bronze">{{bronze}}</h2>
      </div>
      <div style="width: 15%; text-align: center;">
        <h2 class="total">{{total}}</h2>
      </div>
      <div style="width: 5%;"/>
    </div>
  </div>
</template>

<style scoped>
.rank{
  color: #F33E3E;
  text-align: center;
  font-size: 64px;
  font-style: italic;
  font-weight: 900;
  letter-spacing: 0.1px;
  position: absolute;
  right: 10px;
  bottom: -20px;
}

h2{
  margin-top: -8px;
}

.box{
  border-radius: 26px;
  border: 2px solid #F18787;
  background: white;
  box-shadow: 0px 4px 16px 0px rgba(255, 153, 153, 0.20);
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
  padding: 0.8rem 0;
  margin-bottom: 20px;
}

.gold{
  color: #FFDA35;
  text-align: center;
  font-size: 40px;
  font-style: italic;
  font-weight: 500;
  letter-spacing: 0.1px;
}

.silver{
  color: #B7CED2;
  text-align: center;
  font-size: 40px;
  font-style: italic;
  font-weight: 500;
  letter-spacing: 0.1px;
}

.bronze{
  color: #F2BA71;
  text-align: center;
  font-size: 40px;
  font-style: italic;
  font-weight: 500;
  letter-spacing: 0.1px;
}

.total{
  color: black;
  text-align: center;
  font-size: 40px;
  font-style: italic;
  font-weight: 900;
  letter-spacing: 0.1px;
}
</style>