import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import Auth from './components/Auth.vue'
import Reg from './components/Reg.vue'
// import Main from './components/Header.vue'

Vue.use(VueResource)
Vue.config.productionTip = false
Vue.use(VueRouter)

const router=new VueRouter({
  mode:'history',
  routes:[
    // {path:'/',component:Main},
    {path:'/reg',component:Reg},
    {path:'/',component:Auth}
  ],
  linkActiveClass:'active'
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
