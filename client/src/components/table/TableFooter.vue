<template>
  <div class="tiger-table-footer">
    <div class="grid-footer__left">
      <slot name="left">
        <div class="grid-footer__checkbox">
          <el-checkbox @change="onSelectChange" v-model="checked" :indeterminate="indeterminate"></el-checkbox>
        </div>
        <div class="grid-footer__tools">
          <el-button :disabled="selection.length <= 0 " @click="onDelete">{{deleteText}}</el-button>
          <div class="item">
            <el-dropdown disabled>
              <el-button>
                更多菜单<i class="el-icon-arrow-down el-icon--right"></i>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>黄金糕</el-dropdown-item>
                  <el-dropdown-item>狮子头</el-dropdown-item>
                  <el-dropdown-item>螺蛳粉</el-dropdown-item>
                  <el-dropdown-item>双皮奶</el-dropdown-item>
                  <el-dropdown-item>蚵仔煎</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </slot>
    </div>
    <div class="grid-footer__right">
      <slot name="right">
        <el-pagination background :page-size="pagination.pageSize" :layout="paginationLayout" :total="pagination.totalCount" @size-change="onSizeChange" @current-change="onCurrentChange">
        </el-pagination>
      </slot>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, onUnmounted, getCurrentInstance, inject } from 'vue'
import emitter from 'tiny-emitter/instance'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
export default {
  props: {
    /**
     * 是否显示工具栏
     */
    toolbar: {
      type: Boolean,
      default: true
    },
    /**
     * 分页器数据
     */
    pagination: {
      type: Object,
      default() {
        return {
          currentPage: 1,
          pageCount: 0,
          pageSize: 10,
          totalCount: 0
        }
      }
    }
  },
  setup(props, { emit }) {
    const { ctx } = getCurrentInstance()
    const store = useStore()
    const route = useRoute()

    const primaryKey = inject("primaryKey")

    const selection = ref([]) // 已选中数据项
    const checked = ref(false)  // checkbox是否选中
    const indeterminate = ref(false)  // checkbox是否显示不确定样式
    const table = inject('refTable')  // 上级页面表格的引用
    const selectionIds = ref([])  // 已选中数据主键值


    // 监听selectionChange，获得表格选择项
    onMounted(() => {
      // tableSelection:表格选中项
      // dataLength:当前页总共有多少条数据
      emitter.on('selectionChange', (tableSelection, dataLength) => {
        selection.value = tableSelection
        checked.value = tableSelection.length > 0
        indeterminate.value = tableSelection.length > 0 ? tableSelection.length < dataLength : false
        tableSelection.forEach((item) => {
          selectionIds.value.push(item[primaryKey])
        })
      })
    })
    onUnmounted(() => {
      emitter.off('selectionChange')
    })

    // 删除按钮的文本
    const deleteText = computed(() => {
      return selection.value.length > 0 ? '删除(' + selection.value.length + ')' : '删除'
    })

    // 分页布局
    const paginationLayout = computed(() => ctx.$isMobile ? 'total, prev, next, jumper' : 'total,sizes, prev, pager, next, jumper')

    /**
     * 工具栏全选/取消全选
     */
    const onSelectChange = () => {
      table.value.toggleAllSelection()
    }
    /**
     * 每页条数改变时触发
     */
    const onSizeChange = (pageSize) => {
      store.dispatch('setPageSize', {
        path: route.path,
        pageSize: pageSize
      })
      emit('size-change', pageSize)
    }
    /**
     * 跳转页面时触发
     */
    const onCurrentChange = (page) => {
      emit('current-change', page)
    }

    /**
     * 点击删除按钮触发
     */
    const onDelete = () => {
      ctx.$confirm(`确定要删除选中的${selection.value.length}条数据吗?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        buttonSize: 'medium'
      }).then(() => {
        emit('delete', { ids: selectionIds.value })
      }).catch(() => {

      })
    }
    return {
      table,
      selection,
      checked,
      indeterminate,
      deleteText,
      paginationLayout,
      onSelectChange,
      onSizeChange,
      onCurrentChange,
      onDelete
    }
  }
}
</script>

<style lang="scss">
.tiger-table-footer {
  padding: 15px 0;

  display: flex;
  align-items: center;
  justify-content: space-between;
  .grid-footer__left {
    display: flex;
    align-items: center;
    .grid-footer__checkbox {
      padding: 0 14px;
      text-align: center;
    }
    .grid-footer__tools {
      padding-left: 10px;
      display: flex;
      align-items: center;
      .item {
        margin-left: 10px;
      }
    }
  }

  .el-pagination {
    font-weight: normal;
  }
}

// 手机端
body.device-mobile {
  .tiger-table-footer {
    display: block;
    .grid-footer__left {
    }
    .grid-footer__tools {
    }
    .grid-footer__right {
      text-align: center;
    }
    .el-pagination {
      border-top: 1px solid #efefef;
      padding-top: 20px;
      margin-top: 20px;
    }
  }
}
</style>
