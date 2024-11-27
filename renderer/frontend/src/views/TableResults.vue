<template>
  <div class="row">
    <div class="col">
      <h1 class="text-center mb-4"><strong>GOD GAMER OF BUTTONCZECH 2024</strong></h1>
      <div class="row">
        <div class="col d-flex flex-grow-1 align-items-center">
          <!--<h1 class="text-center">GGOB Results</h1>-->
          <table class="table">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Player name</th>
                <th scope="col">Points</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, idx) in state" :key="idx">
                <th scope="row">{{ idx + 1 }}</th>
                <td>{{ row.player_name }}</td>
                <td>{{ row.points }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

type PlayerScore = {
  player_name: string,
  points: number
};

const state = ref<PlayerScore[]>([]);

wails.EventsOn("tableresults/update", (msg: models.main.Message) => {
  console.log("Received tableresults update message:", msg);
  state.value = JSON.parse(msg.message);
});

//const state = ref<PlayerScore[]>([]);
</script>

<style scoped>
.table {
  --bs-table-bg: transparent !important;
  color: white !important;
  font-size: 24px !important;
}

td, th, caption {
  color: white;
}

th {
  font-size: 28px !important;
}

.col {
  font-family: 'Montserrat', sans-serif;
}

.table>tbody>tr>td,
.table>tbody>tr>th {
  border-top: none;
  border-bottom: none;
}
</style>
