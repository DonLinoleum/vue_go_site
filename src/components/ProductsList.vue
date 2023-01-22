<template>
<div class="products-list">
    <product-item 
    v-for="(product,index) in products"
    :product="product"
    :index="index"
    ></product-item>
</div>
</template>

<script>
import ProductItem from './ProductItem.vue';
export default{
    components:{
        'product-item':ProductItem
    },
    props:{
        products:{
            type:Array
        }
    }
}
</script>

<style>
.products-list{
    display: flex;
    align-items: center;
    flex-direction: column;
}
</style>