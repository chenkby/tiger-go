import { createApp } from 'vue'
import store from './store'
import router from './router'
import { detect, deepMerge } from './utils/helper'

// elementUI
import ElementPlus from 'element-plus'
import './style/element-variables.scss'
// import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale/lang/zh-cn'

// async-validator i18N
import AsyncValidator from 'async-validator'
import validatorMessages from '@/common/validator/zh-CN.validator_messages'
deepMerge(AsyncValidator.messages, validatorMessages)

import App from './App.vue'

const app = createApp(App)
app.use(router)

// 可以通过配置size改变大小，如：size: 'medium'
app.use(ElementPlus, { zIndex: 3000, locale })
app.config.globalProperties.$device = detect().device.type || 'pc' // 设备类型
app.config.globalProperties.$isMobile = detect().device.type === 'mobile' // 是否手机端
app.use(store)
app.mount('#app')
