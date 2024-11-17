<script lang="ts" setup>
//import HelloWorld from './components/HelloWorld.vue'
//import Main from './views/Main.vue'
import { useRouter } from 'vue-router';
import * as models from '../wailsjs/go/models';
import * as wails from '../wailsjs/runtime/runtime.js';
import Footer from './components/Footer.vue';
import Header from './components/Header.vue';
const router = useRouter();

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
    <!--<div class="row header-row-size mt-4 gx-4">
      <div class="col-8">
        <div class="border bg-light p-3 pb-4 pt-4">
          <div class="row">
            <div class="col-6 offset-md-1 bg-light border">
              <p>Header text placeholder</p>
            </div>
            <div class="col-3 offset-md-1 bg-light border">
              <p>Header image placeholder</p>
            </div>
          </div>
        </div>
      </div>
      <div class="col-4">
        <div class="p-3 border bg-light pb-4 pt-4">
          <p>BCZ logo</p>
        </div>
      </div>
    </div>
    -->
    <!--<Main />-->

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
          <p>Side panel</p>
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
