var req = {
	id:"uuid"				,//必须
	describe:"描述" ,
	endpoint:"Local Addr"	//必须
}
fetch('http://127.0.0.1:8000/', {
    body: JSON.stringify(req),
    method: 'POST',
    mode: 'no-cors', // no-cors, cors, *same-origin
    headers: {
        'user-agent': 'Mozilla/4.0 MDN Example',
        'content-type': 'application/x-www-form-urlencoded'
    },
})
    .then(response => response.json());