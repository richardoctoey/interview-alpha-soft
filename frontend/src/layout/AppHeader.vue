<template>
  <header class="header-global">
    <base-nav class="navbar-main" transparent type="" effect="light" expand>
      <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
        <img src="img/brand/logo.svg" alt="logo">
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

      <ul class="navbar-nav navbar-nav-hover align-items-lg-center">
        <router-link class="nav-link" to="/football" v-show="isLoggedTrueFalse === false">{{this.$root.T("navurl1")}}</router-link>
        <router-link class="nav-link" to="/football" v-show="isLoggedTrueFalse === true">{{this.$root.T("navurl1")}}</router-link>
        <router-link class="nav-link" to="/history-lotre" v-show="isLoggedTrueFalse === true">{{this.$root.T("navurl2")}}</router-link>
        <router-link class="nav-link" to="/withdraw" v-show="isLoggedTrueFalse === true">{{this.$root.T("navurl3")}}</router-link>
        <router-link class="nav-link" to="/topup" v-show="isLoggedTrueFalse === true">{{this.$root.T("navurl4")}}</router-link>
      </ul>
      <ul class="navbar-nav align-items-lg-center ml-lg-auto">
        <router-link class="btn btn-neutral btn-icon" to="/login" v-show="isLogged('login')">
            <span class="btn-inner--icon">
              <i class="fa fa-sign-in mr-2"></i>
            </span>
          <span class="nav-link-inner--text">{{this.$root.T("navurl5")}}</span>
        </router-link>
        <b-button class="btn btn-neutral btn-icon" v-on:click="loggedout()" v-show="isLogged('logout')">
            <span class="btn-inner--icon">
              <i class="fa fa-sign-in mr-2"></i>
            </span>
          <span class="nav-link-inner--text">{{this.$root.T("navurl6")}}</span>
        </b-button>
      </ul>
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
