l.requests.makej("GET", "/api/verif", {
	query: {
	  token : encodeURIComponent(sessionStorage.getItem("token")),
	  login : encodeURIComponent(sessionStorage.getItem("login"))
	}
  }).then(data => {
	  if (data.status != 1 && data.status != 2 && data.status != 3) {
		window.location.href="/index.html"
	  }
  })