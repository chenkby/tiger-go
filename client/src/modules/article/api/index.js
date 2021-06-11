import request from '@/utils/request'

export const list = function (queryParams) {
  return request({
    url: '/article/list',
    method: 'get',
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
