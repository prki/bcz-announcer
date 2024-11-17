<template>
  <div class="row"> 
    <div class="col">
      <h1>I am match callout view</h1>
    </div>
  </div>
  <div class="row">
    <div class="col">
      <div class="row row-cols-md-4">
        <!--<div v-for="(card, idx) in cards" :key="idx" ref="calloutCards">
          <CalloutCard :state="'default'" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id"/>
        </div>
        -->
        <CalloutCard v-for="(card, idx) in cards" :key="idx" ref="calloutCards" :state="'default'" :gameName="card.game_name" :p1Name="card.p1_name" :p2Name="card.p2_name" :id="card.callout_id" />
        <!--<CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>

        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        -->
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CalloutCard from '../components/CalloutCard.vue';
import { reactive, ref, useTemplateRef } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

/*const state = reactive({
  A: {
    p1Name: '',
    p2Name: ''
  }
});
*/

//type CalloutCardType = InstanceType<typeof CalloutCard>;

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

const calloutCardRef = useTemplateRef<CalloutCard>("calloutCards");

const cards = ref<cardJson[]>([]);

wails.EventsOn("callout/new", (msg: models.main.Message) => {
  console.log("Received callout new message:", msg);
  const msgJson: cardJson = JSON.parse(msg.message);  // This could be given a type since we know what the message is.
  console.log("Message as obj:", msgJson);

  cards.value.push(msgJson);

  //state.A.p1Name = msgJson.p1_name;
  //state.A.p2Name = msgJson.p2_name;
});

wails.EventsOn("callout/update", (msg: models.main.Message) => {
  console.log("Received callout update message:", msg);

  const updateMsgJson: cardUpdateJson = JSON.parse(msg.message);

  const cardToUpdate = calloutCardRef.value.find((card: CalloutCard) => card.id === updateMsgJson.callout_id);
  console.log("Card to update:", cardToUpdate);
  if (cardToUpdate !== undefined) {
    cardToUpdate.changeStatus(updateMsgJson.update_request);
  } else {
    console.log("[ERROR] Card id:", updateMsgJson.callout_id, "not found, cannot update.");
  }

});

wails.EventsOn("callout/delete", (msg: models.main.Message) => {
  console.log("[INFO] Received callout delete message:", msg);
  const idx = cards.value.findIndex((card) => card.callout_id === msg.message); 
  if (idx !== -1) {
    cards.value.splice(idx, 1);
  }
  else {
    console.log("[ERROR] No card with id:", msg.message, "found!")
  }
});
</script>

<style scoped>
</style>
