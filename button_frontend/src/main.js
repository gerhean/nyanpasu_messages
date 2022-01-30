import { createApp } from 'vue'

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

import App from './App.vue'
import store from './store/store'

const app = createApp(App)
app.use(store)
app.mount('#app')
