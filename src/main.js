import Vue from 'vue'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false

import '@/assets/css/style.css'
import '@/assets/js/iconfont.js'

import Icon from '@/components/Icon/Icon.vue'
Vue.component('Icon', Icon)

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
