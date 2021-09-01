import { createApp, h } from "vue";
import vuetify from "./plugins/vuetify";
import VueRouter, {createRouter, createWebHashHistory} from "vue-router";

import App from "./App.vue";
import Homepage from "./pages/Homepage.vue";
import Place from "./pages/Place.vue";

const routes = [
    { path: '/', component: Homepage }, 
    { path: '/place', component: Place }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

const app = createApp(App);

app.use(router)

app.mount("#app");
