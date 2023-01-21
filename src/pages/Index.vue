<template>
    <div class="main-index">
        <div class="h-wrapper">
            <h2>Huge selection of wines!</h2>
        </div>
        <div class="main-index-img">
            <img src="/img/header_img_1.jpg" />
        </div>
        <div class="h-wrapper">
            <h2>Best prices!</h2>
        </div>
        <div class="main-index-img__left">
            <img src="/img/header_img_2.jpg" />
        </div>
        <div class="h-wrapper">
            <h2>Glad to see you!</h2>
        </div>
        <div class="main-index-img">
            <img src="/img/header_img_3.jpg" />
        </div>
    </div>
</template>

<script>
export default{

}
</script>

<style scoped>
    .main-index{
        display: flex;
        font-family: 'Montserrat', sans-serif;
        flex-direction: column;
        width: 100%;
    }
    .h-wrapper{
        margin-top: 1rem;
        font-size: 4rem;
    }
    .main-index-img{
        margin-top:2rem;
        align-self:flex-end;
        margin-right: 3rem;
    }
    .main-index-img img{
        border-radius: 20px;
        max-width: 1200px;
        max-height: 600px;
        box-shadow: 1px 1px 11px black;
    }
    .main-index-img__left{
        margin-top:2rem;
        align-self:flex-start;
        margin-left: 3rem;
    }
    .main-index-img__left img{
        border-radius: 20px;
        max-width: 1200px;
        max-height: 600px;
        box-shadow: 1px 1px 11px black;
    }
</style>