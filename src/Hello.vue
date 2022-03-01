<template>
    <div>
        Hello {{name}}
        <Button>
            <template v-slot:title>
                测试标题
            </template>
            内容
        </Button>
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