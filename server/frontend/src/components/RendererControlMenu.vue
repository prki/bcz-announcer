<template>
  <div class="row mb-4">
    <h2>Renderer controls</h2>
    <div class="col gap-1 d-flex">
      <button class="btn btn-primary" @click="changeView('/')">Change view Main</button>
      <button class="btn btn-primary" @click="changeView('/matchcallout')">Change view MatchCallout</button>
      <button class="btn btn-primary" @click="changeView('/textdisplay')">Change view TextDisplay</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { SendChangeViewRequest } from '../../wailsjs/go/main/App';

function changeView(viewName: string) {
  console.log("Requesting view change to:", viewName);
  SendChangeViewRequest(viewName);
}
</script>

<style scoped>
</style>
