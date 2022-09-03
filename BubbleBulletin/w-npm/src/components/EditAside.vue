<template>
  <aside id="edit" class="sidebar">
    <span style="float:right; cursor: pointer; margin: 10px; transform: scale(2)" @click="$emit('toggleMode','vanilla')" >×</span>
    <input id="edit-title" type="text" placeholder="标题（不会直接显示）" :value="title" @input="$emit('update:title', $event.target.value)"/>
    <textarea id="edit-content" placeholder="正文，支持Markdown" :value="content" @input="$emit('update:content', $event.target.value)"></textarea>
    <div style="text-align: center">
      <span class="buttons" id="edit-position" @click="$emit('toggleMode','edit-position')">设定位置</span>
      <span class="buttons" id="submit" @click="$emit('submit')">提交</span>
    </div>
  </aside>
</template>

<script>
export default {
  name: 'EditAside',
  props:['title','content'],
  emits:['update:title','update:content','toggleMode'],
}
</script>

<style>
input[type=text] {
  width: 90%;
  padding: 12px 20px;
  margin: 8px 0;
  box-sizing: border-box;
}
#edit-content{
  width: 100%;
  height: 85%;
  padding: 12px 20px;
  margin: 8px 0;
  box-sizing: border-box;
}
.buttons{

  width: 100%;
  text-align: center;
  padding: 12px 20px;
  margin: 10px;
  box-shadow: 1px 1px 2px slategray;
  cursor: pointer;  
  background: grey;

}
</style>