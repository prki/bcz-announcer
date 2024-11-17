<template>
  <div class="row">
    <div class="col">
      <h2 class="text-center mb-2">God Gamer of Buttonczech 2024</h2>
      <table class="table">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Player name</th>
            <th scope="col">Points</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(playerScore, idx) in state" :key="idx">
            <th scope="row">{{ idx + 1 }}</th>
            <td>{{ playerScore.name }}</td>
            <td>{{ playerScore.points }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
type PlayerScore = {
  name: string,
  points: number
};

const state = ref<PlayerScore[]>([]);

state.value.push({name: "Pida", points: 10});
state.value.push({name: "Arin", points: 9});
state.value.push({name: "prk", points: 8});
state.value.push({name: "Motionman", points: 7});
</script>

<style scoped>
</style>
