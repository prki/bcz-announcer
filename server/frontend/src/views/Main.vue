<template>
  <div class="row">
    <div class="col">
      <button class="btn btn-primary" @click="sendMessage">Button</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { SendMessage } from '../../wailsjs/go/main/App';

function sendMessage() {
  console.log("sending message");
  SendMessage().then(result => {
    console.log("Message sent")
  });
}
</script>

<style scoped>
</style>
