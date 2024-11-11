<template>
  <div class="row"> 
    <div class="col">
      <h1>I am match callout view</h1>
    </div>
  </div>
  <div class="row">
    <div class="col">
      <h2>Player 1: {{ state.A.p1Name }}</h2>
    </div>
    <div class="col">
      <h2>Player 2: {{ state.A.p2Name }}</h2>
    </div>
  </div>
  <div class="row">
    <div class="col">
      <div class="row row-cols-md-4">
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>

        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
        <CalloutCard gameName="Guilty Gear Xrd" p1Name="Pida" p2Name="Kidiot"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import CalloutCard from '../components/CalloutCard.vue';
import { reactive } from 'vue';
import * as wails from '../../wailsjs/runtime/runtime.js';
import * as models from '../../wailsjs/go/models';

const state = reactive({
  A: {
    p1Name: '',
    p2Name: ''
  }
});

wails.EventsOn("callout/update", (msg: models.main.Message) => {
  console.log("Received callout update message:", msg);
  const msgJson = JSON.parse(msg.message);  // This could be given a type since we know what the message is.
  console.log("Message as obj:", msgJson);

  state.A.p1Name = msgJson.p1_name;
  state.A.p2Name = msgJson.p2_name;
});
</script>

<style scoped>
</style>
