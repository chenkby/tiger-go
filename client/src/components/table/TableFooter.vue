<template>
  <div class="tiger-table-footer">
    <div class="grid-footer__left">
      <slot name="left">
        <div class="grid-footer__checkbox">
          <el-checkbox @change="onSelectChange"></el-checkbox>
        </div>
        <div class="grid-footer__tools">
          <el-button type="danger" disabled>删除</el-button>
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
import { ref, computed } from 'vue'
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
     * 分页器
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
  emits: ['selectChange'],
  setup(props, { emit }) {
    //const refTable = ref(null)
    const onSelectChange = () => {
      //console.log('测试', refTable)
      //refTable.value.toggleAllSelection()
      emit('selectChange')
    }
    return {
      //refTable,
      onSelectChange
    }
  },
  computed: {
    /**
     * 分页布局
     */
    paginationLayout() {
      if (this.$device === 'mobile') {
        return 'total, prev, next, jumper'
      }
      return 'total,sizes, prev, pager, next, jumper'
    }
  },
  methods: {
    /**
     * 页码改变时触发
     */
    onSizeChange(pageSize) {
      this.$store.dispatch('setPageSize', {
        path: this.$route.path,
        pageSize: pageSize
      })
      this.$emit('size-change', pageSize)
    },
    /**
     * 当前页改变时触发
     */
    onCurrentChange(page) {
      this.$emit('current-change', page)
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
