<template>
  <el-form-item class="tiger-form-item" :prop="prop" :label="propLabel" :show-message="showMessage" :label-width="labelWidth" :required="required" :rules="rules" :size="size">
    <slot name="default">

    </slot>
    <span class="hint">
      <slot name="hint">{{hint}}</slot>
    </span>
    <span class="hint hint-error">{{errMessage}}</span>
    <template #error>
      <span></span>
    </template>
  </el-form-item>
</template>

<script>
import { ref, onMounted, onUnmounted, inject, computed } from 'vue'
import emitter from 'tiny-emitter/instance'
export default {
  props: {
    // 提示文本
    hint: {
      type: String
    },
    prop: {
      type: String,
      required: true
    },
    label: {
      type: String
    },
    labelWidth: {
      type: String
    },
    required: {
      type: Boolean,
      default: false
    },
    rules: {
      type: Object
    },
    showMessage: {
      type: Boolean,
      default: true
    },
    size: {
      type: String
    }
  },
  setup(props, { emit }) {
    const modelValue = ref("")
    // 输入框改变时
    const onValueInput = () => {
      emit('update:modelValue', modelValue.value)
    }

    // 接收这个字段的错误信息
    const errMessage = ref("")
    onMounted(() => {
      emitter.on(`onValidated:${props.prop}`, (msg) => {
        errMessage.value = msg
      })
    })
    onUnmounted(() => {
      emitter.off(`onValidated:${props.prop}`)
    })

    // label
    const labels = inject('labels', {})
    const propLabel = computed(() => {
      if (props.label !== undefined && props.label !== null) {
        return props.label
      } else {
        return labels[props.prop]
      }
    })
    return {
      modelValue,
      onValueInput,
      errMessage,
      propLabel
    }
  }
}
</script>
<style lang="scss">
.tiger-form-item {
  .hint {
    font-size: 12px;
    color: #888;
    line-height: 1.5em;
    display: block;
    margin-top: 3px;
    a {
      color: $--color-primary;
      text-decoration: none;
      &:hover {
        text-decoration: underline;
      }
    }
    .el-link {
      font-size: 12px;
      font-weight: normal;
      line-height: 1.5em;
      vertical-align: inherit;
    }
  }
  .hint-error {
    color: $--color-danger;
    display: none;
    &.el-zoom-in-top-enter-active,
    &.el-zoom-in-top-leave-active {
      opacity: 1;
      transform: none !important;
      transition: none !important;
      transform-origin: center top;
    }
    &.el-zoom-in-top-enter,
    &.el-zoom-in-top-leave-active {
      opacity: 0;
      transform: none !important;
    }
  }
  &.is-error {
    .hint {
      display: none;
    }
    .hint-error {
      display: block;
    }
  }
}
</style>