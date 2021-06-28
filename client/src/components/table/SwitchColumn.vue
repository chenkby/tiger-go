<template>
  <el-switch v-model="row[prop]" @click="onClick(row[primaryKey])" :disabled="disabled" active-color="#13ce66" :active-value="activeValue" :inactive-value="inactiveValue" @change="onChange(row[primaryKey])" :loading="loading"></el-switch>
</template>

<script>
import { ref, inject } from 'vue'
export default {
  props: {
    // prop: {
    //   type: String,
    //   required: true
    // },
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
  mounted() {
    console.log(this)
  },
  setup(props, { slots }) {
    const modelValue = ref(1)
    const loading = ref(false)
    const primaryKey = inject('primaryKey')

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
      onChange,
      beforeChange,
      onClick
    }
  }
}
</script>