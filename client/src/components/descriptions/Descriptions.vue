<template>
  <el-descriptions class="tiger-descriptions" :column="columnNum" :title="title" :border="border" :direction="direction" :size="size" :extra="extra">
    <template #title>
      <slot name="title"></slot>
    </template>
    <template #extra>
      <slot name="extra"></slot>
    </template>
    <slot name="default"></slot>
  </el-descriptions>
</template>

<script>
export default {
  props: {
    labelWidth: {
      type: [String, Number],
      default: '100px'
    },
    title: {
      type: String
    },
    border: {
      type: Boolean,
      default: false
    },
    column: {
      type: Number,
      default: 3
    },
    direction: {
      type: String,
      default: 'horizontal'
    },
    size: {
      type: String
    },
    extra: {
      type: String
    }
  },
  computed: {
    /**
     * 列数，如果手机端，只使用1列，否则使用设置的列数
     */
    columnNum() {
      return this.$isMobile ? 1 : this.column
    }
  }
}
</script>

<style lang="scss">
.tiger-descriptions + .tiger-descriptions {
  margin-top: 30px;
}

.el-descriptions__label {
  text-align: right !important;
  color: #666;
  min-width: 150px;
  margin-right: 20px;
  display: inline-block;
  &.is-bordered-label {
    font-weight: normal;
  }
}
.el-descriptions__body {
  color: #111;
}
.el-descriptions__content {
  color: #000;
}
</style>
