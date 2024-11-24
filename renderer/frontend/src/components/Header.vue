<template>
  <div class="row header-row-size mt-4 align-items-stretch">
    <div class="col-8">
      <!--Separated rows, ignore for now.-->
      <!--<div class="border p-3 pb-4 pt-4">
        <div class="row">
          <div class="col-6 offset-md-1 border">
            <div v-html="state.headerText"></div>
          </div>
          <div class="col-3 offset-md-1 border">
            <div v-html="state.headerImage"></div>
          </div>
        </div>
      </div>
      -->
      <div class="row h-100">
        <div class="col h-100">
          <div class="border h-100 d-flex align-items-center header-text justify-content-center px-2">
            <div v-html="state.headerText"></div>
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
  headerText: '<b>Welcome to BUTTONCZECH 2024!</b>',
  headerImage: '<i>Header image</i>'
});

wails.EventsOn("header/update", (msg: models.main.Message) => {
  console.log("Received header update request:", msg);
  state.headerText = msg.message;
});
</script>

<style scoped>
.header-text {
  font-family: 'Montserrat', sans-serif;
  font-weight: 800;
  font-size: 36px;
  color: white;
}
</style>
