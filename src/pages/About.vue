<template>
    <div class="main-about">
        <h1>About our shop</h1>
        <div class="main-about__text">
            A good “About Us” page is concise, pithy, and informative. Not only does it give valuable information about a company, its history, its products, and its goals; it also has personality and is visually engaging. Never underestimate the power of your About Us page, and don’t forget to  alter your story, mission, and accomplishments as your business evolves.
        </div>
        <div class="main-about__img"><img src="/img/about-img.jpg"/></div>
    </div>
</template>

<script>
export default{
    
}
</script>

<style>
.main-about{
        font-family: 'Montserrat', sans-serif;
        align-self: center;
        display:flex;
        flex-direction: column;
        align-items: center;
        min-height:60vh;
        width: 85%;
    }
    .main-about__img{
        border-radius: 20px;
    }
    .main-about__img img{
        width:60rem;
        height:40rem;
        border-radius: 20px;
        box-shadow: 11px 1px 30px #9c91a7;
    }
    .main-about__text{
        font-size: 2rem;
        margin-top:2rem;
        margin-bottom:2rem;
    }
</style>