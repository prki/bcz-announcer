<template>
  <div class="row">
    <div class="col">
      <div class="border mt-4 footer-text px-3">
        <div v-html="state.footerHtml" class="d-flex align-items-center justify-content-center"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

const state = reactive({
  footerHtml: "<span>Welcome to <b>Buttonczech 2024</b>!</span>"
});

wails.EventsOn("footer/update", (msg: models.main.Message) => {
  console.log("Received footer update request, msg:", msg);

  state.footerHtml = msg.message;
});
</script>

<style scoped>
.footer-text {
  font-family: 'Montserrat', sans-serif;
  font-weight: 500;
  font-size: 1.5rem;
  color: white;
}
</style>
