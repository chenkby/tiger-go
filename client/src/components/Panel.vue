<template>
  <div class="tiger-panel" :class="size" :id="name">
    <div class="panel-header">
      <div class="title">{{title}}</div>
      <slot class="right" name="right"></slot>
    </div>
    <div class="panel-body">
      <slot></slot>
    </div>
    <div class="panel-footer">

    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      required: true
    },
    name: {
      type: [String, Number],
      default: 'panel1'
    },
    size: {
      type: String,
      default: 'large'
    }
  },
  setup(props) {

  }
}
</script>

<style lang="scss">
.tiger-panel {
  .panel-header {
    padding: 0 0 30px 0;
    line-height: 1;
    align-items: center;
    display: flex;
    justify-content: space-between;
    .title {
      border-left: 3px solid $--color-primary;
      font-size: 16px;
      padding-left: 10px;
      color: #111;
      font-weight: 400;
    }
    .right {
      text-align: right;
    }
  }
}
.tiger-panel + .tiger-panel {
  margin-top: 30px;
}
</style>
