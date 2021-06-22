<template>
  <div class="tiger-form-container">
    <!-- 非对话框模式布局 -->
    <div class="page" v-if="!dialogMode">
      <page-header :title="pageTitle"></page-header>
      <el-row>
        <el-col :lg="3">
          <slot name="nav"></slot>
        </el-col>
        <el-col :lg="15">
          <slot name="default"></slot>
        </el-col>
      </el-row>
      <slot name="footer"></slot>
    </div>

    <!-- 对话框模式布局 -->
    <form-dialog ref="dialog" :title="pageTitle" @open="onOpen" @closed="onClosed" :width="dialogWidth" v-else>
      <slot name="default"></slot>
      <template #footer>
        <slot name="footer"></slot>
      </template>

    </form-dialog>
  </div>
</template>

<script>
import FormDialog from '@/components/Dialog.vue'
import PageHeader from "@/components/PageHeader.vue"
import { ref, inject } from 'vue'
export default {
  name: "FormContainer",
  components: { PageHeader, FormDialog },
  props: {
    dialogWidth: {
      type: String,
      default: '50%'
    },
    // 页面标题
    pageTitle: {
      type: String,
      default: '标题'
    }
  },
  setup(props, { emit }) {
    const dialog = ref(null)
    const dialogMode = inject('dialogMode', false)

    /**
     * 打开对话框
     */
    const open = () => {
      dialog.value.open()
    }

    /**
     * 关闭对话框
     */
    const close = () => {
      dialog.value.close()
    }

    const onOpen = (params) => {
      emit('open', params)
    }
    const onClosed = (params) => {
      emit('closed', params)
    }
    return {
      dialogMode,
      dialog,
      open,
      close,
      onOpen,
      onClosed
    }
  },
  emits: ['open', 'closed']
}
</script>

<style lang="scss">
.tiger-form-container {
  .el-form-item__label {
    // label后面加个中文冒号
    &:after {
      content: '：';
    }
  }
  .el-form-item__error--inline {
    display: inherit;
    margin-left: 0;
    font-size: 12px;
    line-height: 1.5em;
    margin-top: 3px;
    padding-top: 0;
  }
}
</style>
