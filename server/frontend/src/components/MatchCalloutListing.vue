<template>
  <div class="row"> 
    <div class="col">
      <h3>I am match callout listing component</h3>

      <h4>Created cards:</h4>
      <ul class="list-group">
        <li class="list-group-item d-flex justify-content-between" v-for="(card, idx) in cards" :key="idx">
          {{card.p1Name}} vs {{ card.p2Name }}
          <div class="d-flex gap-1">
            <button type="button" class="btn btn-success" @click="sendColorUpdate(card, 'default')">Default color</button>
            <button type="button" class="btn btn-warning" @click="sendColorUpdate(card, 'dq')">Color DQ</button>
            <button type="button" class="btn btn-danger" @click="sendCardRemove(card)">Remove card</button>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import { DeleteCalloutCard } from '../../wailsjs/go/main/App';

type CardDescription = {
  p1Name: string,
  p2Name: string,
  id: string,
  ackState: string
};

const cards = reactive<CardDescription[]>([]);

wails.EventsOn("callout/new", (match) => {
  console.log("Adding match into list:", match);
  const newCard: CardDescription = {
    p1Name: match.p1_name,
    p2Name: match.p2_name,
    id: match.callout_id,
    ackState: "PENDING",
  };
  cards.push(newCard);
});

// [TODO] Isn't there a race condition? Is there only one event handled at once guaranteed? Should be? I think so?
wails.EventsOn("callout/delete", (cardId) => {
  console.log("[DEBUG] Attempting to delete card id:", cardId);
  const idx = cards.findIndex((card) => card.id === cardId);
  if (idx !== -1) {
    cards.splice(idx, 1);
  }
  // [TODO] Consider some logging component on frontend
  else {
    console.log("[ERROR] Attempted to delete callout card, but not found");
  }
});

function sendColorUpdate(card: CardDescription, colorType: string) {
  console.log("[DEBUG] Sending color update request. Card:", card, "colorType:", colorType);
}

function sendCardRemove(card: CardDescription) {
  console.log("[DEBUG] Sending card remove request. Card:", card);
  DeleteCalloutCard(card.id);
}
</script>

<style scoped>
</style>
