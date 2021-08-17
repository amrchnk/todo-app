import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import Main from './components/Main.vue'
import Reg from './components/Reg.vue'

Vue.use(VueResource)
Vue.config.productionTip = false
Vue.use(VueRouter)

const router=new VueRouter({
  mode:'history',
  routes:[
    {path:'/reg',component:Reg},
    {path:'/',component:Main}
  ],
  linkActiveClass:'active'
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
