import {createRouter, createWebHistory} from 'vue-router'
import Index from './../pages/Index.vue'
import Products from './../pages/Products.vue'
import About from './../pages/About.vue'
import ProductItemDetail from './../pages/ProductItemDetail.vue'

const routes = [
    {path:'/',component:Index},
    {path:'/products',component:Products},
    {path:'/about',component:About},
    {path: '/products/:id',component:ProductItemDetail}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router