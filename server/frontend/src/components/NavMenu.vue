<template>
  <div class="row mb-4">
    <div class="col">
      <ul class="nav nav-tabs">
        <li class="nav-item">
          <RouterLink class="nav-link" to="/">Index</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" to="/callout">Callouts</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" to="/footer">Footer</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" to="/header">Header</RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style scoped>
.nav-link {
  color: white;
}

.nav-link .active {
  color: white;
}
</style>
