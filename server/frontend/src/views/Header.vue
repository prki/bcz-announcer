<template>
  <div class="row border bg-light mt-1">
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
