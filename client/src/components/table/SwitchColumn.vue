<template>
  <el-switch v-model="row[prop]" :disabled="disabled" active-color="#13ce66" :active-value="activeValue" :inactive-value="inactiveValue" @change="onChange" :loading="loading"></el-switch>
</template>

<script>
import { ref, inject } from 'vue'
export default {
  props: {
    prop: {
      type: String,
      required: true
    },
    // label: {
    //   type: String
    // },
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
    api: {
      type: [Function, Promise]
    },
    row: {
      tyep: Object,
      default() {
        return {

        }
      }
    }
  },
  setup(props, { slots }) {
    const modelValue = ref(1)
    const loading = ref(false)
    const primaryKey = inject('primaryKey')
    // const prop = inject('prop')
    console.log("aaaaaa", primaryKey)

    const beforeChange = async (id) => {
      // loading.value = true
      // if (props.api) {
      //   const res = await props.api({
      //     ids: id,
      //     attribute: props.prop,
      //     value: 1 - modelValue.value
      //   })
      //   loading.value = false
      //   if (res.err_code === 0) {
      //     return true
      //   } else {
      //     return false
      //   }
      // }
    }
    const onChange = (value, value2) => {
      console.log('aa', value, value2)
    }

    const onClick = (id) => {
      console.log('click', id, modelValue.value)
    }


    return {
      loading,
      modelValue,
      primaryKey,
      //prop,
      onChange,
      beforeChange,
      onClick
    }
  }
}
</script>