import {createApp} from 'vue'
import App from './App.vue'
import './assets/css/bootstrap.css'
import './style.css';
import router from './router/index'

createApp(App).use(router).mount('#app')
