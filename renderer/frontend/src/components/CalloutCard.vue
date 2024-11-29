<template>
  <div class="col">
    <div class="card card-font" 
         :class="[cardSize, { border: status === 'dq', 'border-warning': status === 'dq', 'border-danger': status === 'stream', 'border-secondary': status === 'stream_next', 'bg-transparent': status === 'default', 'dq': status === 'dq', 'stream': status === 'stream', 'stream-next': status === 'stream_next'}]">
      <img :src="gameNameToLogoPath(gameName)" class="card-img-top" :class="cardSize">
      <div class="card-body bg-transparent" :class="status">
        <ul class="list-group list-group-flush text-center">
          <li class="list-group-item bg-transparent card-font text-truncate">
            {{ getFlagEmoji(p1CountryCode) }} {{ p1Name }}
            <!--<span v-html="getFlagEmoji(p1CountryCode)"></span> {{ p1Name }}-->
          </li>
          <li class="list-group-item bg-transparent card-font border-bottom-0 text-truncate">
            {{ getFlagEmoji(p2CountryCode) }} {{ p2Name }}
             <!--<span v-html="getFlagEmoji(p2CountryCode)"></span> {{ p2Name }}-->
          </li>
          <li class="list-group-item bg-transparent card-font border-bottom-0">
            <span v-if="status === 'dq'" style="color: yellow;">DQ ALERT</span>
            <span v-else-if="status === 'stream'" style="color: red;">STREAM</span>
            <span v-else-if="status === 'stream_next'" style="color: darkorange;">NEXT STREAM</span>
            <span v-else style="opacity: 0;">Dummy</span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps(['gameName', 'p1Name', 'p1CountryCode', 'p2Name', 'p2CountryCode', 'id', 'state', 'cardSize', 'status']);

//const status = ref('default');
console.log("[DEBUG] Callout card status:", props.status);
const id = props.id;

defineExpose({
  //matchStatus,
  id
});

// Function returning image path for a particular logo.
// Simple switch/case since low number of options/static content
function gameNameToLogoPath(gamename: string): string {
  const basePath = "/images/"
  switch(gamename) {
    case "bbcf":
      return basePath + "BBCF_logo.svg";
      //return new URL(basePath + 'BBCF_logo.svg', import.meta.url).href;
    case "strive":
      return basePath + "GGST_logo_resize.svg";
      //return new URL(basePath + 'GGST_logo_resize.svg', import.meta.url).href;
    case "xrd":
      return basePath + "ggxrdrev2_logo.svg";
    case "plusr":
      return basePath + "GGAC+R_logo.svg";
    case "sf6":
      return basePath + "SF6_logo.svg";
    case "tekken":
      return basePath + "Tekken-8-logo_white.svg";
    case "gov":
      return basePath + "govlogo.svg";
    default:
      return "";
  }
}


// Taken from https://dev.to/jorik/country-code-to-flag-emoji-a21
function getFlagEmoji(countryCode: string) {
  const codePoints = countryCode
    .toUpperCase()
    .split('')
    .map(char =>  127397 + char.charCodeAt(0));
  return String.fromCodePoint(...codePoints);

  /*const emoji = String.fromCodePoint(...codePoints);

  return twemoji.parse(emoji, {
    folder: 'svg',
    ext: '.svg'
  });
  */
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
  /*font-family: 'Montserrat', sans-serif;*/
  font-family: 'Montserrat', 'Noto Color Emoji', 'Segoe UI Emoji', sans-serif;
  font-weight: 800;
  /*font-size: 22px;*/
  color: white;
}

.default {
  background-color: green;
}
.dq {
  background-color: transparent;
  animation: fadeBackgroundDq 3s infinite;
}

.stream-next {
  background-color: transparent;
  animation: fadeBackgroundStreamNext 3s infinite;
}

@keyframes fadeBackgroundStreamNext {
  0% {
    background-color: transparent;
  }
  50% {
    background-color: rgba(255, 140, 0, 0.2);
  }
  100% {
    background-color: transparent;
  }
}

@keyframes fadeBackgroundDq {
  0% {
    background-color: transparent;
  }
  50% {
    background-color: rgba(255, 255, 0, 0.2);
  }
  100% {
    background-color: transparent;
  }
}

.stream {
  background-color: transparent;
  animation: fadeBackgroundStream 3s infinite;
}

@keyframes fadeBackgroundStream {
  0% {
    background-color: transparent;
  }
  50% {
    background-color: rgba(255, 0, 0, 0.2);
  }
  100% {
    background-color: transparent;
  }
}

.card-img-top {
  width: 100%;
  height: 15vh;
  object-fit: fill;
}

.card-img-top.cardsize-large {
  height: 12vh;
}

.card-img-top.cardsize-medium {
  height: 11vh;
}

.card-img-top.cardsize-small {
  height: 7vh;
}

.list-group-item {
  overflow: hidden;
}

.cardsize-large {
  font-size: 36px !important;
}

.cardsize-medium {
  font-size: 24px !important;
}

.cardsize-small {
  font-size: 20px !important;
}
</style>
