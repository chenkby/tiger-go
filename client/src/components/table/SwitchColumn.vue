<template>
  <el-table-column prop="free">
    <template #default="scope">
      <el-switch v-model="scope.row.free" :disabled="disabled" active-color="#13ce66" :active-value="activeValue" :inactive-value="inactiveValue" @change="onChange" :loading="loading"></el-switch>
    </template>
  </el-table-column>
</template>

<script>
import { ref } from 'vue'
export default {
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    activeValue: {
      type: [Boolean, String, Number],
      default: 1
    },
    inactiveValue: {
      type: [Boolean, String, Number],
      default: 0
    },
    scope: {
      type: Object
    }
  },

  setup(props) {
    console.log(props.scope)
    const modelValue = ref(1)
    const loading = ref(false)

    const onBeforeChange = (value) => {
      loading.value = true
      console.log(value)
    }
    const onChange = (value) => {
      console.log(value)
    }
    return {
      loading,
      modelValue,
      onChange,
      onBeforeChange
    }
  }
}
</script>