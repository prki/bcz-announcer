<template>
  <div class="row">
    <div class="col">
      <h1>I am text display view</h1>
      <p>{{ state.content }}</p>
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
</style>
