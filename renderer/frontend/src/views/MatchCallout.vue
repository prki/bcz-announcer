<template>
  <div class="row h-100">
    <div class="col h-100">
      <div class="row justify-content-center align-items-center h-100" :class="rowLayout">
        <CalloutCard v-for="(card, idx) in cards" :key="idx" ref="calloutCards" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id" :cardSize="calloutCardSize" :status="card.status" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CalloutCard from '../components/CalloutCard.vue';
import { reactive, ref, useTemplateRef, computed, toRefs } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

type cardJson = {
  game_name: string,
  p1_name: string,
  p2_name: string,
  callout_id: string,
  status: string,
};

type cardUpdateJson = {
  callout_id: string,
  update_request: string,
};

const calloutCardRef = useTemplateRef<typeof CalloutCard>("calloutCards");
const cardCounter = reactive({counter: 0});

const cards = reactive<cardJson[]>([]);

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

// Computed property governing CSS style of row/col layout.
const rowLayout = computed(() => {
  if (cardCounter.counter === 0) {
    return "row-cols-md-2";
  }

  if (cardCounter.counter <= 4) {
    return "row-cols-md-2";
  } else if (cardCounter.counter > 4 && cardCounter.counter <= 6) {
    return "row-cols-md-3";
  } else if (cardCounter.counter === 9) {
    return "row-cols-md-3";
  } else {
    return "row-cols-md-4";
  }
});

const addNewCard = (msgJson: cardJson) => {
  cards.push(msgJson);
  cardCounter.counter = cards.length;
}

const removeCard = (card_id: string) => {
  const idx = cards.findIndex((card) => card.callout_id === card_id);
  if (idx !== -1) {
    cards.splice(idx, 1);
    cardCounter.counter = cards.length;
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

  const cardToUpdateIdx = cards.findIndex((card) => card.callout_id === updateMsgJson.callout_id);

  // Update done as delete of old + replace by new.
  // Possibly easier with some deep compute/ref, but this is easier to see.
  if (cardToUpdateIdx !== -1) {
    //cards[cardToUpdateIdx].status = updateMsgJson.update_request;
    const updatedCard = { ...cards[cardToUpdateIdx], status: updateMsgJson.update_request };
    cards.splice(cardToUpdateIdx, 1, updatedCard);
    console.log("Successfully updated card to:", cards[cardToUpdateIdx]);
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
