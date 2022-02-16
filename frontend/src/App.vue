<template>
<!--  <div id="nav">-->
<!--    <router-link to="/">Chat</router-link> |-->
<!--    <router-link to="/about">About</router-link>-->
<!--  </div>-->
  <div class="container">
    <router-view/>
  </div>
</template>

<style>
@font-face {
  font-family: 'Source Code Pro', monospace;
  src: url('./assets/fonts/SourceCodePro-Regular.ttf') format('woff');
  font-style: normal;
  font-weight: 400;
}
@font-face {
  font-family: 'Source Code Pro', monospace;
  src: url('./assets/fonts/SourceCodePro-Medium.ttf') format('woff');
  font-style: normal;
  font-weight: 500;
}
@font-face {
  font-family: 'Source Code Pro', monospace;
  src: url('./assets/fonts/SourceCodePro-SemiBold.ttf') format('woff');
  font-style: normal;
  font-weight: 600;
}
*{
  background: #0f0f23;
  box-sizing: border-box;
  margin: 0;
}

body{
  font-size: 62.5%;
}
#app {
  font-family: "Source Code Pro", monospace;
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
.container{
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-column-gap: 2.4rem;
}
</style>
