package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var l []string

func main() {
	l = []string{
		"https://pximg.moonchan.xyz/img-master/img/2013/07/08/01/50/25/36919122_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/26/00/00/05/90097034_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/26/00/34/29/90098100_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2020/11/06/00/00/05/85485312_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/15/22/20/38/89861476_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/17/20/06/17/89908804_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2020/09/25/20/44/24/84604468_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/28/01/34/38/90139691_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/28/07/51/02/90142930_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/15/20/00/02/89857549_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/16/16/12/37/89879470_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/12/22/37/53/89794654_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2020/12/22/21/25/29/85982941_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/03/28/19/22/02/88766749_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/04/10/49/02/89586986_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/08/23/17/23/89701608_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/23/00/02/29/90023683_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/29/20/12/48/90176688_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/25/21/54/00/90093650_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/29/19/08/51/90175112_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/01/17/00/22/27/87098538_p0_master1200.jpg",
		"https://pximg.moonchan.xyz/img-master/img/2021/05/04/12/42/35/89588689_p0_master1200.jpg",
	}

	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, l[rand.Intn(len(l))], 302)
}
