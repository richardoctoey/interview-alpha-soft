<template>
  <header class="header-global">
    <base-nav class="navbar-main" transparent type="" effect="light" expand>
      <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">

      </router-link>

      <div class="row" slot="content-header" slot-scope="{closeMenu}">
        <div class="col-6 collapse-brand">
          <router-link to="/">
            <img src="img/brand/logo_purple.svg">
          </router-link>
        </div>
        <div class="col-6 collapse-close">
          <close-button @click="closeMenu"></close-button>
        </div>
      </div>
    </base-nav>
  </header>
</template>
<script>
import BaseNav from "@/components/BaseNav";
import BaseDropdown from "@/components/BaseDropdown";
import CloseButton from "@/components/CloseButton";
import {getCookie, isLoggedIn, logout} from "@/utils";

export default {
  components: {
    BaseNav,
    CloseButton,
    BaseDropdown
  },
  data: () => ({
    isLoggedTrueFalse: isLoggedIn(),
  }),
  mounted() {
    this.$bus.$on('logged', () => {
      this.isLoggedTrueFalse = isLoggedIn();
    });
  },
  methods: {
    loggedout() {
      this.isLoggedTrueFalse = false;
      this.axios.post('user/logout').then((response)=>{}).catch((err) => {});
      logout();
    },
    isLogged(typ) {
      if(typ==="logout" && this.isLoggedTrueFalse) {
        return true;
      }
      if(typ==="login" && !this.isLoggedTrueFalse) {
        return true;
      }
      return false;
    }
  },
};
</script>
<style>
</style>
