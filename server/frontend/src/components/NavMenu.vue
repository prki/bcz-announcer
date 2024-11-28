<template>
  <div class="row mb-4">
    <div class="col justify-content-center">
      <ul class="nav nav-tabs">
        <li class="nav-item">
          <RouterLink class="nav-link" to="/">Guide</RouterLink>
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
        <li class="nav-item">
          <RouterLink class="nav-link" to="/textdisplay">TextDisplay</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" to="/tableresults">TableResults</RouterLink>
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
