<template>
  <div class="brand">
    <img class="logo" src="https://ergev2.tuxiaobei.com/manage/img/logo.93fa7efe.png" />
    <div class="name">{{name}}</div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default {
  props: {
    logo: {
      type: String
    },
    name: {
      type: String,
      default: '系统名称'
    }
  },
  data() {
    return {

    }
  }
}
</script>

<style lang="scss">
.brand {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: auto;
  .logo {
    width: 30px;
    height: 30px;
  }
  .name {
    color: #2c3e50;
    font-size: 18px;
    font-weight: normal;
    margin-left: 10px;
  }
}
</style>
