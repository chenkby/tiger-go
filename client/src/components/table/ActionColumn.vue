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

      <router-link :to="{path: '/article/view', query:{id:scope.row[primaryKey]}}">查看</router-link>
    </template>
  </el-table-column>
</template>

<script>
import { inject } from 'vue'
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
  emits: ['update', 'delete'],
  setup(props, { emit }) {
    // 当前表主键，默认为id
    const primaryKey = inject('primaryKey', 'id')

    const onDelete = (row) => {
      emit("delete", { ids: [row[primaryKey]] })
    }
    const onUpdate = (row) => {
      emit('update', row)
    }
    return {
      primaryKey,
      onUpdate,
      onDelete
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