<script lang="ts" setup>
//import HelloWorld from './components/HelloWorld.vue'
//import Main from './views/Main.vue'
import { useRouter } from 'vue-router';
import * as models from '../wailsjs/go/models';
import * as wails from '../wailsjs/runtime/runtime.js';
const router = useRouter();

wails.EventsOn("control/change_view", (msg: models.main.Message) => {
  console.log("Received change view request:", msg);

  console.log("Router:", router);
  console.log("msg.message:", msg.message);
  router.push(msg.message);
});
</script>

<template>
  <!--<img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
  <div class="container">
    <div class="row">
      <h1>BCZ 2024 Callout Renderer</h1>
    </div>
    <!--<Main />-->
    <RouterView />
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
