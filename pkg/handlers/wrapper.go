package handlers

/*func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}*/
