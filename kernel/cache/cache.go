package cache

type Cache struct {
	Adapt string
}

func NewCache(args ...string) *Cache {
	adapt := "file"
	if len(args) > 0 {
		adapt = args[0]
	}
	return &Cache{Adapt: adapt}
}

func (c *Cache) Get(key string, defaults ...interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) Set(key string, value interface{}, ttl int) bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) Delete(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) Clear() bool {
	//TODO implement me
	panic("implement me")
}
