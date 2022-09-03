<template>
  <nav id="nav">
    <span class="bulletine-item" id="bulletine-title">{{title}}</span>

    <span class="bulletine-item" id="bulletine-timeline" @click="$emit('toggleMode','timeline')">时间线</span>
    <span class="bulletine-item" id="bulletine-edit" @click="$emit('toggleMode','edit')">编辑</span>
  </nav>
</template>

<script>
export default {
  name: 'NaviBar',
  props: ['title'],
  emits:['toggleMode']
}
</script>

<style>

#nav{
  position: fixed;
  margin: auto;
  top: 0px;
  left: 10%;
  width: 80%;
  /* height: 32px; */
  z-index: 123456789;
}
#bulletine-title{
  float:left; 
}
#bulletine-edit{
  float: right;
}
#bulletine-timeline{
  float: right;
}
.bulletine-item{
  background: white;  
  border-radius: 5px;
  padding: 12px;
  margin: 10px;
  box-shadow: 1px 1px 2px slategray;
}
.bulletine-item:hover{
  cursor: pointer;  
  box-shadow: 5px 5px 10px slategray;
}

</style>