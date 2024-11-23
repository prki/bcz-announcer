<template>
  <div class="col">
    <div class="card card-font bg-transparent" :class="[cardSize, { border: status === 'dq', 'border-warning': status === 'dq', 'border-danger': status === 'stream' }]">
      <img :src="gameNameToLogoPath(gameName)" class="card-img-top" :class="cardSize">
      <div class="card-body bg-transparent" :class="status">
        <ul class="list-group list-group-flush text-center">
          <li class="list-group-item bg-transparent card-font">{{ p1Name }}</li>
          <li class="list-group-item bg-transparent card-font border-bottom-0">{{ p2Name }}</li>
          <li class="list-group-item bg-transparent card-font border-bottom-0" v-show="status === 'dq'" style="color: yellow;">DQ ALERT</li>
          <li class="list-group-item bg-transparent card-font border-bottom-0" v-show="status === 'stream'" style="color: red;">STREAM</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
const props = defineProps(['gameName', 'p1Name', 'p2Name', 'id', 'state', 'cardSize']);
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

// Function returning image path for a particular logo.
// Simple switch/case since low number of options/static content
function gameNameToLogoPath(gamename: string): string {
  const basePath = "src/assets/images/"
  switch(gamename) {
    case "bbcf":
      return basePath + "BBCF_logo.svg";
    case "strive":
      return basePath + "GGST_logo.svg";
    case "xrd":
      return basePath + "ggxrdrev2_logo.svg";
    case "plusr":
      return basePath + "GGAC+R_logo.svg";
    case "sf6":
      return basePath + "SF6_logo.svg";
    case "tekken":
      return basePath + "Tekken-8-logo_white.svg";
    default:
      return "";
  }

}

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
  /*font-size: 22px;*/
  color: white;
}

.default {
  background-color: green;
}
.dq {
  background-color: yellow;
}

.card-img-top {
  width: 100%;
  height: 15vh;
  object-fit: fill;
}

.card-img-top.cardsize-large {
  height: 30vh;
}

.card-img-top.cardsize-medium {
  height: 20vh;
}

.card-img-top.cardsize-small {
  height: 10vh;
}

.list-group-item {
  overflow: hidden;
}

.cardsize-large {
  font-size: 40px !important;
}

.cardsize-medium {
  font-size: 20px !important;
}

.cardsize-small {
  font-size: 10px !important;
}
</style>
