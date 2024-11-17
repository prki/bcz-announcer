import { createRouter, createWebHashHistory } from "vue-router";
//import Main from '../views/Main.vue';
import Index from '../views/Index.vue';
import MatchCallout from '../views/MatchCallout.vue';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "index",
      component: Index
    },
    {
      path: "/callout",
      name: "callout",
      component: MatchCallout
    }
    /*,
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
    */
  ]
});

export default router
