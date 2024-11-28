import { createRouter, createWebHashHistory } from "vue-router";
import Main from '../views/Main.vue';
import MatchCallout from '../views/MatchCallout.vue';
import TextDisplay from '../views/TextDisplay.vue';
import TableResults from '../views/TableResults.vue';
import DummyView from '../views/DummyView.vue';

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
    },
    {
      path: "/tableresults",
      name: "tableresults",
      component: TableResults
    },
    {
      path: "/dummyview",
      name: "dummyview",
      component: DummyView
    }
  ]
});

export default router
