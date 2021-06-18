import { useRoute } from 'vue-router'
import { getCurrentInstance, ref, inject } from 'vue'

export default function () {
  const { ctx } = getCurrentInstance()
  const id = ref(ctx.$route.query.id || null)
  const loading = ref(false)

  const loadData = () => {}

  // 取消，返回上一页或者关闭对话框
  const onCancel = () => {
    if (ctx.dialogMode) {
      ///this.$refs.container.close()
    } else {
      ctx.$router.go(-1)
    }
  }
  return {
    onCancel
  }
}
