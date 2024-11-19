<template>
  <div class="row border">
    <div class="col">
      <h2 class="text-center mb-2">Edit Textdisplay content</h2>
      <form>
        <div class="row">
          <div class="col">
            <label for="textdisplay-text" class="form-label">Textdisplay content (HTML)</label>
            <textarea type="text" class="form-control" id="footer-text" rows="4" v-model="state.textcontent" placeholder="<p>Next event starting in 30 minutes!</p>"></textarea>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <button type="button" class="btn btn-primary" @click="sendTextdisplayUpdate()">Change Textdisplay content</button>
          </div>
        </div>
      </form>
    </div>
  </div>

  <div class="row border mt-4">
    <div class="col">
      <h3 class="text-center mb-2">Guide</h3>
      <p>Edit content in TextDisplay view. Text can be sent as direct text or direct HTML markup. Note that this <b>should only be used</b> for styling purposes - e.g. using bold/italics tags.</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { SendTextdisplayUpdate } from '../../wailsjs/go/main/App';

const state = reactive({
  textcontent: '',
});

function sendTextdisplayUpdate() {
  console.log("[INFO] Sending textdisplay content change request. Current content:", state.textcontent);
  SendTextdisplayUpdate(state.textcontent);
}
</script>

<style scoped>
</style>
