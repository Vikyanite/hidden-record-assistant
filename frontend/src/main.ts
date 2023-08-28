import {createApp} from 'vue'
import App from './App.vue'

const app = createApp(App)

// ==================== vue-router ====================
import router from './routes/router'
app.use(router.router)

// ==================== element-plus ====================
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
app.use(ElementPlus)


app.mount('#app')
