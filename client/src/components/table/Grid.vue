<template>
  <el-table :data="data" @selection-change="onSelectionChange" :size="size" ref="refTable">
    <template #default>
      <slot name="default"></slot>
    </template>
    <template #empty>
      <el-empty :image-size="140">
        <el-button type="primary" icon="el-icon-refresh" @click="refresh">刷新</el-button>
      </el-empty>
    </template>
  </el-table>
</template>

<script>
import { ref, provide } from 'vue'
export default {
  props: {
    data: {
      type: [Array]
    },
    primaryKey: {
      type: String,
      required: true
    },
    size: {
      type: String,
      default: 'medium'
    }
  },
  setup(props, { emit }) {
    provide("primaryKey", props.primaryKey)
    const onSelectionChange = () => {
      emit('selection-change')
    }
    const refresh = () => {
      emit('refresh')
    }
    return {
      onSelectionChange,
      refresh
    }
  }
}
</script>