import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import App from './App.vue'

// router
import router from './router'

// element ui plus
// import 'element-plus/packages/theme-chalk/src/base.scss'
import 'element-plus/lib/theme-chalk/index.css'

const app = createApp(App)
app.use(router)

app.use(ElementPlus, { size: 'medium', zIndex: 2000 })

app.mount('#app')
