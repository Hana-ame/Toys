
var catchEvent

var md = window.markdownit();
md.use(markdownitContainer, 'spoiler', {

    validate: function (params) {
        return params.trim().match(/^spoiler\s+(.*)$/);
    },

    render: function (tokens, idx) {
        var m = tokens[idx].info.trim().match(/^spoiler\s+(.*)$/);

        if (tokens[idx].nesting === 1) {
            // opening tag
            return '<details><summary>' + md.utils.escapeHtml(m[1]) + '</summary>\n';

        } else {
            // closing tag
            return '</details>\n';
        }
    }
});
md.use(markdownitContainer, 'warning', {

    validate: function (params) {
        return params.trim().match(/^warning$/);
    },

    render: function (tokens, idx) {
        var m = tokens[idx].info.trim().match(/^warning$/);

        if (tokens[idx].nesting === 1) {
            // opening tag
            return '<div class="warning" style="padding: 1px;">';

        } else {
            // closing tag
            return '</div>\n';
        }
    }
});

console.log(md.render('::: warning \n*here be dragons*\n:::'));
console.log(md.render("asd\n\n![](https://twimg.moonchan.xyz/media/E8gxdp4UUAYgWo6?format=png&name=360x360)\n<>\n>\n\n::: spoiler sb\n\nsd\"\"ffs\n\n:::\n"))


const b2Main = {
    data() {
        return {
            counter: 0,
            width: 0,
            height: 0,

            // 拖动组 (done)

            // 全局坐标
            globalPos: {
                x: 0, y: 0, z: 0
            },
            //当前鼠标坐标
            mousePos: {
                x: 0, y: 0
            },
            //鼠标按下时坐标
            mousePosPre: {
                x: 0, y: 0
            },
            mousePressed: false,
            // 拖动组end

            content: "12312321\n# 111\nd`<p></p>`dfs\ndfdsf![](https://twimg.moonchan.xyz/media/E8gxdp4UUAYgWo6?format=png&name=360x360)\n\n123\n\n\n<span style=\"color:blue\">some *blue* text</span>\n\n::: spoiler sb\n\nsdffs\n\n:::\n\n::: warning\nsd\n:::"
        }
    },
    computed: {
        // 拖动组 (done)
        referPos() {
            if (this.mousePressed) {
                return {
                    x: this.globalPos.x + (this.width / 2) + this.mousePos.x - this.mousePosPre.x,
                    y: this.globalPos.y + (this.height / 2) + this.mousePos.y - this.mousePosPre.y,
                    z: this.globalPos.z
                }
            }
            return {
                x: this.globalPos.x + (this.width / 2),
                y: this.globalPos.y + (this.height / 2),
                z: this.globalPos.z
            }
        },
        // 暂定
        transCoe() {
            return Math.pow(2, -this.globalPos.z / 2000).toFixed(1)
        },
        translatePos() {
            if (this.mousePressed) {
                return {
                    x: (this.width / 2) + (this.globalPos.x) * this.transCoe + this.mousePos.x - this.mousePosPre.x,
                    y: (this.height / 2) + (this.globalPos.y) * this.transCoe + this.mousePos.y - this.mousePosPre.y,
                    // z: this.globalPos.z
                }
            }
            return {
                x: (this.width / 2) + this.globalPos.x * this.transCoe,
                y: (this.height / 2) + this.globalPos.y * this.transCoe,
                // z: this.globalPos.z
            }
        },
        // 拖动组end
        contentRendered() {
            return md.render(this.content)
        },
    },
    created() {
        window.addEventListener("resize", this.onResize);
        window.addEventListener("mousewheel", this.onWheel);
    },
    mounted() {
        this.onResize()
    },
    unmounted() {
        window.removeEventListener("resize", this.onResize);
        window.removeEventListener("mousewheel", this.onWheel);
    },
    methods: {
        // 当窗口大小变动时，得到窗口大小
        onResize() {
            this.width = $(window).width()
            this.height = $(window).height()
        },
        // 滚动组
        onWheel(event) {
            catchEvent = event
            this.globalPos.z += event.deltaY
            if (this.globalPos.z > 5000) {
                this.globalPos.z = 5000
            } else if (this.globalPos.z < -5000) {
                this.globalPos.z = -5000
            }
        },
        //// 滚动组end
        onDrag() {
            return
        },
        // 拖动组 (done)
        mouseMove(event) {
            // catchEvent = event
            if (this.mousePressed) {
                this.mousePos.x = event.clientX;
                this.mousePos.y = event.clientY;
            }
        },
        mouseDown(event) {
            if (!this.mousePressed) {
                this.mousePressed = true
                this.mousePosPre.x = event.clientX;
                this.mousePos.x = event.clientX;
                this.mousePosPre.y = event.clientY;
                this.mousePos.y = event.clientY;
            }
        },
        mouseUp(event) {
            this.mousePressed = false
            this.globalPos.x += (this.mousePos.x - this.mousePosPre.x) / this.transCoe
            this.globalPos.y += (this.mousePos.y - this.mousePosPre.y) / this.transCoe
        },
        //// 拖动组end




    },
    components: {
        // contextMenu 
    },
    filters: {
        // removed
    }
}



const box = {
    data() {
        return {
            height: 0,
            width: 0,
            text: 0,
        }
    },
    methods: {
        a() {
            console.log("box")
            this.text += 1
        }
    }
}

Vue.createApp(box).mount('#box')




const app = Vue.createApp(b2Main)

app.component('button-counter', {
    data() {
        return {
            count: 0
        }
    },
    template: `
      <button v-on:click="count++">
        You clicked me {{ count }} times.
      </button>`
})




app.mount('body')






////////////////////////////////////////////////////
//              end of the file                   //
////////////////////////////////////////////////////

        // debug
        // a(){
        //     console.log("a")
        //     this.counter += 1
        // }




        // setInterval(() => {
        //     this.counter++
        // }, 100)


// window.onresize = function(){
//     console.log($(window).width())
//     console.log($(window).height())

//     width = $(window).width()
//     height = $(window).height()
// }()