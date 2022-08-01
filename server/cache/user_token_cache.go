package cache

import (
	"errors"
	"time"

	"github.com/goburrow/cache"
	"github.com/mlogclub/simple/sqls"

	"bbs-go/model"
	"bbs-go/repositories"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
	cache cache.LoadingCache
}

// newUserTokenCache 生成 token 缓存
func newUserTokenCache() *userTokenCache {
	return &userTokenCache{
		cache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				value = repositories.UserTokenRepository.GetByToken(sqls.DB(), key.(string))
				if value == nil {
					e = errors.New("数据不存在")
				}
				return
			},
			cache.WithMaximumSize(1000),
			cache.WithExpireAfterAccess(60*time.Minute), // cache 过期时间
		),
	}
}

// Get 获取用户token
func (c *userTokenCache) Get(token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	val, err := c.cache.Get(token) // Get 返回与 Key 关联的值或调用底层 LoaderFunc 以加载值（如果它不存在）。
	if err != nil {
		return nil
	}
	if val != nil {
		return val.(*model.UserToken)
	}
	return nil
}

func (c *userTokenCache) Invalidate(token string) {
	c.cache.Invalidate(token)
}
