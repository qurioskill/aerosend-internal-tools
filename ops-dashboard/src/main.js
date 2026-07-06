import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import QueueHealth from './views/QueueHealth.vue'
import IpPoolHealth from './views/IpPoolHealth.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/queues' },
    { path: '/queues', component: QueueHealth },
    { path: '/ip-pools', component: IpPoolHealth }
  ]
})

createApp(App).use(router).mount('#app')
