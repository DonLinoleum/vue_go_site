
export const goodsModule = {
    state(){
        return {
            host: "http://localhost:8080",
            products: [],
            searchQuery: "",
            totalPages:10
        }
    },
    getters:{
        getHost(state){
            return state.host
        },
        getProducts(state){
            return state.products
        },
        getSearchedProducts(state){
            if (state.searchQuery)
                return state.products.filter(el=>el.Title.toLowerCase().includes(state.searchQuery.toLowerCase()))
            else
                return state.products
        },
        getSearchQuery(state){
            return state.searchQuery
        },
        getTotalPages(state){
            return state.totalPages
        }
    },
    mutations:{
        setProducts(state,products){
            state.products = products
        },
        setSearchQuery(state,value){
            state.searchQuery = value
        },
        setTotalPages(state,value){
            state.totalPages = value
        }    
    },
    actions:{
        async fetchProducts({commit,state},payload){
                try{
                    const response = await fetch(state.host + "/get-all-products?limit="+payload.limit+"&pageNumber="+payload.pageNumber,{
                        method:"GET",
                        headers:{"Accept":"application/json"}
                    })
                    const responseJson = await response.json()
                    commit ('setProducts',responseJson)
                    commit ('setTotalPages',Math.ceil(response.headers.get('Total-Count')/payload.limit))
                }
                catch(e){
                    alert(e)
                }
        },
        async fetchOneProduct({commit,state},id){
            try{
                const response = await fetch(state.host + "/get-product/"+id,{
                    method:"GET",
                    headers:{"Accept":"application/json"}
                })
                const responseJson = await response.json()
                return responseJson
            }
            catch(e){
                alert(e)
            }
    }
    },
    namespaced:true 
}