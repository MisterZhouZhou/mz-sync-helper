import {createApp} from 'vue'
import App from './App.vue'
import router from './routes/router'
import './initializers'

const app = createApp(App)
app.use(router)
app.mount('#app')
