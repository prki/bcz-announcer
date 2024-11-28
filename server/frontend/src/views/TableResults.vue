<template>
  <div class="row">
    <div class="col">
      <h1>Modify GGOB Results</h1>
      <div class="row">
        <div class="col d-flex gap-4">
          <button class="btn btn-primary" @click="previewCsv()">Preview CSV</button>
          <button class="btn btn-success" @click="sendCsv()">Send CSV</button>
        </div>
      </div>
    </div>
  </div>

  <!--[TODO] Possibly create a component, we'll see.-->
  <div class="row">
    <div class="col">
      <h1>GGOB Results preview</h1>
      <table class="table table-sm">
        <caption>
          Preview created on {{ previewCreatedTime }}
        </caption>
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
            <td><span class="truncated-row text-truncate">{{ row.player_name }}</span></td>
            <td>{{ row.points }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

  <div class="row">
    <div class="col">
      <h1>Guide</h1>
      <div class="border">
        <p class="p-2">This section enables to upload GGOB result data. The result data is provided via the file <span class="font-monospace">ggob.csv</span> located in the folder with this application.</p>
        <p class="p-2">Before uploading the data to the renderer, run <strong>preview</strong> to validate that the data is handled properly. If not, the data will need to be fixed.</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';
import * as models from '../../wailsjs/go/models';
import { ReadGGOBCsvPreview, SendGGOBCsv } from '../../wailsjs/go/main/App';

const state = reactive<models.main.GGOBRow[]>([]);
const previewCreatedTime = ref('');


const previewCsv = () => {
  ReadGGOBCsvPreview().then((res: models.main.Message) => {
    const rows: models.main.GGOBRow[] = JSON.parse(res.message);

    state.splice(0, state.length, ...rows);
    state.splice(15, state.length);  // Only keep 15 best results - issue - what if there's more people on 15th place? Probably try to take the one first alpabetically?
    previewCreatedTime.value = new Date().toLocaleString();
  });
}

const sendCsv = () => {
  SendGGOBCsv();
}
</script>

<style scoped>
.table {
  --bs-table-bg: transparent !important;
  color: white !important;
}

.truncated-row {
  display: inline-block;
  max-width: 25rem;
}

td, th, caption {
  color: white;
}

/*th {
  color: white;
}
*/
</style>
