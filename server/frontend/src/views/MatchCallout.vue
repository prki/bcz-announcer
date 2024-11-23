<template>
  <div class="row">
    <div class="col-5">
      <!--<h2>Match callout view</h2>-->
      <MatchCalloutCreate/>
    </div>
    <div class="col-7">
      <MatchCalloutListing/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import MatchCalloutListing from '../components/MatchCalloutListing.vue';
import MatchCalloutCreate from '../components/MatchCalloutCreate.vue';
</script>

<style scoped>
</style>
