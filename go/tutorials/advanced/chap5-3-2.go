package advanced

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	// 将Handler强制转换为了HandlerFunc
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main() {
	// 注意这里是用的http.Handle(调用的是一个Handler)
	// 而不是用之前的http.HandleFunc(调用一个HandlerFunc)
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
}

// http: Handler, HandlerFunc, ServeHTTP

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// 注意, HandlerFunc实现了Handler接口,因此我们只需要实现handlerfunc即可.
// type HandlerFunc func(ResponseWriter, *Request)
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }
