<template>
    <div>
        Hello {{name}}
        <Button>新文本</Button>
    </div>
</template>
<script>

import Button from "./Button.vue"
export default {
    data(){
        return {

        }
    },
    props: {
        name: {
            type: String,
        }
    },
    components: {
        Button
    }
}
</script>