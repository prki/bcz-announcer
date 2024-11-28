<template>
  <div class="row mb-4">
    <h2 class="">Renderer controls</h2>
    <div class="col gap-4 d-flex">
      <button class="btn btn-primary" @click="changeView('/')">Change view Banner</button>
      <button class="btn btn-primary" @click="changeView('/matchcallout')">Change view MatchCallout</button>
      <button class="btn btn-primary" @click="changeView('/textdisplay')">Change view TextDisplay</button>
      <button class="btn btn-primary" @click="changeView('/tableresults')">Change view TableResults</button>
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
