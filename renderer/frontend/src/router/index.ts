import { createRouter, createWebHashHistory } from "vue-router";
import Main from '../views/Main.vue';
import MatchCallout from '../views/MatchCallout.vue';
import TextDisplay from '../views/TextDisplay.vue';

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
    },
    {
      path: "/textdisplay",
      name: "textdisplay",
      component: TextDisplay
    }
  ]
});

export default router
