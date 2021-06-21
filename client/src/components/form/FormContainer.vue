<template>
  <div class="tiger-form-container">
    <div class="page" v-if="!dialogMode">
      <page-header title="添加活动"></page-header>
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

    <form-dialog ref="dialog" :title="pageTitle" @open="onOpen" @closed="onClosed" v-else>
      <slot name="default"></slot>
      <slot name="footer"></slot>

    </form-dialog>
  </div>
</template>

<script>
import FormDialog from '@/components/Dialog.vue'
import PageHeader from "@/components/PageHeader.vue"
import { inject } from 'vue'
export default {
  name: "FormContainer",
  components: { PageHeader, FormDialog },
  props: {
    dialogMode: {
      type: Boolean,
      default: true
    },
    // 页面标题
    pageTitle: {
      type: String,
      default: '标题'
    }
  },
  setup(props) {
    //const dialogMode = inject('dialogMode', true)
    return {
      //dialogMode
    }
  },
  methods: {
    open() {
      this.$refs.dialog.open()
    },
    close() {
      this.$refs.dialog.close()
    },
    onOpen(params) {
      this.$emit('open', params)
    },
    onClosed(params) {
      this.$emit('closed', params)
    }
  }
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
}
</style>
