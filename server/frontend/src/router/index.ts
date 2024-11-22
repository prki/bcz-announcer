import { createRouter, createWebHashHistory } from "vue-router";
//import Main from '../views/Main.vue';
import Index from '../views/Index.vue';
import MatchCallout from '../views/MatchCallout.vue';
import Footer from '../views/Footer.vue';
import Header from '../views/Header.vue';
import TextDisplay from '../views/TextDisplay.vue';
import TableResults from '../views/TableResults.vue';

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
    },
    {
      path: "/footer",
      name: "footer",
      component: Footer
    },
    {
      path: "/header",
      name: "header",
      component: Header
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
    }
  ]
});

export default router
