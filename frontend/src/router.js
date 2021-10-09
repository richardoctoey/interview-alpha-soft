import Vue from "vue";
import Router from "vue-router";
import AppHeader from "./layout/AppHeader";
import AppFooter from "./layout/AppFooter";
import Music from "./views/Football.vue";

Vue.use(Router);

export default new Router({
    linkExactActiveClass: "active",
    routes: [
        {
            path: "/",
            name: "components",
            components: {
                header: AppHeader,
                default: Music,
                footer: AppFooter,
            },
            meta: {
                title: "",
            }
        },
    ],
});
