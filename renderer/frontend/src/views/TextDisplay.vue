<template>
  <div class="row">
    <div class="col">
      <div v-html="state.content" class="p-4"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as models from '../../wailsjs/go/models';
import * as wails from '../../wailsjs/runtime/runtime.js';

const state = reactive({
  content: ''
});

wails.EventsOn("textdisplay/update", (msg: models.main.Message) => {
  console.log("Received text display update request, msg:", msg);
  state.content = msg.message;
});
</script>

<style scoped>
.col {
  font-family: 'Montserrat', sans-serif;
  font-weight: 500;
  font-size: 3rem;
}
</style>
