import { getCurrentInstance, ref, computed } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

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
    console.log('来了', searchForm.value)
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
    loading.value = false
    tableData.value = res.data
    pagination.value = res.pagination
  }

  /**
   * 确认删除单行数据
   * @param {*} row 同table中的scope.row
   */
  const onConfirmDelete = async (row) => {
    const res = deleteApi({ ids: row.article_id })
    console.log(res)
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
    console.log(refTable, refTable.value)
    refTable.value.toggleAllSelection()
  }

  loadData()

  // 根据设备判断搜索表单是否行内表单
  const inlineForm = computed(() => !ctx.$isMobile)

  return {
    searchForm,
    onSearch,

    tableData,
    pagination,
    loading,
    inlineForm,

    onCurrentChange,
    onConfirmDelete,
    onSelectChange
  }
}
