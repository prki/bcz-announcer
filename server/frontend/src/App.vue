<script lang="ts" setup>
//import HelloWorld from './components/HelloWorld.vue'
import Main from './views/Main.vue'
import NavMenu from './components/NavMenu.vue';
import RendererControlMenu from './components/RendererControlMenu.vue';
</script>

<template>
  <!--<img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
  <div class="container" id="main-container">
    <div class="row">
      <h1 class="text-center">BCZ 2024 Callout Admin Console</h1>
    </div>
    <hr />

    <RendererControlMenu />
    <hr />

    <NavMenu />
    <router-view v-slot:="{ Component }">
      <KeepAlive>
        <component :is="Component" />
      </KeepAlive>
    </router-view>
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
