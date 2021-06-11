<template>
  <div class="tiger-table-footer">
    <div class="grid-footer__left">
      <slot name="left">
        <div class="grid-footer__checkbox">
          <el-checkbox></el-checkbox>
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
        <el-pagination background :page-size="100" :layout="paginationLayout" :total="1000">
        </el-pagination>
      </slot>
    </div>
  </div>
</template>

<script>
import { computed } from '@vue/runtime-core'
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
     * 是否显示分页
     */
    pagination: {
      type: Boolean,
      default: true
    }
  },
  setup(props, context) {
  },
  computed: {
    /**
 * 分页器布局
 */
    paginationLayout() {
      if (this.$device === 'mobile') {
        return 'total, prev, next, jumper'
      }
      return 'total,sizes, prev, pager, next, jumper'
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
</style>
