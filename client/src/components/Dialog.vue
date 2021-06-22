<template>
  <el-dialog custom-class="tiger-dialog" :title="title" v-model="visible" :width="width" :destroy-on-close="true" :close-on-click-modal="closeOnClickModal" @open="onOpen" @close="onClose" @closed="onClosed">
    <slot></slot>
    <template v-slot:footer>
      <slot name="footer"></slot>
    </template>
  </el-dialog>
</template>

<script>
import { defineComponent, ref } from 'vue'
export default {
  props: {
    title: {
      type: String,
      default: '标题'
    },
    width: {
      type: [String, Number],
      default: '30%'
    },
    closeOnClickModal: {
      type: [Boolean],
      default: false
    }
  },
  setup(props, { emit }) {
    const visible = ref(false)

    const open = () => {
      visible.value = true
    }

    const close = () => {
      visible.value = false
    }

    const onOpen = () => {
      emit('open')
    }

    const onClose = () => {
      emit('close')
    }

    const onClosed = () => {
      emit('closed')
    }

    return {
      visible,
      open,
      close,
      onOpen,
      onClose,
      onClosed
    }
  }
}
</script>

<style lang="scss">
.tiger-dialog {
  .el-dialog__header {
    background-color: #fafbfc;
    padding-top: 15px;
    padding-bottom: 15px;
    border-bottom: 1px solid #eee;
  }
  .el-dialog__headerbtn {
    top: 15px;
  }
  .el-dialog__footer {
    text-align: center;
    background-color: #fafbfc;
    padding-bottom: 10px;
  }
}
</style>
