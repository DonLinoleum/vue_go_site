<template>
    <div class="background">
        <header-app></header-app>
        <div class="app">
            <router-view></router-view>
        </div>
        <footer-app></footer-app>
        </div>
</template>

<script>
    import Header from './components/Header.vue';
    import Footer from './components/Footer.vue'
    export default{
        components:{
            'header-app' : Header,
            'footer-app' : Footer
        }
    }
</script>

<style>
*{
    margin:0;
    padding:0;
    box-sizing: border-box;
}
.app{
   display: flex; 
   font-family: 'Montserrat', sans-serif;
}
.background{
    background: url('/img/background.jpg'),rgba(255,255,255,0.8);
    background-size: cover;
    background-blend-mode: color;
    background-repeat: no-repeat;
    
}
</style>