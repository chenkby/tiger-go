/**
 * 列表页逻辑
 */
import { getCurrentInstance, ref, computed, provide } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import emitter from 'tiny-emitter/instance'

export default function (getListApi, deleteApi) {
  const { ctx } = getCurrentInstance()
  const store = useStore()
  const route = useRoute()

  // 搜索model
  const searchForm = ref({
    keyword: ''
  })

  // 提交搜索
  const onSearch = () => {
    loadData()
  }

  const tableData = ref([]) // 表格数据
  const loading = ref(true) // 是否加载数据中
  const pagination = ref({}) // 分页配置
  const currentPath = route.path // 当前路由，用于配置分页器里面的每页数据数量

  /**
   * 请求列表数据
   * @param {*} page 要请求第几页
   */
  const loadData = async (page) => {
    const queryParams = searchForm.value || {}
    queryParams.page = page || pagination.value.currentPage
    queryParams.pageSize =
      store.state.table.pageSize[currentPath] || pagination.value.pageSize || 10
    const res = await getListApi(queryParams)
    emitter.emit('loadDataComplate')
    loading.value = false
    tableData.value = res.data
    pagination.value = res.pagination
  }

  /**
   * 确认删除单行数据
   * @param {*} row 同table中的scope.row
   */
  const onDelete = async (ids) => {
    const { err_code, data } = await deleteApi({ ids: ids })
    if (err_code === 0) {
      if (data > 0) {
        ctx.$message.success(`成功删除${data}条数据`)
        loadData(1)
      } else {
        ctx.$message.warning('没有删除数据')
      }
    }
  }

  const onUpdate = (row) => {
    ctx.$router.push({
      append: true,
      path: 'update',
      query: { id: row['article_id'] }
    })
  }

  /**
   * 当前页改变时触发
   * @param {*} page
   */
  const onCurrentChange = (page) => {
    loadData(page)
  }
  const refTable = ref(null)
  const onSelectChange = () => {
    refTable.value.toggleAllSelection()
  }

  loadData()

  return {
    searchForm,
    onSearch,

    tableData,
    pagination,
    loading,

    onCurrentChange,
    onUpdate,
    onDelete,
    onSelectChange
  }
}
