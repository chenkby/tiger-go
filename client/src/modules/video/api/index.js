import request from '@/utils/request'

export const listApi = function (queryParams) {
  return request({
    url: '/video/list',
    method: 'get',
    params: queryParams
  })
}

export const infoApi = function (queryParams) {
  return request({
    url: '/video/info',
    method: 'get',
    params: queryParams
  })
}

export const createApi = (data, queryParams) => {
  return request({
    url: '/video/create',
    method: 'post',
    data: data,
    params: queryParams
  })
}

export const updateApi = (data, queryParams) => {
  return request({
    url: '/video/update',
    method: 'post',
    data: data,
    params: queryParams
  })
}

export const deleteApi = (data, queryParams) => {
  return request({
    url: '/video/delete',
    method: 'post',
    data: data,
    params: queryParams
  })
}

export const toggleApi = (data, queryParams) => {
  return request({
    url: '/video/toggle',
    method: 'post',
    data: data,
    params: queryParams
  })
}

// or
// export default {
//   list(queryParams) {
//     return request.get('/article/list', {
//       params: queryParams
//     })
//   },
//   // @Summary 添加文章
//   // @Security ApiKeyAuth
//   // @accept application/json
//   // @Produce application/json
//   // @Param data body {authorityId uint} true "添加文章"
//   // @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
//   // @Router /article/create [post]
//   create(data, queryParams) {
//     return request.post('/article/create', data, {
//       params: queryParams
//     })
//   },
//   update(data, queryParams) {
//     return request.post('/article/update', data, {
//       params: queryParams
//     })
//   },
//   view(queryParams) {
//     return request.get('/article/view', {
//       params: queryParams
//     })
//   },
//   delete(data, queryParams) {
//     return request.post('/article/delete', data, {
//       params: queryParams
//     })
//   },
//   toggle(data, queryParams) {
//     return request.post('/article/toggle', data, {
//       params: queryParams
//     })
//   }
// }
