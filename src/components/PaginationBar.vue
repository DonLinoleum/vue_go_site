<template>
 <div class="pagination-wrapper">
    <div :class="{'current' : item==pageNumber,'paginationItem':true}" 
    v-for="item in totalPages"
    v-on:click="changePage(item)">
        <span class="pagination-item__text">{{item}}</span>
    </div>
 </div>

</template>

<script>
export default{
    props:{
        limit:{type:Number},
        pageNumber:{type:Number},
        totalPages:{type:Number}
    },
    methods:{
        changePage(item){
            this.$emit("changePage",item)
        }
    }
}
</script>

<style scoped>
    .pagination-wrapper{
        display: flex;
        margin-top:1rem;
    }
    .paginationItem{
        display: flex;
        justify-content: center;
        align-items: center;
        width: 4rem;
        height: 4rem;
        border: 2px solid black;
        border-radius: 10px;
        margin-right: 1rem;
        cursor: pointer;
        box-shadow: 1px 1px 12px black;
        background-color: aliceblue;
        transition: all 0.5s ease-out;
    }

    .current{
        box-shadow: 1px 1px 12px rgb(255, 0, 0);
        border-color: rgb(255, 0, 0);
    }
    .paginationItem:hover{
        background-color: black;
    }

    .paginationItem:hover .pagination-item__text{
        color: white;
    }

    .pagination-item__text{
        font-size: 2rem;
        transition: all 0.5s ease-out;
    }
</style>