import { createApp } from 'vue'
import  App  from './App.vue'
import router from './router/index'
import store from './store'
import UI from './components/UI'

const app = createApp(App)

UI.forEach(el=>{
    app.component(el.name,el)
})

app.use(store)
app.use(router)
app.mount('#app')
