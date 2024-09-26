<script setup lang="ts">
import {ref} from "vue";

const props = defineProps<{
  country : string
}>();

const isLoaded = ref(false);

function onLoad(){
  isLoaded.value = true;
}
</script>

<template>
  <div v-loading="!isLoaded">
    <img
        @load="onLoad()"
        class="flag"
        :style="isLoaded ? 'display: inline-flex;' : 'display: none;'"
        :src="'https://olympics.com/OG2024/assets/images/flags/OG2024/' + country + '.webp'"
        :alt="country"
    />
  </div>
</template>

<style scoped>
.flag{
  box-shadow: 0px 4px 8px 0px rgba(0, 0, 0, 0.40);
  width: 36px;
  height: 24px;
}

div{
  --el-mask-color: transparent;
  --el-loading-spinner-size: 36px;
  --el-color-primary: #f33e3e;
}
</style>