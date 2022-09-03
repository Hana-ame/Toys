<template>
  <div id="wrapper" v-on="{mousemove: mouseMove,  mousedown: mouseDown, mouseup: mouseUp}">
    <Bulletine 
      :windowSize="windowSize" 
      :mousePos="mousePos"
      :relativeVector="calcVector" 
      :scaleRate="scaleRate"
      :opacity="opacity"
      :bubbleList="bubbleListReversed"
      />
    <NaviBar
      :title="title"
      @toggle-mode="mode=($event)"
      />
    <EditAside v-if="mode=='edit'||mode=='edit-position'||mode=='edit-position2'||mode=='edit-position-done'" 
      @toggle-mode="mode=($event)"
      v-model:title="bubblePost.bubbleMetaData.title"
      v-model:content="bubblePost.content"
      @submit="submit"
      />
    <span v-else-if="mode=='timeline'" id="timeline" class="sidebar">
      timeline
    </span>

    <Bubble v-if="mode=='edit'||mode=='edit-position'||mode=='edit-position2'||mode=='edit-position-done'"
      :windowSize="{x:0,y:0}" 
      :mousePos="mousePos"
      :relativeVector="{x:0,y:0}" 
      :scaleRate="1"
      :bubblePos="bubblePost.displayPos"
      :bubbleSize="bubblePost.bubbleSize"
      :bubbleScale="1"
      :bubbleBackground="bubblePost.background"
      :bubbleContent="bubblePost.content"
      :bubbleMetaData="bubblePost.bubbleMetaData"
    />
  </div>

  <div v-if="debug" id="scoper">
    <p>windowSize:{{windowSize}}</p>
    <p>mousePos:{{mousePos}}</p>
    <p>relativeVector:{{relativeVector}}</p>
    <p>scaleRate:{{scaleRate}}</p>
    <p>bubblePost:{{bubblePost}}</p>
    <p>truePos:{{truePos}}</p>
  </div>
</template>

<script>
import $ from "jquery";
import Bulletine from '@/components/Bulletine.vue';
import NaviBar from '@/components/NaviBar.vue';
import EditAside from '@/components/EditAside.vue';
// import Timeline from '@/components/Timeline.vue';
import Bubble from '@/components/Bubble.vue';
var md = require('markdown-it')();

