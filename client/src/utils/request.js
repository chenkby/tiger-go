// http request
import axios from 'axios' // 引入axios
import { ElMessage as Message } from 'element-plus'

// let acitveAxios = 0
// let timer
// const showLoading = () => {
//   acitveAxios++
//   if (timer) {
//     clearTimeout(timer)
//   }
//   timer = setTimeout(() => {
//     if (acitveAxios > 0) {
//       context.$bus.emit('showLoading')
//     }
//   }, 400)
// }

// const closeLoading = () => {
//   acitveAxios--
//   if (acitveAxios <= 0) {
//     clearTimeout(timer)
//     context.$bus.emit('closeLoading')
//   }
// }

const request = axios.create({
  baseURL: import.meta.env.VITE_BASE_API_URL,
  timeout: 5000
})
// http request 拦截器
request.interceptors.request.use(
  (config) => {
    // if (!config.donNotShowLoading) {
    //   showLoading()
    // }
    // const token = store.getters['user/token']
    // const user = store.getters['user/userInfo']
    config.data = JSON.stringify(config.data)
    config.headers = {
      'Content-Type': 'application/json'
      // 'x-token': token,
      // 'x-user-id': user.ID
    }
    return config
  },
  (error) => {
    // closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)
// http response 拦截器
request.interceptors.response.use(
  (response) => {
    // closeLoading()

    // if (response.headers['new-token']) {
    //   store.commit('user/setToken', response.headers['new-token'])
    // }
    if (response.data.err_code === 0) {
      if (response.data.data?.needInit) {
        Message({
          type: 'info',
          message: '您是第一次使用，请初始化'
        })
        // store.commit('user/NeedInit')
        // router.push({ name: 'init' })
      }
    }
    if (response.data.err_code === 0) {
      return response.data
    }
    Message({
      showClose: true,
      message: response.data.msg || decodeURI(response.headers.msg),
      type: response.headers.msgtype || 'error'
    })
    // if (response.data.data && response.data.data.reload) {
    //   store.commit('user/LoginOut')
    // }
    return response.data.msg ? response.data : response
  },
  (error) => {
    // closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

export default request
