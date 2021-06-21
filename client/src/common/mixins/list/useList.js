/**
 * 列表页逻辑
 */
import {
  getCurrentInstance,
  ref,
  computed,
  provide,
  onMounted,
  onUnmounted,
  inject
} from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import emitter from 'tiny-emitter/instance'

const useSearch = () => {
  // 搜索model
  const searchForm = ref({
    keyword: ''
  })
  // 提交搜索
  const onSearch = () => {
    loadData()
  }
  return {
    searchForm
  }
}

const useLoad = () => {}

/**
 * 操作
 * @param {*} ctx
 * @returns
 */
const useDataOperate = (ctx, dialogMode, deleteApi, loadData) => {
  // 点击行编辑时触发
  const onUpdate = (row) => {
    if (dialogMode) {
      // @todo 打开对话框
    } else {
      ctx.$router.push({
        append: true,
        path: 'update',
        query: { id: row['article_id'] }
      })
    }
  }

  /**
   * 确认删除单行数据
   * @param {*} row 同table中的scope.row
   */
  const onDelete = async (params) => {
    const { err_code, data } = await deleteApi({ ids: params.ids })
    if (err_code === 0) {
      if (data > 0) {
        ctx.$message.success(`成功删除${data}条数据`)
        loadData(1)
      } else {
        ctx.$message.warning('没有删除数据')
      }
    }
  }
  return {
    onUpdate,
    onDelete
  }
}

const useToolbar = (loadData) => {
  // 每页条数改变时重新加载数据
  const onSizeChange = () => {
    loadData(1)
  }
  return {
    onSizeChange
  }
}

/**
 * 表格操作
 * @param {*} tableData
 * @returns
 */
const useTable = (tableData) => {
  // 当选择项发生变化时会触发该事件
  const onSelectionChange = (selection) => {
    emitter.emit('selectionChange', selection, tableData.value.length)
  }
  return {
    onSelectionChange
  }
}
/**
 *
 * @param {Promise} getListApi 列表api
 * @param {Promise} deleteApi 删除api
 * @param {String} primaryKey 表主键名
 * @param {Boolean} dialogMode 是否使用对话框模式
 * @returns
 */
export default function (getListApi, deleteApi, primaryKey, dialogMode) {
  const { ctx } = getCurrentInstance()
  const store = useStore()
  const route = useRoute()
  // 引用el-table
  const refTable = ref(null)

  provide('primaryKey', primaryKey) // 主键名
  provide('dialogMode', dialogMode) // 是否对话框模式
  provide('refTable', refTable) // el-table的引用

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
      store.state.table.pageSize[currentPath] ||
      pagination.value.pageSize ||
      import.meta.env.VITE_DEFAULT_PAGESIZE
    const res = await getListApi(queryParams)
    emitter.emit('loadDataComplate')
    loading.value = false
    tableData.value = res.data
    pagination.value = res.pagination
  }

  /**
   * 当前页改变时触发
   * @param {*} page
   */
  const onCurrentChange = (page) => {
    loadData(page)
  }

  loadData()

  return {
    ...useTable(tableData),
    refTable,
    searchForm,
    onSearch,

    tableData,
    pagination,
    loading,

    onCurrentChange,
    ...useToolbar(loadData),
    ...useDataOperate(ctx, dialogMode, deleteApi, loadData)
  }
}
