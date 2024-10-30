import { createRouter, createWebHashHistory } from "vue-router";
import Main from '../views/Main.vue';
import MatchCallout from '../views/MatchCallout.vue';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "main",
      component: Main
    },
    {
      path: "/matchcallout",
      name: "matchcallout",
      component: MatchCallout
    }
  ]
});

export default router
