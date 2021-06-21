import { useRoute } from 'vue-router'
import { getCurrentInstance, ref, inject, computed, provide } from 'vue'

const useForm = (refForm) => {
  const onSubmit = () => {
    console.log('refForm', refForm)
    refForm.value.validate((valid) => {
      if (valid) {
        alert('submit!')
      } else {
        console.log('error submit!!')
        return false
      }
    })
  }
  return {
    onSubmit
  }
}

export default function (dialogMode) {
  const { ctx } = getCurrentInstance()
  const id = ref(ctx.$route.query.id || null)
  const loading = ref(false)
  provide('dialogMode', dialogMode)

  const refForm = ref(null)

  // 表单label位置
  const labelPosition = computed(() => (ctx.$isMobile ? 'top' : 'right'))

  const loadData = () => {}

  // 取消，返回上一页或者关闭对话框
  const onCancel = () => {
    if (dialogMode) {
      ///this.$refs.container.close()
    } else {
      ctx.$router.go(-1)
    }
  }
  return {
    refForm,
    dialogMode,
    labelPosition,
    onCancel,
    ...useForm(refForm)
  }
}
