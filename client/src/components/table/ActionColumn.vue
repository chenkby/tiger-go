<template>
  <el-table-column fixed="right" label="操作" :width="width" class-name="action-column">
    <template #default="scope">
      <span @click="onUpdate(scope.row)">编辑</span>
      <el-divider direction="vertical"></el-divider>
      <el-popconfirm placement="top-start" @confirm="onDelete(scope.row)" confirmButtonText='删除' confirmButtonType="danger" icon="el-icon-info" title="确定要删除这条数据吗？">
        <template #reference>
          <span>删除</span>
        </template>
      </el-popconfirm>
      <el-divider direction="vertical"></el-divider>

      <span @click="onView(scope.row)">查看</span>
    </template>
  </el-table-column>
</template>

<script>
import { getCurrentInstance, ref, inject } from 'vue'
export default {
  props: {
    /**
     * 列宽度
     */
    width: {
      type: String,
      default: '100px'
    }
  },
  emits: ['delete'],
  setup(props, { emit }) {
    const { ctx } = getCurrentInstance()
    // 上级页面的对话框组件引用
    const refFormDialog = inject('refFormDialog')
    // 当前表主键，默认为id
    const primaryKey = inject('primaryKey', 'id')
    // 对话框模式
    const dialogMode = inject('dialogMode', true)

    const { path } = ctx.$route
    const basePath = ref(path + (path.endsWith('/') ? '' : '/'))

    /**
     * 点击删除时触发
     */
    const onDelete = (row) => {
      emit("delete", { ids: [row[primaryKey]] })
    }

    /**
     * 点击行编辑链接时触发
     * @param {*} row
     */
    const onUpdate = (row) => {
      if (dialogMode) {
        refFormDialog.value.open({
          id: row[primaryKey],
          model: JSON.parse(JSON.stringify(row))
        })
      } else {
        ctx.$router.push({
          path: basePath.value + 'update',
          query: { id: row[primaryKey] }
        })
      }
    }

    /**
     * 查看
     */
    const onView = (row) => {
      ctx.$router.push({
        path: basePath.value + 'view',
        query: { id: row[primaryKey] }
      })
    }

    return {
      primaryKey,
      basePath,
      refFormDialog,
      onUpdate,
      onDelete,
      onView
    }
  }
}
</script>
<style lang="scss">
// 操作列
.action-column {
  text-align: right !important;
  .el-link,
  .el-button {
    font-weight: normal;
  }
  .el-link {
    line-height: 1;
  }
  span,
  a {
    color: $--color-primary;
    text-decoration: none;
    cursor: pointer;
    &:hover {
      text-decoration: underline;
    }
  }
}
</style>