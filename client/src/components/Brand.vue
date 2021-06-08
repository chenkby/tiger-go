<template>
  <div class="brand" :class="{cursor: this.link}">
    <img class="logo" src="https://ergev2.tuxiaobei.com/manage/img/logo.93fa7efe.png" @click="goHome" />
    <div class="name" @click="goHome">{{name}}</div>
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
    },
    /**
     * 点击时的链接地址
     */
    link: {
      type: [String, Boolean],
      default: '/'
    }
  },
  data() {
    return {

    }
  },
  methods: {
    goHome() {
      if (this.link) {
        this.$router.push(this.link)
      }
    },
    onClick() {
      if (this.link) {
        this.$router.push('/')
      }
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
  &.cursor {
    cursor: pointer;
  }
}
</style>
