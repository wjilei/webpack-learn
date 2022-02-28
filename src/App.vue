<template>
    <div>
        <Hello name="zhangsan" />
    </div>
</template>
<script>
import Hello from "./Hello.vue"
export default {
    props: {

    },
    components: {
        Hello
    }
}
</script>