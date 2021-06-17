<template>
  <el-dialog :title="title" v-model:visible="visible" :width="width" :destroy-on-close="true" :close-on-click-modal="closeOnClickModal" @open="onOpen" @close="onClose" @closed="onClosed">
    <slot></slot>
    <template v-slot:footer>
      <slot name="footer"></slot>
    </template>
  </el-dialog>
</template>

<script>
import { ref } from 'vue'
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
  setep(props, { emit }) {
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

<style scoped>
</style>
