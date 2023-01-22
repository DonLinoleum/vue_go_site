<template>
    <div class="footer">
        <div class="footer_img_wrapper">
            <img src="/img/footer_img_1.png" />
        </div>
        <div class="footer_text_wrapper">
            <h2>Excellent wine shop of Sophocles!</h2>
            <h2>2023 ©</h2>
        </div>
        <div class="footer_img_wrapper right">
            <img src="/img/footer_img_1.png" />
        </div>
    </div>
</template>

<script>
    export default{
        
    }
</script>

<style scoped>
    .footer{
        display:flex;
        justify-content: center;
        min-height: 10vh;
        background-color: #000000;
        font-family: 'Montserrat', sans-serif;
        color:white;
        margin-top: 3rem;
        padding: 0.5rem;
    }
    .footer_text_wrapper{
        margin-left:2rem;
        margin-right: 2rem;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
    .footer_img_wrapper img{
        width:6rem;
        height: 9rem;
    }

    .right{
        transform: scaleX(-1);
    }
</style>