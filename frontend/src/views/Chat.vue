<template>
  <section class="terminal">
    <header class="header">
      <button class="header-btn">[Create Chat]</button>
      <button class="header-btn">[Connect To Chat]</button>
      <button class="header-btn">[Settings]</button>
    </header>
    <div v-if="isVueChat" class="content">
      <div class="window-chat">
        <div class="answer-mess"></div>
        <div class="mess"></div>
      </div>
      <input type="text" class="input-mess">
      <button class="btn">Send message</button>
    </div>
  </section>
</template>

<script>
const isVueChat = false;
</script>

<style>
.terminal{
  grid-column: 1 / span 12;
  height: 100vh;
  color: green;
  padding: 2em 1em;
}
.header{
  display: flex;
}
.header-btn{
  color: #009900;
  border: none;
  background: transparent;
  cursor: pointer;
  font-weight: 500;
  font-size: 1.8em;
  padding: .1rem .1rem;
  font-family: "Source Code Pro", monospace;
  margin-right: 1rem;
}

.window-chat{
  width: 80%;
  height: 60%;
  background: #ccc;
}
.input-mess{
  border: none;
  border-bottom: 1px solid white;
  outline: none;
  color: #009900;
}
.btn{
  color: #009900;
  border: none;
  background: transparent;
  cursor: pointer;
}
</style>
