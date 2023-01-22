<template>
    <div class="main-products">
      <h1>Our wines:</h1>
        <products-list :products="getProducts"></products-list>
    </div>
</template>

<script>
import { mapState,mapActions,mapGetters,mapMutations } from 'vuex'
import ProductsList from '../components/ProductsList.vue'
export default{
    data(){
        return{

        }
    },
    components:{
        'products-list':ProductsList
    },
    methods:{
        ...mapMutations({
            setProducts:'products/setProducts'
        }),
        ...mapActions({
            fetchProducts:'products/fetchProducts'
        }),

    },
    computed:{
        ...mapGetters({
            getProducts:'products/getProducts'
        })
    },
    watch:{},
    mounted(){
        this.fetchProducts()
    }
}
</script>

<style>
    .main-products{
        display:flex;
        flex-direction: column;
        width: 100%;
        align-items: center;
        min-height:60vh;
        font-family: 'Montserrat', sans-serif;
    }
    h1{
        font-size: 4rem;
    }
</style>