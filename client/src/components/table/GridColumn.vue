<template>
  <el-table-column :prop="prop" :label="label" :min-width="minWidth">
    <template #default="scope">
      <slot name="default" :model="scope"></slot>
    </template>
  </el-table-column>
</template>

<script>
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
  setup(props, { slots }) {
    console.log(slots.default())
  },
  created() {
    console.log(this)
  }

}

</script>