<template>
  <el-descriptions-item :span="span" :label="label">

    abcdefg

    <template #label>
      <slot name="label"></slot>
    </template>
  </el-descriptions-item>
</template>

<script>
export default {
  props: {
    label: {
      type: String
    },
    span: {
      type: Number,
      default: 1
    }
  }
}
</script>

<style lang="sass">
</style>
