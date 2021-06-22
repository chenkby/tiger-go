<template>
  <!-- 对话框按钮 -->
  <div class="dialog-form-buttons" v-if="dialogMode">
    <el-button @click="onCancel">取消</el-button>
    <el-button type="primary" @click="onSubmit">确定</el-button>
  </div>
  <!--非对话框模式-->
  <el-row v-else>
    <el-col :lg="3"></el-col>
    <el-col :lg="15" class="form-buttons" :style="{'padding-left':indentWidth}">
      <el-button v-if="showSubmit" type="primary" @click="onSubmit">{{submitTitle}}</el-button>
      <el-button v-if="showCancel" @click="onCancel">取消</el-button>
    </el-col>
  </el-row>

</template>

<script>
import { computed, getCurrentInstance, inject } from 'vue'
export default {
  props: {
    /**
     * 提交按钮标题
     */
    submitTitle: {
      type: String,
      default: '保存'
    },
    /**
     * 是否显示提交按钮
     */
    showSubmit: {
      type: Boolean,
      default: true
    },
    /**
     * 是否显示取消按钮
     */
    showCancel: {
      type: Boolean,
      default: true
    },
    /**
     * 按钮条左边距，移动端不起作用
     */
    indent: {
      type: [String],
      default: '100px'
    }
  },
  setup(props, { emit }) {
    const { ctx } = getCurrentInstance()
    const dialogMode = inject('dialogMode', true)
    /**
     * 点击提交表单按钮触发
     */
    const onSubmit = () => {
      emit('submit')
    }
    /**
     * 点击取消按钮触发
     */
    const onCancel = () => {
      emit('cancel')
    }

    const indentWidth = computed(() => ctx.$isMobile ? 0 : props.indent)
    return {
      dialogMode,
      indentWidth,
      onSubmit,
      onCancel
    }
  },
  emits: ['submit', 'cancel']
}
</script>

<style lang="scss">
.form-buttons {
  padding: 10px 0 32px 0;
}
body.device-mobile {
  .form-buttons {
    text-align: center;
  }
}
</style>

