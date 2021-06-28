<template>
  <el-table-column :prop="prop" :label="label" :min-width="minWidth">
    <template #default="{row,column,$index,store}">
      <slot :row="row" :column="column" :$index="$index" :store="store">{{row[prop]}}</slot>
    </template>
  </el-table-column>
</template>

<script>
import { ref, provide } from 'vue'
export default {
  props: {
    prop: {
      type: String,
      required: true
    },
    label: {
      type: String
    },
    minWidth: {
      type: String,
      default: '100px'
    },
    width: {
      type: String
    }
  },
  setup(props, { emit }) {
    provide('prop', props.prop)
  }
}

</script>