<template>
  <div class="page" v-if="!dialogMode">
    <page-header title="添加活动"></page-header>
    <el-row>
      <el-col :lg="3">
        <slot name="nav"></slot>
      </el-col>
      <el-col :lg="15">
      </el-col>
    </el-row>
    <div class="page-content-body">
      <slot></slot>
      <slot name="footer"></slot>
    </div>
  </div>

  <form-dialog ref="dialog" :title="pageTitle" @open="onOpen" @closed="onClosed" v-else>
    <slot></slot>
    <slot name="footer" slot="footer"></slot>

  </form-dialog>

</template>

<script>
import FormDialog from '@/components/Dialog'
import PageHeader from "@/components/PageHeader"

export default {
  name: "FormContainer",
  components: { PageHeader, FormDialog },
  props: {
    dialogMode: {
      type: Boolean,
      default: true
    },
    pageTitle: {
      type: String,
      default: '标题'
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

<style scoped>
</style>
