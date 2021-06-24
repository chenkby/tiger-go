<template>
  <el-table-column :prop="prop" class-name="download-column">
    <template #default="scope">
      <template v-if="scope.row[prop]">
        <el-tooltip :content="scope.row[prop]" :placement="placement" v-if="showTooltip">
          <el-link :href="scope.row[prop]" :target="target" :underline="false">
            <i :class="icon"></i><span>{{text}}</span>
          </el-link>
        </el-tooltip>
        <a :href="scope.row[prop]" :target="target" v-else>
          <i :class="icon"></i><span>{{text}}</span>
        </a>
      </template>
      <template v-else>-</template>
    </template>
  </el-table-column>

</template>

<script>
export default {
  props: {
    prop: {
      type: String
    },
    /**
     * tooltip位置
     */
    placement: {
      type: String,
      default: 'left'
    },
    /**
     * 是否显示tooltip
     */
    showTooltip: {
      type: Boolean,
      default: true
    },
    /**
     * 链接打开窗口
     */
    target: {
      type: String,
      default: '_blank'
    },
    /**
     * 显示的文字，不设置为字段值
     */
    text: {
      type: String,
      default: '下载'
    },
    icon: {
      type: String,
      default: 'el-icon-download'
    }
  },
  setup() {

  }
}
</script>

<style lang="scss">
.download-column {
  .el-link {
    font-weight: normal;
  }
  .el-link [class*='el-icon-'] + span {
    margin-left: 2px;
  }
}
</style>