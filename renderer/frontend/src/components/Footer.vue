<template>
  <div class="row">
    <div class="col">
      <div class="border mt-4">
        <div v-html="state.footerHtml"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

const state = reactive({
  footerHtml: "<p>Welcome to <b>Buttonczech 2024</b>!</p>"
});

wails.EventsOn("footer/update", (msg: models.main.Message) => {
  console.log("Received footer update request, msg:", msg);

  state.footerHtml = msg.message;
});
</script>

<style scoped>
</style>
