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
    methods:{
        startAnimation(){
            setTimeout(()=>{
            let images = document.querySelectorAll('img')
            let headers_text = document.querySelectorAll('h2')
            headers_text.forEach(el=>{
                el.style.transform = 'translateX(0)'
            })

            images.forEach(el=>{
            el.style.opacity = '1'
        }) 
        },10)
        }
    },
    mounted(){
        this.startAnimation()
    }
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

    .h-wrapper h2 {
        transform: translateX(-60rem);
        transition: all 2s ease-out;
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
        opacity: 0;
        transition: opacity 2.5s ease-out;
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
        opacity: 0;
        box-shadow: 1px 1px 11px black;
        transition: opacity 2.5s ease-out;
    }
</style>