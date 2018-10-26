import "material-design-icons-iconfont/dist/material-design-icons.css";
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import Vuetify from "vuetify";
import "vuetify/src/stylus/app.styl";

Vue.use(Vuetify, {
  iconfont: "md"
});

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
