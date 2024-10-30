import {createApp} from 'vue'
import App from './App.vue'
//import './style.css';
import './assets/css/bootstrap.css'
import router from './router/index'

createApp(App).use(router).mount('#app')
