<template>
  <div class="row h-100">
    <div class="col h-100">
      <div v-if="cardCounter.counter <= 4" class="row row-cols-md-2 justify-content-center align-items-center h-100">
        <CalloutCard v-for="(card, idx) in cards" :key="idx" ref="calloutCards" :state="'default'" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id" :cardSize="calloutCardSize" />
      </div>
      <div v-else-if="cardCounter.counter >= 4 && cardCounter.counter <= 6" class="row row-cols-md-3 justify-content-center align-items-center h-100">
        <CalloutCard v-for="(card, idx) in cards" :key="idx" ref="calloutCards" :state="'default'" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id" :cardSize="calloutCardSize" />
      </div>
      <div v-else class="row row-cols-md-4 justify-content-center align-items-center h-100">
        <CalloutCard v-for="(card, idx) in cards" :key="idx" ref="calloutCards" :state="'default'" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id" :cardSize="calloutCardSize" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CalloutCard from '../components/CalloutCard.vue';
import { reactive, ref, useTemplateRef, computed } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

type cardJson = {
  game_name: string,
  p1_name: string,
  p2_name: string,
  callout_id: string,
};

type cardUpdateJson = {
  callout_id: string,
  update_request: string,
};

const calloutCardRef = useTemplateRef<typeof CalloutCard>("calloutCards");
const cardCounter = reactive({counter: 0});

const cards = ref<cardJson[]>([]);


const calloutCardSize = computed(() => {
  console.log("Computing cardsize, card counter:", cardCounter.counter);

  if (cardCounter.counter === 0) {
    return "cardsize-null";
  }

  if (cardCounter.counter <= 4) {
    return "cardsize-large";
  } else if (cardCounter.counter <= 8) {
    return "cardsize-medium";
  } else {
    return "cardsize-small";
  }
});

const addNewCard = (msgJson: cardJson) => {
  cards.value.push(msgJson);
  cardCounter.counter = cards.value.length;
}

const removeCard = (card_id: string) => {
  const idx = cards.value.findIndex((card) => card.callout_id === card_id);
  if (idx !== -1) {
    cards.value.splice(idx, 1);
    cardCounter.counter = cards.value.length;
  } else {
    console.log("[ERROR] No card with id:", card_id, "found.");
  }
}

wails.EventsOn("callout/new", (msg: models.main.Message) => {
  console.log("Received callout new message:", msg);
  const msgJson: cardJson = JSON.parse(msg.message);  // This could be given a type since we know what the message is.
  console.log("Message as obj:", msgJson);

  addNewCard(msgJson);
});

wails.EventsOn("callout/update", (msg: models.main.Message) => {
  console.log("Received callout update message:", msg);

  const updateMsgJson: cardUpdateJson = JSON.parse(msg.message);

  if (calloutCardRef.value === null) {
    console.log("No callout cards in calloutCardRef, cannot update. Retry.");
    return
  }
  const cardToUpdate = calloutCardRef.value.find((card: typeof CalloutCard) => card.id === updateMsgJson.callout_id);
  console.log("Card to update:", cardToUpdate);

  if (cardToUpdate !== undefined) {
    cardToUpdate.changeStatus(updateMsgJson.update_request);
  } else {
    console.log("[ERROR] Card id:", updateMsgJson.callout_id, "not found, cannot update.");
  }

});

wails.EventsOn("callout/delete", (msg: models.main.Message) => {
  console.log("[INFO] Received callout delete message:", msg);
  removeCard(msg.message);
});
</script>

<style scoped>
h1 {
  font-family: 'Montserrat', sans-serif;
  font-weight: 800;
}
</style>
