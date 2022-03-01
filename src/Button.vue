<template>
   <div>
       <h3><slot name="title">标题</slot></h3>
       <p>
           <slot>内容</slot>
       </p>
   </div>
</template>
<script>
export default {
    
}
</script>