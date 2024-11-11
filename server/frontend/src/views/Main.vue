<template>
  <div class="row">
    <div class="col">
      <button class="btn btn-primary" @click="changeView('/')">Change view Main</button>
      <button class="btn btn-primary" @click="changeView('/matchcallout')">Change view MatchCallout</button>
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
      <!--<h2>Match callout view form</h2>
      <form>
        <div class="row">
          <label for="match-a-p1" class="form-label">[A] Player 1 name</label>
          <div class="col">
            <input type="text" class="form-control" id="match-a-p1" v-model="state.A.p1Name">
          </div>
          <div class="col">
            <input type="text" class="form-control" id="match-a-p2" v-model="state.A.p2Name">
          </div>
          <button type="button" class="btn btn-primary" @click="updateCalloutView()">Update view</button>
        </div>
      </form>
      -->
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
