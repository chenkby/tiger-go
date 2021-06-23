<template>
  <div class="tiger-form-container">
    <!-- 非对话框模式布局 -->
    <div class="page" v-if="!dialogMode">
      <page-header :title="pageTitle"></page-header>
      <el-row>
        <el-col :lg="3">
          <slot name="left"></slot>
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
    // 对话框宽度
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

  .quick-nav {
    border-right: 1px solid #efefef;
    text-align: right;
    padding-right: 20px;
    width: 120px;
    position: fixed;

    .item {
      padding: 5px 0;
      font-size: 12px;
      .el-link.el-link--default {
        color: #333;
        &:hover {
          color: $--color-primary;
        }
      }
    }
    .title {
      font-size: 16px;
      margin-bottom: 5px;
      color: #777;
    }
    .el-link {
      font-weight: normal;
      font-size: 80%;
    }
  }
}
body.device-mobile {
  .tiger-form-container {
    // 手机模式下，表单label下边距
    .el-form--label-top .el-form-item__label {
      padding: 0;
      line-height: 30px;
    }
  }
}
</style>
