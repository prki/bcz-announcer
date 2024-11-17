<template>
  <div class="col">
    <!--<div class="card" style="width: 18rem;">-->
    <div class="card card-font bg-transparent">
      <div class="card-header text-center">
        {{ friendlyGameTitle(gameName) }}
      </div>
      <div class="card-body bg-transparent" :class="status">
        <ul class="list-group list-group-flush text-center">
          <li class="list-group-item bg-transparent card-font">{{ p1Name }}</li>
          <li class="list-group-item bg-transparent card-font">{{ p2Name }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
const props = defineProps(['gameName', 'p1Name', 'p2Name', 'id', 'state']);
console.log(props.state);

const status = ref('default');
const id = props.id;

function changeStatus(newStatus: string) {
  status.value = newStatus;
}

defineExpose({
  changeStatus,
  id
});

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
.card-font {
  font-family: 'Montserrat', sans-serif;
  font-weight: 800;
  font-size: 28px;
  color: white;
}

.default {
  background-color: green;
}
.dq {
  background-color: yellow;
}
</style>
