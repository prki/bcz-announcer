<template>
  <div class="row">
    <div class="col">
      <!--<p>Message will be displayed here.</p>
      <p>Message action: {{ receivedMsg.action }}</p>
      <p>Message content: {{ receivedMsg.message }}</p>
      -->
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
//import { Message } from '../../wailsjs/go/models';
import * as models from '../../wailsjs/go/models';

const receivedMsg = reactive<models.main.Message>({
  action: "",
  message: "",
});


wails.EventsOn("message/display", (msg: models.main.Message) => {
  //console.log("action: ", msg.action, " message: ", msg.message);
  receivedMsg.action = msg.action;
  receivedMsg.message = msg.message;
});

// Commented since not relevant + should be on top level. Kept for pattern
// recognition, [DELME] when confident in how this works.
/*wails.EventsOn("control/change_view", (msg: models.main.Message) => {
  console.log("Received change view request:", msg);
});
*/
</script>

<style scoped>
</style>
