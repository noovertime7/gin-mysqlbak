import { createApp } from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import store from './store'
import router from './router'
const app = createApp(App).use(store).use(router).use(router).use(store).use(Antd);

app.mount('#app');