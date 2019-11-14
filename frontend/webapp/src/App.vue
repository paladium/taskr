<template>
  <div id="app">
    <div class="container">
      <div class="row">
        <div class="column">
          <h4>taskr 📓</h4>
        </div>
        <div class="column-50 column-offset-25">
          <a class="button button-clear">About</a>
        </div>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class App extends Vue {}
</script>

<style lang="css" src="milligram/dist/milligram.min.css"></style>

<style lang="scss">
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin-top: 10px;
}
.button-small{
  font-size: .8rem;
  height: 2.8rem;
  line-height: 2.8rem;
  padding: 0 1.5rem;
}
.button-black{
  background: #000;
  border-color: transparent;
}

.modal{
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: black;
  background-color: rgba(0, 0, 0, 0.4);

  .content{
    background-color: #fefefe;
    margin: 15% auto;
    padding: 20px;
    border-radius: 5px;
    width: 50%;
    z-index: 10;
  }
}
</style>
