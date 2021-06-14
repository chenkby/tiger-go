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
  timeout: 3000
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
    console.log('请求发生了错误', error)
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
    console.log('res', response)
    // 如果设置了responseType，说明不是json，有可能是blob等，这种情况下不做处理
    if (response.config.responseType) {
      return response
    }
    // closeLoading()

    // if (response.headers['new-token']) {
    //   store.commit('user/setToken', response.headers['new-token'])
    // }
    return new Promise((resolve, reject) => {
      if (response.data.hasOwnProperty('err_code')) {
        if (response.data.err_code === 0) {
          return resolve(response.data)
        } else {
          //errorHandle(res.data.err_code, res.data)
          Message({
            showClose: true,
            message: response.data.err_msg,
            type: 'error'
          })
          return resolve(response.data)
        }
      } else {
        Message({
          showClose: true,
          message: '请求发生了错误',
          type: 'error'
        })
        resolve(response.data)
      }
    })
  },
  (error) => {
    const { response } = error
    // 请求已发出，但是不在2xx的范围
    if (response) {
    } else {
      // 断网或者超时
      // @todo 重试N次
      Message({
        showClose: true,
        message: '请求失败，请检查你的网络状态',
        type: 'error'
      })
    }
    return error
  }
)

export default request
