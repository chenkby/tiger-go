<template>
  <div class="page" v-if="!dialogMode">
    <page-header title="添加活动"></page-header>
    <el-row>
      <el-col :lg="3">
        <slot name="nav"></slot>
      </el-col>
      <el-col :lg="15">
        <slot></slot>
      </el-col>
    </el-row>
    <slot name="footer"></slot>
  </div>

  <form-dialog ref="dialog" :title="pageTitle" @open="onOpen" @closed="onClosed" v-else>
    <slot></slot>
    <slot name="footer"></slot>

  </form-dialog>

</template>

<script>
import FormDialog from '@/components/Dialog.vue'
import PageHeader from "@/components/PageHeader.vue"
import { inject } from 'vue'
export default {
  name: "FormContainer",
  components: { PageHeader, FormDialog },
  props: {
    labelWidth: {
      type: String,
      default: '100px'
    },
    pageTitle: {
      type: String,
      default: '标题'
    }
  },
  setup() {
    const dialogMode = inject('dialogMode', true)
    return {
      dialogMode
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
