import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import Argon from "./plugins/argon-kit";
import axios from 'axios';
import VueAxios from 'vue-axios';
import {getBaseUrl} from "@/utils";
import {BootstrapVue} from 'bootstrap-vue';
import EventBus from './EventBus';
import {data as dataEnUs} from './lang/en-us';
import {data as dataIdId} from './lang/id-id';

Vue.config.productionTip = false;
Vue.use(Argon);
Vue.use(VueAxios, axios);
Vue.use(BootstrapVue);
Vue.axios.defaults.withCredentials = false;
Vue.axios.defaults.baseURL = getBaseUrl();
Vue.prototype.$bus = EventBus;

new Vue({
    router,
    render: h => h(App),
    data: {
      "lang": {
          "selectedLang": "en-us",
          "en-us": dataEnUs,
          "id-id": dataIdId,
      },
    },
    methods: {
        T(key) {
            return this.lang[this.lang.selectedLang][key];
        }
    },
}).$mount("#app");
