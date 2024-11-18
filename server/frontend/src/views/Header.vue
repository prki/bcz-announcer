<template>
  <div class="row border mt-1">
    <div class="col">
      <h2 class="text-center">Modify header content</h2>
      <form>
        <div class="row">
          <div class="col">
            <label for="header-text" class="form-label">Header text</label>
            <input type="text" class="form-control" id="header-text" v-model="state.headerText">
          </div>
        </div>
        <div class="row mt-2 d-flex">
          <div class="col">
            <div class="text-center">
              <button type="button" class="btn btn-primary" @click="updateHeaderContent()">Update header content</button>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>

  <div class="row border mt-4">
    <div class="col">
      <h2 class="text-center">Header content guide</h2>
      <p>This lets the TO desk person modify the header of the renderer. It is to be used to call out the currently running tournament game.</p>
      <p>Format which should be used for games is as follows:</p>
      <ul>
        <li>&lt;b&gt;TEKKEN 8 POOLS&lt;/b&gt; -- in case of a single running tournament</li>
        <li>&lt;b&gt;TEKKEN 8 POOLS // GUILTY GEAR XRD TOP 4&lt;/b&gt; -- in case of concurrently running tournaments</li>
      </ul>
      <p>Meaning that the text is sent in bold tag, all text is uppercase. Names of games to be used are:</p>
      <ul>
        <li>Guilty Gear Xrd: GUILTY GEAR XRD</li>
        <li>Street Fighter 6: STREET FIGHTER 6</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'; 
import { SendHeaderUpdate } from '../../wailsjs/go/main/App';

type HeaderContent = {
  headerText: string,
};

const state = reactive<HeaderContent>({headerText: ''});

function updateHeaderContent() {
  console.log("[INFO] Updating header content. New text:", state.headerText);
  SendHeaderUpdate(state.headerText);
}
</script>

<style scoped>
</style>
