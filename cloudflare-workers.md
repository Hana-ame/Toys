```js
const destinationURL = "https://raw.githubusercontent.com"
const statusCode = 301

addEventListener("fetch",  event => { 
  return event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  const urlStr = request.url  // request的域名
  const urlObj = new URL(urlStr) // 同样是域名
  const path = urlObj.href.substr(urlObj.origin.length) // 提取请求域名中的path
  // console.log(path)

  // 不要http!
  if (urlObj.protocol == `http:`){
    urlObj.protocol = 'https:'
    return Response.redirect(urlObj.href, statusCode)
  }

  // 抓
  const res = await fetch(
    destinationURL+path,
    {
      method: request.method,
      headers: request.headers,
      redirect: 'manual',
    }
  )

  // 换头
  newHead = new Headers(res.headers)
  newHead.delete('access-control-allow-origin')
  newHead.set('access-control-allow-origin', '*')
  newHead.delete('content-security-policy')
  newHead.delete('content-security-policy-report-only')
  newHead.delete('clear-site-data')

  newHead.delete(`content-type`)
  newHead.set(`content-type`,`application/javascript`)

  return new Response(
    res.body,
    {
      status: res.status,
      headers: newHead
    }
  )
}
```

```js
addEventListener("fetch", (event) => {
  event.respondWith(
    handleRequest(event.request).catch(
      (err) => new Response(err.stack, { status: 500 })
    )
  );
});

/**
 * Many more examples available at:
 *   https://developers.cloudflare.com/workers/examples
 * @param {Request} request
 * @returns {Promise<Response>}
 */
async function handleRequest(request) {
  a = request.headers.get("CF-Connecting-IP")
  return new Response(a);
}
```


```js
return Response.redirect(`https://google.com/`, 302)
```

```js
addEventListener('fetch', event => {
    event.respondWith(handleRequest(event.request))
  })
  
  /**
   * Respond to the request
   * @param {Request} request
   */
  async function handleRequest(request) {
  
    const urlStr = request.url  // request的域名
    const urlObj = new URL(urlStr) // 同样是域名
    const path = urlObj.href.substr(urlObj.origin.length) // 提取请求域名中的path
  
    let resHead = new Headers(request.headers)
    resHead.delete('Referer')
    resHead.set('Referer', 'https://pixiv.net/')
    // 抓取目标域名内容
    const res = await fetch(
      "https://i.pximg.net"+path, // 注意要加入path
      {
        method: request.method,
        headers: resHead,
        redirect: 'manual',
      }
    )
  
    // 好像没用到，管他呢
    let newHead = new Headers(res.headers)
    newHead.set('access-control-allow-origin', '*')
    newHead.delete('content-security-policy')
    newHead.delete('content-security-policy-report-only')
    newHead.delete('clear-site-data')
  
    // 返回获得的页面
    return new Response(
      res.body, 
      {
        status: res.status,
        headers: newHead
      }
    )
    
  }
  ```