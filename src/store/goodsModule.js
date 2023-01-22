
export const goodsModule = {
    state(){
        return {
            host: "http://localhost:8080",
            products: []
        }
    },
    getters:{
        getHost(state){
            return state.host
        },
        getProducts(state){
            return state.products
        }
    },
    mutations:{
        setProducts(state,products){
            state.products = products
        }    
    },
    actions:{
        async fetchProducts({commit,state}){
                try{
                    const response = await fetch(state.host + "/get-all-products",{
                        method:"GET",
                        headers:{"Accept":"application/json"}
                    })
                    const responseJson = await response.json()
                    commit ('setProducts',responseJson)
                }
                catch(e){
                    alert(e)
                }
        }
    },
    namespaced:true 
}