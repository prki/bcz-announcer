<template>
  <div class="row">
    <div class="col gap-1 d-flex">
      <button class="btn btn-primary" @click="changeView('/')">Change view Main</button>
      <button class="btn btn-primary" @click="changeView('/matchcallout')">Change view MatchCallout</button>
      <button class="btn btn-primary" @click="changeView('/textdisplay')">Change view TextDisplay</button>
    </div>
  </div>

  <div class="row">
    <div class="col">
      <h2>Main view form</h2>
      <form>
        <p>TODO</p>
      </form>
    </div>
    <div class="col">
      <MatchCallout />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { SendChangeViewRequest } from '../../wailsjs/go/main/App';
import MatchCallout from './MatchCallout.vue';
import * as models from '../../wailsjs/go/models';

type SetupMatch = {
  p1Name: string,
  p2Name: string
};

const state = reactive({
  A: {
    p1Name: 'Player 1 name',
    p2Name: 'Player 2 name'
  } as SetupMatch
});

function changeView(viewName: string) {
  console.log("Requesting view change to:", viewName);
  SendChangeViewRequest(viewName);
}

function updateCalloutView() {
  console.log("Requesting update callout view");
  console.log("setup A: ", state.A);
  const data = {
    p1_name: state.A.p1Name,
    p2_name: state.A.p2Name
  } as models.main.SetupMatch;
  //SendUpdateCalloutRequest(data);
}

// [TODO] DELME, kept for pattern recognition
/*function sendMessage() {
  console.log("sending message");
  SendMessage().then(result => {
    console.log("Message sent")
  });
}
*/
</script>

<style scoped>
</style>
