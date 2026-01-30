import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './route'
import store from './store'
import Vant from 'vant'
import 'vant/lib/index.css'
import { Button, Field, Cell, IndexBar, NavBar, Search, Grid, GridItem, List, Checkbox, Icon } from 'vant';

const app = createApp(App)

app.use(Vant)
app.use(store)
app.use(router)

// 按需注册需要的组件
app.use(Button);
app.use(Field);
app.use(Cell);
app.use(IndexBar);
app.use(NavBar);
app.use(Search);
app.use(Grid);
app.use(GridItem);
app.use(List);
app.use(Checkbox);
app.use(Icon);

app.mount('#app')
