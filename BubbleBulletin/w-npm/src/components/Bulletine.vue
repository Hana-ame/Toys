<template>
  <main id="bulletin" :style="{opacity:opacity}">
    <Bubble v-for="bubble in bubbleList" :key="bubble"
      :windowSize="windowSize" 
      :mousePos="mousePos"
      :relativeVector="relativeVector" 
      :scaleRate="scaleRate"
      :bubblePos="bubble.pos"
      :bubbleSize="bubble.size" 
      :bubbleScale="bubble.scale"
      :bubbleBackground="bubble.background"
      :bubbleContent="bubble.content"
      :bubbleMetaData="bubble.metaData"
    />
    <!-- <Bubble 
      :windowSize="windowSize" 
      :mousePos="mousePos"
      :relativeVector="relativeVector" 
      :scaleRate="scaleRate"
      :bubblePos="{x:0,y:0}"
      :bubbleSize="{x:100,y:100}"
      :bubbleScale="1"
      bubbleBackground="yellow"
      bubbleContent="# markdown-it rulezz!"
      bubbleMetaData="Null"
    />
    <Bubble 
      :windowSize="windowSize" 
      :mousePos="mousePos"
      :relativeVector="relativeVector" 
      :scaleRate="scaleRate"
      :bubblePos="{x:-100,y:-100}"
      :bubbleSize="{x:500,y:200}"
      :bubbleScale="0.5"
      bubbleBackground="red"
      bubbleContent="# markdown-it rulezz!
      <p></p>"
      bubbleMetaData="Null" 
    /> -->
  </main>
</template>

<script>
import Bubble from '@/components/Bubble.vue';

export default {  
  name: 'Bulletine',
  components: {
    Bubble
  },
  props: ['windowSize','mousePos','relativeVector','scaleRate','opacity','bubbleList'],
  data () {
    return {
    //   mousePos: {x:0,y:0},
    //   relaVector
    }
  },
  computed: {
  },
  methods: {      
  },
}
</script>

<style>

main{
    height: 100%;
    background: wheat;
}
</style>  