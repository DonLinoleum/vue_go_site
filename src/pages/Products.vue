<template>
    <div class="main-products">
    <custom-input
    placeholder="Search...Type name of product."
    v-model="searchQuery"
    v-on:update:modelValue="setSearchQuery(searchQuery)"
    ></custom-input>
    
      <h1>Our wines:</h1>
      <h2 v-if="getSearchedProducts.length<=0">Sorry, no products found</h2>
      <products-list v-else :products="getSearchedProducts"></products-list>
      
    <pagination-bar 
        v-bind:pageNumber="pageNumber"
        v-bind:limit="limit"
        v-bind:totalPages="getTotalPages"
        v-on:changePage="changePage"
        >
    </pagination-bar>

    </div>
</template>

<script>
import {mapActions,mapGetters,mapMutations } from 'vuex'
import ProductsList from '../components/ProductsList.vue'
import PaginationBar from '../components/PaginationBar.vue'
export default{
    data(){
        return{
          searchQuery:"",
          limit : 3,
          pageNumber : 1
        }
    },
    components:{
        'products-list':ProductsList,
        'pagination-bar': PaginationBar
    },
    methods:{
        ...mapMutations({
            setProducts:'products/setProducts',
            setSearchQuery:'products/setSearchQuery'
        }),
        ...mapActions({
            fetchProducts:'products/fetchProducts'
        }),
        changePage(e){
            this.pageNumber = e
            this.fetchProducts(
            {
                pageNumber:this.pageNumber,
                limit:this.limit
            })
            window.scrollTo({ top: 0, behavior: 'smooth' });
        }
    },
    computed:{
        ...mapGetters({
            getProducts:'products/getProducts',
            getSearchedProducts: 'products/getSearchedProducts',
            getSearchQuery:'products/getSearchQuery',
            getTotalPages : 'products/getTotalPages'
        }),   
    },
    watch:{},
    mounted(){
        this.fetchProducts(
            {
                pageNumber:this.pageNumber,
                limit:this.limit
            })
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