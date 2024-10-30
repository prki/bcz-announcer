<template>
  <div class="row">
    <div class="col">
      <button class="btn btn-primary" @click="changeView('/')">Change view Main</button>
      <button class="btn btn-primary" @click="changeView('/matchcallout')">Change view MatchCallout</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
// [TODO] DELME
import { SendMessage } from '../../wailsjs/go/main/App';
import { SendChangeViewRequest } from '../../wailsjs/go/main/App';

function changeView(viewName: string) {
  console.log("Requesting view change to:", viewName);
  SendChangeViewRequest(viewName);
}

// [TODO] DELME, kept for pattern recognition
function sendMessage() {
  console.log("sending message");
  SendMessage().then(result => {
    console.log("Message sent")
  });
}
</script>

<style scoped>
</style>
