import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Axios from 'axios'

Vue.config.productionTip = false

Vue.prototype.$http = Axios.create({
  baseURL: 'http://localhost:2233/',
  headers: {
    'Content-Type': 'application/json'
  }
})


// 引入自定义文件
import '@/assets/css/style.css'
import '@/assets/js/iconfont.js'
// 引入字体文件
import Icon from '@/components/Icon/Icon.vue'
Vue.component('Icon', Icon)
// 引入UI框架
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
