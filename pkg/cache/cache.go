package cache

type CacheEl struct {
	Token    string
	FullName string
	Email    string
	Pass     string
}

type Cache map[string]CacheEl

func NewCache() *Cache {
	b := make(map[string]CacheEl)
	c := Cache(b)
	return &c
}

func (C *Cache) Add(token, fullName, email, pass string) {
	(*C)[token] = CacheEl{Token: token, FullName: fullName, Email: email, Pass: pass}
}

func (C *Cache) Get(token string) CacheEl {
	return (*C)[token]
}

func (C *Cache) Del(token string) {
	delete(*C, token)
}
