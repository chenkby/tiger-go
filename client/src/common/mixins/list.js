import { getCurrentInstance, ref, computed } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

export default function (getListApi, deleteApi) {
  const { ctx } = getCurrentInstance()
  const store = useStore()
  const route = useRoute()

  // 搜索model
  const searchForm = ref({
    keyword: '',
    status: ''
  })

  // 提交搜索
  const onSubmitQueryForm = () => {
    loadData()
  }

  // 表格数据
  const tableData = ref([])

  const loading = ref(true)
  const pagination = ref({})
  const currentPath = route.path

  // 请求列表api
  const loadData = async (page) => {
    const queryParams = searchForm.value || {}
    queryParams.page = page || pagination.currentPage
    queryParams.pageSize =
      store.state.table.pageSize[currentPath] || pagination.pageSize || 10
    const res = await getListApi(queryParams)
    loading.value = false
    tableData.value = res.data
    pagination.value = res.pagination
  }

  // 表格操作
  const onConfirmDelete = async (row) => {
    const res = deleteArticle({ ids: row.article_id })
    console.log(res)
  }

  // 分页器事件
  const onCurrentChange = (page) => {
    loadData(page)
  }

  loadData()

  console.log('aa', ctx)
  // 是否行内表单
  const inlineForm = computed(() => !ctx.$isMobile)

  return {
    searchForm,
    onSubmitQueryForm,

    tableData,
    pagination,
    loading,
    inlineForm,
    onCurrentChange,

    onConfirmDelete
  }
}
