<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>


<style>
  .tool-bar{
    border-bottom:1px solid #ebeef5;
    padding:5px;
  }
  .tool-bar dev {
    margin-right: 10px;
  }
  .tool-bar .left {
    float:left;  
  }
  .tool-bar .right {
    float:right;
  }
</style>