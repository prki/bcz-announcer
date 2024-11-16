<template>
  <div class="row"> 
    <div class="col">
      <h3>I am match callout listing component</h3>

      <h4>Created cards:</h4>
      <ul class="list-group">
        <li class="list-group-item" v-for="(card, idx) in cards" :key="idx">{{card.p1Name}} vs {{ card.p2Name }} || {{ card.ackState }}</li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';

type CardDescription = {
  p1Name: string,
  p2Name: string,
  id: string,
  ackState: string
};

const cards = reactive<CardDescription[]>([]);

wails.EventsOn("callouts/new", (match) => {
  console.log("Adding match into list:", match);
  const newCard: CardDescription = {
    p1Name: match.p1_name,
    p2Name: match.p2_name,
    id: match.callout_id,
    ackState: "PENDING",
  };
  cards.push(newCard);
});
</script>

<style scoped>
</style>