export default {
  name: 'App',

  components: {
    Bulletine,
    NaviBar,
    EditAside,
    // Timeline,
    Bubble,
  },

  data() {
    return {
      title:"学习不爱我",
      debug: false,
      mode: "vanilla",

      windowSize: {x:0,y:0},
      mousePos : {x:0,y:0},
      relativeVector : {x:0,y:0,z:0},
      mouseLast : {x:0,y:0},
      isMouseDown: false,

      bubblePost : {
        bubbleMetaData: {
          title: "",
          author: "",
        },
        displayPos: {x:0,y:0},
        bubbleSize: {x:0,y:0},
        background: "rgba(255,255,255,0.5)",
        content: "",
      },
      bubbleList : [],
    }
  },

  computed: {
    bubbleListReversed(){
      var t = [...this.bubbleList]
      return t.reverse()
    },
    scaleRate(){
      return Math.pow(2, -this.relativeVector.z / 2000).toFixed(2)
    },
    calcVector(){
      if (this.isMouseDown){
        return{
          x: this.relativeVector.x + (this.mousePos.x - this.mouseLast.x) / this.scaleRate,
          y: this.relativeVector.y + (this.mousePos.y - this.mouseLast.y) / this.scaleRate,
          z: this.relativeVector.z,
        }
      }else{
        return this.relativeVector
      }
    },
    opacity(){
      if (this.mode != "edit-position" && this.mode!='edit-position2'){
        return 1
      }else{
        return 0.4  
      }
    },
    truePos(){
      return {
        x: (this.mousePos.x - this.windowSize.x/2) / this.scaleRate - this.relativeVector.x,
        y: (this.mousePos.y - this.windowSize.y/2) / this.scaleRate - this.relativeVector.y,
      }
    },
    contentRendered(){
      return md.render(this.bubblePost.content)
    }
  },
  methods: {  
    log(event){
      console.log(event)
    },

    submit(){
      var postJson = {
        pos:{
          x: (this.bubblePost.displayPos.x - this.windowSize.x/2) / this.scaleRate - this.relativeVector.x,
          y: (this.bubblePost.displayPos.y - this.windowSize.y/2) / this.scaleRate - this.relativeVector.y,
        },
        size:{
          x: this.bubblePost.bubbleSize.x,
          y: this.bubblePost.bubbleSize.y,
        },
        scale:1/this.scaleRate,
        background:this.bubblePost.background,
        content:this.bubblePost.content,
        metaData:{
          imgSrc:"",
          title:this.bubblePost.bubbleMetaData.title,
        }
      }

      // this.bubbleList.push(postJson)
      fetch('/bb/api/bubble', {
          body: JSON.stringify(postJson),
          method: 'POST',
          // mode: 'no-cors', // no-cors, cors, *same-origin
          headers: {
              // 'user-agent': 'Mozilla/4.0 MDN Example',
              'content-type': 'application/x-www-form-urlencoded'
          },
      })
      .then(response => response.json())
      .then(data => this.bubbleList=data)




      this.bubblePost = {
        bubbleMetaData: {
          title: "",
          author: "",
        },
        displayPos: {x:0,y:0},
        bubbleSize: {x:0,y:0},
        background: "",
        content: "",
      }
      return postJson
    },


    onResize(){
      this.windowSize.x = $(window).width()
      this.windowSize.y = $(window).height()
    },
    mouseMove(event){
      this.mousePos.x = event.x
      this.mousePos.y = event.y
      if (this.mode=='edit-position2'){
        this.bubblePost.bubbleSize.x = Math.abs( this.bubblePost.displayPos.x - event.x)
        this.bubblePost.bubbleSize.y = Math.abs( this.bubblePost.displayPos.y - event.y)
      }
    },
    onWheel(event){
      this.relativeVector.z += event.deltaY
      if (this.relativeVector.z > 10000) {
        this.relativeVector.z = 10000
      } else if (this.relativeVector.z < -10000) {
        this.relativeVector.z = -10000
      }
    },
    mouseDown(event){
      this.isMouseDown = true
      this.mouseLast.x = event.x
      this.mouseLast.y = event.y  
      // console.log(this.mode)
      if (this.mode == 'edit-position'){
        this.isMouseDown = false
        this.bubblePost.displayPos.x = event.x
        this.bubblePost.displayPos.y = event.y
        this.bubblePost.bubbleSize.x = 0
        this.bubblePost.bubbleSize.y = 0
        this.bubblePost.background = ""
        this.mode = 'edit-position2'        
      } 
    },
    mouseUp(event){
      this.isMouseDown = false
      if (this.mode!='edit-position2'){
        this.relativeVector.x += (event.x - this.mouseLast.x) / this.scaleRate
        this.relativeVector.y += (event.y - this.mouseLast.y) / this.scaleRate
      }
      if (this.mode=='edit-position2'){
        // console.log(this.mode)
        // 之后再做，只能先左上角
        this.bubblePost.bubbleSize.x = Math.abs( this.bubblePost.displayPos.x - event.x)
        this.bubblePost.bubbleSize.y = Math.abs( this.bubblePost.displayPos.y - event.y)
        this.bubblePost.background = "rgba(0,0,0,0.3)"
        this.mode = 'edit-position-done'        
      }
    }
  },

  created() {
    window.addEventListener("resize", this.onResize);
    window.addEventListener("mousewheel", this.onWheel);
  },
  mounted() {
    this.onResize()
    fetch('/bb/api/bubble', {
      // body: JSON.stringify(postJson),
      method: 'GET',
      // mode: 'no-cors', // no-cors, cors, *same-origin
      headers: {
        // 'user-agent': 'Mozilla/4.0 MDN Example',
        'content-type': 'application/x-www-form-urlencoded'
      },
    })
    .then(response => response.json())
    .then(data => this.bubbleList=data)
  },
  unmounted() {
    window.removeEventListener("resize", this.onResize);
    window.removeEventListener("mousewheel", this.onWheel);
  },
}

</script>

<style>

.sidebar{
  position: fixed;
  right: 45px;
  top: 45px;
  
  width: 500px;
  height: 80%;
  background: white;
  
  border-radius: 25px;
  box-shadow: 5px 5px 10px slategray;
  padding: 25px;

  z-index: 123456;
}




html{
  height: 100%;
}
body{
  height: 100%;
  margin: 0px;
}
#wrapper{
  height: 100%;
}


#scoper  {
  position: fixed;
  top: 200px;
  right: 0px;
}



#app {
  height: 100%;
  margin-top: 0px;
}
</style>
