<template>
        <div class="product-item">
        <div class="product-item__props-wrapper">
            <div class="product-item__title"><h2>{{product.Title}}</h2></div>
            <div class="product-item__picture"><img :src="product.Picture"/></div>
            <div class="product-item__size">Size: <span>{{product.Size}} L</span></div>
            <div class="product-item__price">Price: <span>{{product.Price}} $</span></div>
        </div>
        <div class="product-item__description-wrapper">
            <div class="product-item__description">{{product.Description}}</div>
        </div>
    </div>   
</template>

<script>
import {mapActions} from 'vuex'
export default{
    data(){
        return{
            id: this.$route.params.id,
            product: {}
        }
    },
    methods:{
        ...mapActions({
            fetchOneProduct: 'products/fetchOneProduct'
        })
    },
    mounted(){
      let response = this.fetchOneProduct(this.id)
      response.then(data=>this.product = data)
    }
    
}
</script>

<style scoped>
.product-item__detail{
    display: flex;
    min-height:60vh;
}
.product-item{
    margin-top:1rem;
    margin-bottom: 1rem;
    display:flex;
    justify-content: center;
    width:85%;
    box-shadow: 11px 1px 30px #9c91a7;
    border-radius: 20px;
    padding:2rem;
    cursor: pointer;
}

.product-item__title{
    font-size: 2rem;
}

.product-item__description{
    font-size: 2rem;
}

.product-item__size{
    font-size: 3rem;
    align-self: baseline;
    margin-top:1rem;
}

.product-item__price{
    font-size: 3rem;
    align-self: baseline;
}

.product-item__size span{
    font-size: 4rem;
}

.product-item__price span{
    font-size: 4rem;
}

.product-item__props-wrapper{
    display: flex;
    flex-direction: column;
    align-items: center;
}

.product-item__description-wrapper{
    display: flex;
    flex-direction: column;
    align-items: center;
    margin:5rem 5rem;
}

.product-item__picture img{
    margin-top: 2rem;
    width: 30rem;
    height:30rem;
    border-radius: 20px;
    box-shadow: 1px 1px 11px black;
}
</style>