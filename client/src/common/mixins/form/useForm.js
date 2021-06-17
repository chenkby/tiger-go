import { getCurrentInstance, ref } from 'vue'

export default function () {
  const { ctx } = getCurrentInstance()
  const id = ref(ctx.$route.query.id || null)
  const loading = ref(false)

  const loadData = () => {}
}
