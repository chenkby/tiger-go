<template>
  <div class="tiger-hint-block">
    <span :style="{visibility:show ? 'visible': 'hidden'}">
      <slot></slot>
    </span>
  </div>
</template>

<script>
/**
 * 用于在表单控件下方放一段输入提示文字
 */
import { ref, onMounted, onUnmounted } from 'vue'
import emitter from 'tiny-emitter/instance'
export default {
  props: {
    prop: {
      type: String,
      required: true
    }
  },
  setup(props) {
    console.log('p', props.prop)
    const show = ref(true)
    onMounted(() => {
      emitter.on(`onValidated:${props.prop}`, (isPass) => {
        show.value = isPass
        console.log('我触发了', isPass)
      })
    })
    onUnmounted(() => {
      emitter.off(`onValidated:${props.prop}`)
    })
    return {
      show
    }
  }
}
</script>

<style lang="scss">
.tiger-hint-block {
  color: #777;
  font-size: 12px;
  line-height: 1.5em;
  min-height: 16px;
  margin-top: 3px;
  visibility: visible;

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
</style>
