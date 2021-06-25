<template>
  <el-radio-group v-model="controlValue" @change="onChange" :disabled="disabled">
    <template v-for="(item, key) in items" :key="key">
      <el-radio :disabled="getItemDisabled(item)" :label="getValue(item, key)">{{getLabel(item)}}</el-radio>
    </template>

  </el-radio-group>
</template>

<script>
import { ref } from 'vue'
export default {
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    items: {
      type: [Array, Object],
      default() {
        return [
          { value: 0, label: '禁用' },
          { value: 1, label: '启用' }
        ]
      }
    },
    /**
     * items的label键名
     */
    labelKey: {
      type: String,
      default: 'label'
    },
    /**
     * items的value键名
     */
    valueKey: {
      type: String,
      default: 'value'
    }
  },
  setup(props, { emit }) {
    const controlValue = ref(props.modelValue)
    const getValue = (item, key) => {
      return typeof item === 'object' ? item[props.valueKey] : key
    }
    const getLabel = (item) => {
      return typeof item === 'object' ? item[props.labelKey] : item
    }

    const getItemDisabled = (item) => {
      return item.disabled === true
    }

    const onChange = (value) => {
      emit('change', value)
    }
    return {
      controlValue,
      getValue,
      getLabel,
      getItemDisabled,
      onChange
    }
  }
}
</script>
