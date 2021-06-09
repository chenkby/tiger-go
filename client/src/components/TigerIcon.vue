<template>
  <svg class="tiger-icon" aria-hidden="true">
    <use :xlink:href="iconHref"></use>
  </svg>
</template>

<script>
export default {
  props: {
    icon: {
      type: String,
      default: ''
    }
  },
  computed: {
    iconHref() {
      return `#tiger-icon-${this.icon}`
    }
  }
}
</script>

<style lang="scss">
.tiger-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
