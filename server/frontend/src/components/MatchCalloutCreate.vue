<template>
  <div class="row"> 
    <div class="col">
      <h3>I am Match Callout Create component</h3>
      <form>
        <div class="row">
          <div class="col">
            <label for="gamename" class="form-label">Game title</label>
            <select v-model="state.gameName" class="form-select" id="gamename">
              <option value="bbcf">Blazblue Central Fiction</option>
              <option value="strive">Guilty Gear Strive</option>
              <option value="xrd">Guilty Gear Xrd</option>
              <option value="plusr">Guilty Gear XX ACPR</option>
              <option value="sf6">Street Fighter 6</option>
              <option value="tekken">Tekken 8</option>
            </select>
          </div>
        </div>
        <div class="row">
          <div class="col-6">
            <label for="match-p1" class="form-label">Player 1 name</label>
            <input type="text" class="form-control" id="match-p1" v-model="state.p1Name">
          </div>
          <div class="col-6">
            <label for="match-p2" class="form-label">Player 2 name</label>
            <input type="text" class="form-control" id="match-p2" v-model="state.p2Name">
          </div>
          <button type="button" class="btn btn-primary" @click="createNewCallout()">Create callout card</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as models from '../../wailsjs/go/models';
import { SendUpdateCallout } from '../../wailsjs/go/main/App';

const state = reactive({
  p1Name: 'Player 1 name',
  p2Name: 'Player 2 name',
  gameName: 'gameNamePlaceholder',
});

// Function creating new callout.
// Logic is as follows:
//   * Pass info to backend - will store information.
//   * Backend returns ID of the card created.
//     -> Ideally, this should only be done when the backend verifies the renderer received this!
//   * Emit callout created in order for parent to put match into listing.
function createNewCallout() {
  console.log("Selected game:", state.gameName);
  SendUpdateCallout(state.p1Name, state.p2Name, state.gameName).then((result) => {
    console.log("[DEBUG] [MatchCalloutCreate] Go returned:", result);
  });
}
</script>

<style scoped>

</style>
