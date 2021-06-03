<template>
<div class="page">

</div>
</template>

<script>
import { defineComponent } from '@vue/composition-api'

export default defineComponent({
    setup() {
        const name = ref("aaa")
        function test() {
            console.log('这是测试方法')
        }
        return {
            test
        }
    },
})
</script>

