import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import constant from './constants/upload_status'

Vue.config.productionTip = false

Vue.use(axios)
Vue.use(constant)

new Vue({
  render: h => h(App),
}).$mount('#app')
