<template>
    <div class="header">
        <div class="header_logo"><img src="/img/header_sophocles.png"/></div>
        <div class="header_navbar__wrapper" >
            <h2>Excellent wine shop of Sophocles!</h2>
            <ul class="header_navbar">
                <li v-on:click="$router.push('/')">Main</li>
                <li v-on:click="$router.push('/products')">Products</li>
                <li v-on:click="$router.push('/about')">About</li>
            </ul>
        </div>
        <div class="header_logo__right"><img src="/img/header_sophocles.png"/></div>
    </div>
</template>

<script>
    export default{
        
    }
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Lobster&family=Montserrat:wght@300;400;600&display=swap');
    .header{
        display: flex;
        justify-content: center;
        min-height: 10vh;
        background-color: #000000;
        font-family: 'Montserrat', sans-serif;
        box-shadow: 0.1px 8px 31px rgb(177, 190, 57);
        align-items: center;
        color: white;
    }

    .header_navbar__wrapper{
        display:flex;
        flex-direction: column;
    }
    .header_navbar{
        list-style: none;
        display: flex;
        padding: 1rem;
    }
    .header_navbar li{
        position: relative;
        margin-right:1rem;
        cursor: pointer;
        font-size: 2rem;
        transition: width 0.3s ease-out;
    }
    .header_navbar li:after{
        display: block;
        position: absolute;
        content: "";
        border-bottom: 3px solid rgb(124, 124, 45);
        width: 0;
        transition: width 0.3s ease-out;
    }
    .header_navbar li:hover::after{
        width: 100%;
    }

    .header_logo img{
        width:11rem;
        height: 10rem;
        margin-top: 1rem;
        margin-bottom: 1rem;
    }
    .header_logo__right img{
        width:11rem;
        height: 10rem;
        transform: scaleX(-1);
        margin-top: 1rem;
        margin-bottom: 1rem;
    }
</style>