<template>
  <div class="col">
    <!--<div class="card" style="width: 18rem;">-->
    <div class="card">
      <div class="card-header">
        <h5>{{ friendlyGameTitle(gameName) }}</h5>
      </div>
      <div class="card-body">
        <ul class="list-group list-group-flush">
          <li class="list-group-item">{{ p1Name }}</li>
          <li class="list-group-item">{{ p2Name }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps(['gameName', 'p1Name', 'p2Name', 'id']);
console.log(props.id);

// Simple switch/case since low number of options/static content
function friendlyGameTitle(gamename: string): string {
  switch(gamename) {
    case "bbcf":
      return "Blazblue Central Fiction";
    case "strive":
      return "Guilty Gear Strive";
    case "xrd":
      return "Guilty Gear Xrd";
    case "plusr":
      return "Guilty Gear ACPR";
    case "sf6":
      return "Street Fighter 6";
    case "tekken":
      return "Tekken 8";
    default:
      return "";
  }
}
</script>

<style scoped>
</style>
