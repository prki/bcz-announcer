<script lang="ts" setup>
//import HelloWorld from './components/HelloWorld.vue'
//import Main from './views/Main.vue'
import { useRouter } from 'vue-router';
import * as models from '../wailsjs/go/models';
import * as wails from '../wailsjs/runtime/runtime.js';
import Footer from './components/Footer.vue';
import Header from './components/Header.vue';
const router = useRouter();
let currFullscreen = false;

// Preload routes by navigating through them. This is important so that data
// can be sent into a particular view without it having been opened yet =>
// add all possible router views into the cached tree.
// [TODO] There **must** be a better solution than this. Even if this preloading
// is done in a loop, it should not have to be hard coded like this. To find out
// at some point, though this works for now.
router.push("/matchcallout").then(() => {
  router.push("/textdisplay").then(() => {
    router.push("/tableresults").then(() => {
      router.push("/")
    });
  });
});

function handleKeyPress(event: any) {
    if (event.key === 'F11') {
        event.preventDefault(); // Prevent the browser default fullscreen action
        toggleFullscreen(currFullscreen);
        currFullscreen = !currFullscreen;
    }
}

function toggleFullscreen(currFullscreen: boolean) {
  if (!currFullscreen) {
    wails.WindowFullscreen(); // Switch this with `WindowUnfullscreen()` if already fullscreen
  } else {
    wails.WindowUnfullscreen();
  }
}

window.addEventListener('keydown', handleKeyPress);

wails.EventsOn("control/change_view", (msg: models.main.Message) => {
  console.log("Received change view request:", msg);

  console.log("Router:", router);
  console.log("msg.message:", msg.message);
  router.push(msg.message);
});
</script>

<template>
  <!--<img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
  <div class="container-fluid px-4">
    <!-- Header row -->
    <Header />

    <!-- Content row -->
    <div class="row mt-4 mb-4 gx-4">
      <div class="col-8">
        <div class="row">
          <div class="col">
            <div class="border content-row-size">
              <router-view v-slot="{ Component }">
              <KeepAlive>
                <component :is="Component" />
              </KeepAlive>
              </router-view>
            </div>
          </div>
        </div>
        <!-- Footer-->
        <Footer />
      </div>
      <div class="col-4 d-flex">
        <div class="border flex-grow-1">
          <img class="img-fluid" src="./assets/images/sidepanel_logo.png" style="object-fit: fill;">
        </div>
      </div>
    </div>

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

.box-bg-blue {
  background-color: blue;
  color: white;
}

.box-bg-yellow {
  background-color: yellow;
  color: black;
}

/* [TODO] Consider percentage + direct margin/padding from top/bottom*/
.content-row-size {
  height: 73vh;
}

.header-row-size {
  height: 10vh;
}
</style>
