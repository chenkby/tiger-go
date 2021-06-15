import { createApp } from 'vue'
import store from './store'
// router
import ElementPlus from 'element-plus'
import './style/element-variables.scss'
// import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale/lang/zh-cn'

import router from './router'

import App from './App.vue'

import { detect } from './utils/helper'

// // element ui plus
// // import 'element-plus/packages/theme-chalk/src/base.scss'
// import 'element-plus/lib/theme-chalk/index.css'

const app = createApp(App)
app.use(router)

// 可以通过配置size改变大小，如：size: 'medium'
app.use(ElementPlus, { zIndex: 2000, locale })
app.config.globalProperties.$device = detect().device.type || 'pc' // 设备类型
app.config.globalProperties.$isMobile = detect().device.type === 'mobile' // 是否手机端
app.use(store)
app.mount('#app')
