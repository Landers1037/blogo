/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package middleware

import (
	"encoding/json"
	"errors"
	"github.com/landers1037/blogo/config"
	"github.com/landers1037/blogo/models/article"
	"github.com/landers1037/blogo/models/dao/post_dao"
	"github.com/landers1037/blogo/models/response"
	"github.com/landers1037/blogo/rediscache"
)

var (
	useOrNo bool
)

// Cache 使用前置redis缓存
func Cache(name string)  interface{}{
	getFlag()
	if useOrNo {
		hit,e := getFromRedis(name)
		if e !=nil{
			//把内容缓存到cache
			hitDB := getFromDB(name)
			_, _ = rediscache.Set(name, hitDB, config.Cfg.Expires)
			return getFromDB(name)
		}
		var hitCache article.DB_BLOG_POST
		_ = json.Unmarshal(hit, &hitCache)
		return hitCache
	}else {
		return getFromDB(name)
	}
}

// PostCache 文章列表的缓存
func PostCache(p int) ([]response.RES_POST, int){
	getFlag()
	if useOrNo {
		_ = rediscache.Setup()
		hit :=  rediscache.Exists("allposts")
		if hit{
			var s []response.RES_POST
			posts,_ := rediscache.Get("allposts")
			_ = json.Unmarshal(posts, &s)
			var length int = len(s)
			res := pageNation(p, s)
			// 击中从缓存读取
			return res,length
		}else {
			var posts []response.RES_POST
			posts, _ = post_dao.PostQueryAll(map[string]interface{}{})
			var length int = len(posts)
			res := pageNation(p, posts)
			_, _ = rediscache.Set("allposts", posts, config.Cfg.PostsTimeout)
			// 未击中缓存更新
			return res,length
		}

	}else {
		//redis未开启
		var posts []response.RES_POST
		posts, _ = post_dao.PostQueryAll(map[string]interface{}{})
		res := pageNation(p, posts)
		var length int = len(posts)
		return res,length
	}

}

// CheckCache 检查是否命中缓存
func CheckCache(name string)  string{
	getFlag()
	if useOrNo {
		_,e := getFromRedis(name)
		if e !=nil{
			return "miss"
		}

		return "hit"
	}else {
		return "fromdb"
	}
}

func getFlag()  {
	useOrNo = config.Cfg.UseRedis
}

// getFromDB 直接从数据库读取
func getFromDB(name string) article.DB_BLOG_POST{
	var dbFetch article.DB_BLOG_POST
	dbFetch, _ = post_dao.PostQuery(map[string]interface{}{"name": name})
	return dbFetch
}

// getFromRedis 从redis读取
func getFromRedis(name string) ([]byte,error) {
	// init
	initErr := rediscache.Setup()
	if initErr != nil{
		return []byte(""),errors.New("initerr")
	}
	exit := rediscache.Exists(name)
	if exit{
		cacheFetch,err := rediscache.Get(name)
		if err != nil{
			return []byte(""),errors.New("nohit")
		}else {
			return cacheFetch,nil
		}
	}else {
		return []byte(""),errors.New("nohit")
	}

}

// pageNation 文章列表的分页
func pageNation(p int, data []response.RES_POST) []response.RES_POST {
	// 分页数据从ini文件里读取
	// 对于最后一页不满足分页长度的 需要进行截取
	page := config.Cfg.PageSize
	var from, end int

	if p<=0 {
		return data
	}else {
		from = (p-1)*page
		end = (p-1)*page+page
		if end >= len(data) {
			data = data[from:]
			return data
		}
		data = data[from:end]
		return data
	}
}