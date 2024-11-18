<template>
  <div class="row border">
    <div class="col">
      <h2>Edit Footer</h2>
      <form>
        <div class="row">
          <div class="col">
            <label for="footer-text" class="form-label">Footer content</label>
            <textarea type="text" class="form-control" id="footer-text" rows="4" v-model="state.textcontent" placeholder="<p>Welcome to <b>Buttonczech 2024</b>!</p>"></textarea>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <button type="button" class="btn btn-primary" @click="sendFooterChange()">Change Footer</button>
          </div>
        </div>
      </form>
    </div>
  </div>

  <div class="row border mt-4">
    <div class="col">
      <p>Edit the bottom footer of the renderer. Text can be sent as direct text or direct HTML markup. Note that this <b>should only be used</b> for styling purposes - e.g. using bold/italics tags.</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { SendFooterUpdate } from '../../wailsjs/go/main/App';

const state = reactive({
  textcontent: '',
});

// [TODO] Consider if there's a way to have a max character limit - it's unreasonable to expect footer
// to have 2000 characters, for instance.
function sendFooterChange() {
  console.log("[INFO] Sending footer change. Current content:", state.textcontent);
  SendFooterUpdate(state.textcontent);
}
</script>

<style scoped>
</style>
