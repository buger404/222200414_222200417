<script setup lang="ts">

</script>

<template>
  <div class="line_container">
    <div class="box1"/>
    <div class="line_container2">
      <div class="box2"/>
      <div style="height: 50%;"/>
    </div>
    <div class="dot" style="top: -6px; left: -6px;"/>
    <div class="dot" style="bottom: -6px; left: -6px;"/>
    <div class="dot" style="right: -6px; top: calc(50% - 6px - 1px);"/>
  </div>
</template>

<style scoped>
.dot {
  position: absolute;
  border-radius: 100px;
  background: #686868;
  width: 12px;
  height: 12px;
  z-index: 1;
}

.line_container {
  display: flex;
  width: 100%;
  height: 100%;
  flex-direction: row;
  position: relative;
}

.box1 {
  border-top: #686868 solid 2px;
  border-right: #686868 solid 2px;
  border-bottom: #686868 solid 2px;
  width: 50%;
  height: 100%;
}

.line_container2 {
  display: flex;
  width: 50%;
  height: 100%;
  flex-direction: column;
}

.box2 {
  border-bottom: #686868 solid 2px;
  width: 100%;
  height: 50%;
}

</style>