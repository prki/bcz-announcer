<template>
  <div class="col">
    <!--<div class="card" style="width: 18rem;">-->
    <div class="card">
      <div class="card-header">
        <h5>{{ gameName }}</h5>
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
</script>

<style scoped>
</style>
