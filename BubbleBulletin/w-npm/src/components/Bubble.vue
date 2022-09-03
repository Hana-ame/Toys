<template>
  <div class="bubble"
    :style="{
        transform: 'scale('+scaleRate*bubbleScale+')',
        left: displayPos.x+'px',
        top: displayPos.y+'px',
        width: bubbleSize.x+'px',
        height: bubbleSize.y+'px',
        background: bubbleBackground,
    }">
    <BubbleMeta
      metaImgSrc="https://9191.ga/imgs/2021/11/c78c57332a461da4.webp"
      :metaTitle="bubbleMetaData.title"
      :metaAuthor="bubbleMetaData.author"
      :metaTimestamp="bubbleMetaData.timestamp"
    />
    <div class="bubble-content" v-html="contentRendered"></div>
  </div>
</template>

<script>
import BubbleMeta from '@/components/BubbleMeta.vue';
var md = require('markdown-it')();


export default {
  name: 'Bubble',
  components: {
    BubbleMeta
  },
  props: ['windowSize','mousePos','relativeVector','scaleRate','bubblePos','bubbleSize','bubbleScale','bubbleBackground','bubbleContent','bubbleMetaData'],
  // props: ['windowSize','mousePos','relativeVector','scaleRate','bubblePos','bubbleSize','bubbleScale','bubbleMeta','bubbleContent'],
  data () {
    return {
        
    }
  },
  computed: {
    displayPos(){
      return{
        // 窗口中心点 + 自身坐标 + 相对坐标修正 + 放缩修正
        x: this.windowSize.x/2 + (this.bubblePos.x + this.relativeVector.x) * this.scaleRate + ( this.bubbleSize.x * (this.scaleRate*this.bubbleScale -1) /2 ),
        y: this.windowSize.y/2 + (this.bubblePos.y + this.relativeVector.y) * this.scaleRate + ( this.bubbleSize.y * (this.scaleRate*this.bubbleScale -1) /2 ),
      }
    },
    contentRendered(){
      return md.render(this.bubbleContent)
    }
  },
  methods: {      
  },
}
</script>

<style>
.bubble{
  position: fixed;  
  border-radius: 5px;
  padding: 5px;
  /* overflow: hidden; */
}
.bubble:hover{
  box-shadow: 5px 5px 10px slategray;
  z-index: 12345;
}
.test{
    position: fixed;
    width: 100px;
    height: 100px;
    background: yellow;
}


.bubble-content{  
    /* white-space: nowrap; */
    height: 100%;
    width: 100%;
    overflow: hidden;
    overflow-wrap: break-word;
    /* overflow-y: scroll; */
}
img{
    max-width: 100%;
}
p{
    margin-block-start: auto;
    margin-block-end: auto;
}
h1{
    margin-block-start: auto;
    margin-block-end: auto;
}
h2{
    margin-block-start: auto;
    margin-block-end: auto;
}
h3{
    margin-block-start: auto;
    margin-block-end: auto;
}
h4{
    margin-block-start: auto;
    margin-block-end: auto;
}
h5{
    margin-block-start: auto;
    margin-block-end: auto;
}
h6{
    margin-block-start: auto;
    margin-block-end: auto;
}
</style>