<template>
  <div class="bubble-meta">
    <img v-if="metaImgSrc" class="bubble-icon" :src="metaImgSrc" alt="头像" />
    <div class="bubble-title">{{metaTitle}}</div>
    <span class="bubble-author">{{metaAuthor}}</span>
    <span class="bubble-timestamp">{{timeFormatted}}</span>
  </div>
</template>

<script>
export default {
  name: 'BubbleMeta',
  props: ['metaImgSrc','metaTitle','metaAuthor','metaTimestamp'],
  
  computed:{
    timeFormatted(){
      return this.metaTimestamp
    }
  }

}
</script>

<style>
.bubble-meta{
  display: none;
  position: absolute;

  top: -50px;
  width: 90%;
  min-width: 300px;
  height: 48px;

  background: white;
  
  border-radius: 25px;

  padding-left: 25px;
  overflow: hidden;
}
.bubble:hover .bubble-meta{  
  display: unset;
  box-shadow: 5px 5px 10px slategray;
}
.bubble-icon{
    width: 48px;
    height: 48px;
    float: left;    
}
.bubble-title{
    /* display: none; */
    font-weight: bolder;
    white-space: nowrap;
}
.bubble-meta span{
    font-size: smaller;
    color: slategray;
    vertical-align: middle;
    margin-left: 5px;
}

</style>