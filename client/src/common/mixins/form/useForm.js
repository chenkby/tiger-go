import { useRoute } from 'vue-router'
import { getCurrentInstance, ref, inject, computed, provide } from 'vue'
import emitter from 'tiny-emitter/instance'

const useForm = (refForm) => {
  const onSubmit = () => {
    refForm.value.validate((valid) => {
      if (valid) {
        alert('submit!')
      } else {
        return false
      }
    })
  }

  const onValidate = (prop, isPass, errMessage) => {
    emitter.emit(`onValidated:${prop}`, errMessage)
  }

  return {
    onSubmit,
    onValidate
  }
}

const useDialog = () => {
  const formContainer = ref(null) // form-container引用
  /**
   * 打开表单对话框(主要供list使用)
   * @param {*} params
   */
  const open = (params) => {
    formContainer.value.open()
  }
  return {
    formContainer,
    open
  }
}

export default function (isDialogMode, labels, hints) {
  const { ctx } = getCurrentInstance()
  const id = ref(ctx.$route.query.id || null)
  const loading = ref(false)

  // 如果是移动端，不使用对话框模式
  const dialogMode = ctx.$isMobile ? false : isDialogMode

  provide('dialogMode', dialogMode)
  provide('labels', labels)
  provide('hints', hints)

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
    ...useForm(refForm),
    ...useDialog()
  }
}
