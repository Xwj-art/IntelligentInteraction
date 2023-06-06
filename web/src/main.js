import Vue from 'vue';
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';
import axios from 'axios'
import './assets/css/style.css'

axios.defaults.baseURL = "http://localhost:3030/api"
Vue.prototype.$http = axios
Vue.use(ElementUI);

Vue.config.productionTip = false

new Vue({
  router,
  el: '#app',
  render: h => h(App)
});