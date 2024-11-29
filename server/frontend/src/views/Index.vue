<template>
  <div class="row">
    <div class="col">
      <h1>BCZ callout server usage guide</h1>
      <hr />
      <div class="border">
        <p class="p-2">This application is used to project relevant information to Buttonczech 2024 entrants.</p>
        <p class="p-2">The application communicates with a <strong>renderer</strong> - accompanying application which will be running on the projector. Both applications should be running. If either crashes, it is recommended to restart both of them to create a connection between the two.</p>
        <p class="p-2">The <strong>renderer</strong> has a static layout with an area defined for various <strong>views</strong> - these views are used for providing various information to entrants.</p>
        <p class="p-2">Multiple views are provided:</p>
        <ul>
          <li><strong>Banner view</strong> - General banner containing image for BCZ 2024. To be displayed when nothing of relevance is happening</li>
          <li><strong>Match callout view</strong> - View for calling out matches - creates callout cards showing players to play. Has capabilities to call out stream games/DQ alerts.</li>
          <li><strong>Text Display view</strong> - A simple view letting the operator send a text message into the renderer. To be used when a particular announcement is to be provided.</li>
          <li><strong>Table Results view</strong> - View showing a table containing GGOB results.</li>
        </ul>
        <p class="p-2">Press <strong>F11</strong> on renderer to make it go fullscreen.</p>
        <p class="p-2">It is recommended to first create content for a particular view and then change to the view - not before.</p>
        <p class="p-2">In case of any questions/technical issues, reach out to <strong>prk`</strong></p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style scoped>
</style>
