<template>
  <v-app id="app">
    <v-toolbar color="indigo" dark fixed app>
      <v-toolbar-title>MtG deck tools</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid grid-list-md>
        <router-view/>
      </v-container>
    </v-content>
  </v-app>
</template>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
