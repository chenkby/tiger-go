import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import './style/element-variables.scss'
import App from './App.vue'

// router
import router from './router'

// // element ui plus
// // import 'element-plus/packages/theme-chalk/src/base.scss'
// import 'element-plus/lib/theme-chalk/index.css'

const app = createApp(App)
app.use(router)

// 可以通过配置size改变大小，如：size: 'medium'
app.use(ElementPlus, { zIndex: 2000 })

app.mount('#app')
