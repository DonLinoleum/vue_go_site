<template>
    <input v-bind:value="modelValue" v-on:input="updateInputValue" />
</template>

<script>
    export default{
        name:"custom-input",
        data(){
            return{

            }
        },
        props:{
            modelValue:{
                type:[String,Number]
            }
        },
        methods:{
            updateInputValue(event){
                this.$emit("update:modelValue",event.target.value)
            }
        }
    }
</script>

<style scoped>
    input{
        width:85%;
        margin-top:2rem;
        font-size: 2rem;
        border-radius: 10px;
        border: 2px solid #6e6e27;
        padding:5px
    }
</style>