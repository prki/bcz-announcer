<template>
  <div class="row">
    <div class="col">
      <h2>Match callout view</h2>
      <MatchCalloutListing />
      <MatchCalloutCreate />
    </div>
  </div>
</template>

<script lang="ts" setup>
import MatchCalloutListing from '../components/MatchCalloutListing.vue';
import MatchCalloutCreate from '../components/MatchCalloutCreate.vue';
</script>

<style scoped>
</style>
