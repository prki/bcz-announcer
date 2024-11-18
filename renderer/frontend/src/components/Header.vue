<template>
  <div class="row header-row-size mt-4 gx-4">
    <div class="col-8">
      <div class="border p-3 pb-4 pt-4">
        <div class="row">
          <div class="col-6 offset-md-1 border">
            <div v-html="state.headerText"></div>
          </div>
          <div class="col-3 offset-md-1 border">
            <div v-html="state.headerImage"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="col-4">
      <div class="border">
        <img src="../assets/images/bcz_logo_horizontal.svg" class="img-fluid">
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

const state = reactive({
  headerText: '<b>Header text static placeholder</b>',
  headerImage: '<i>Header image</i>'
});

wails.EventsOn("header/update", (msg: models.main.Message) => {
  console.log("Received header update request:", msg);
  state.headerText = msg.message;
});
</script>

<style scoped>
</style>
