<template>
  <div class="row border mt-1"> 
    <h2>Create new callout card</h2>
    <div class="col">
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
              <option value="gov">God of Vidya</option>
            </select>
          </div>
        </div>
        <div class="row mb-1">
          <div class="col-6">
            <label for="match-p1" class="form-label">Player 1 name</label>
            <!--<input type="text" class="form-control" id="match-p1" v-model="state.p1Name">-->
            <select v-model="state.p1Name" id="match-p1" class="form-select">
              <option v-for="(value, key, idx) in playerMap" :key="idx" :value="key">{{ key }}</option>
            </select>
          </div>
          <div class="col-6">
            <label for="match-p2" class="form-label">Player 2 name</label>
            <!--<input type="text" class="form-control" id="match-p2" v-model="state.p2Name">-->
            <select v-model="state.p2Name" id="match-p2" class="form-select">
              <option v-for="(value, key, idx) in playerMap" :key="idx" :value="key">{{ key }}</option>
            </select>
          </div>
        </div>
        <div class="row mt-2 d-flex justify-content-center">
          <div class="col">
            <div class="text-center">
              <button type="button" class="btn btn-primary" @click="createNewCallout()">Create callout card</button>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { calloutStore } from '../store/store';
import { toastStore, Toast } from '../store/toastStore';
import * as models from '../../wailsjs/go/models';
import { SendUpdateCallout, GetPlayers } from '../../wailsjs/go/main/App';

type Player = {
  playerName: string,
  countryCode: string
}

const state = reactive({
  p1Name: '',
  p2Name: '',
  gameName: 'bbcf',
});

const playerMap = reactive<Record<string, Player>>({})

GetPlayers().then((res: models.main.Player[]) => {
  console.log("Got players:", res);
  res.map((player: models.main.Player) => {
    playerMap[player.player_name] = {playerName: player.player_name, countryCode: player.country_code};
  })
});

// Function creating new callout.
// Logic is as follows:
//   * Pass info to backend - will store information.
//   * Backend returns ID of the card created.
//     -> Ideally, this should only be done when the backend verifies the renderer received this!
//   * Emit callout created in order for parent to put match into listing.
function createNewCallout() {
  console.log("Selected game:", state.gameName);
  if (calloutStore.cardCount >= calloutStore.maxCards) {
    //alert("Cannot create more than " + calloutStore.maxCards + " callout cards!");
    toastStore.addToast({
      headerMessage: "Callout create error",
      message: "Cannot create more than " + calloutStore.maxCards + " callout cards!",
      toastBg: "bg-danger"
    });
    return;
  }

  console.log("Creating new callout, state:", state);
  // [TODO] Rename to SendNewCallout or something akin to that.
  SendUpdateCallout(
    playerMap[state.p1Name].playerName,
    playerMap[state.p1Name].countryCode,
    playerMap[state.p2Name].playerName,
    playerMap[state.p2Name].countryCode,
    state.gameName
  ).then((result) => {
    console.log("[DEBUG] [MatchCalloutCreate] Go returned:", result);
  });
}
</script>

<style scoped>

</style>
