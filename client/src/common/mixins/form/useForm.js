import { useRoute } from 'vue-router'
import { getCurrentInstance, ref, inject, computed, provide } from 'vue'
import emitter from 'tiny-emitter/instance'

const useForm = (refForm, formData) => {
  const onSubmit = () => {
    console.log(formData.value)
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

const useData = (id, formData, infoApi) => {
  const loading = ref(false)
  const loadData = async () => {
    if (id.value) {
      loading.value = true
      const res = await infoApi({ id: id.value })
      loading.value = false
      if (res.err_code === 0) {
        formData.value = { ...formData.value, ...res.data }
      }
    }
  }
  return {
    loadData
  }
}

export default function (params, apis) {
  const { dialogMode: isDialogMode, labels, hints, formData } = params
  const { createApi, updateApi, infoApi } = apis
  const { ctx } = getCurrentInstance()
  const id = ref(ctx.$route.query.id || null)

  // 如果是移动端，不使用对话框模式
  const dialogMode = ctx.$isMobile ? false : isDialogMode

  provide('dialogMode', dialogMode)
  provide('labels', labels)
  provide('hints', hints)

  const refForm = ref(null)

  // 表单label位置
  const labelPosition = computed(() => (ctx.$isMobile ? 'top' : 'right'))

  const { loadData } = useData(id, formData, infoApi)

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
    hints,
    formData,
    loadData,
    onCancel,
    ...useForm(refForm, formData),
    ...useDialog()
  }
}
