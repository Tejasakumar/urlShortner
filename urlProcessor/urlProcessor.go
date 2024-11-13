package pack

import (
	"math/rand"
	"net/url"
	"strconv"
	"sync"
	"urlShortner/dbmanager"
)


var MyDomain string = "127.0.0.1:8080"

func ProcessURI(resourceLocator string) string {
	shortenedurl := url.URL{
		Scheme:   "http",
        Host:     MyDomain,
        Path:     resourceLocator,
		
	}
	return shortenedurl.String()
}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var symbols = "_-.~"
var dbLock sync.RWMutex

func Shorten(path string) string {
	var lengths []int = []int{6,8,10}
	var index = rand.Intn(3)
	short := getRand(lengths[index])
	dbLock.RLock()
	redirectURL := dbmanager.FindOne(short)
	if redirectURL == ""{
		dbLock.RUnlock()
		dbLock.Lock()
		dbmanager.InsertOne(path,short)
		dbLock.Unlock()
	}else{
		dbLock.RUnlock()
        Shorten(path)
	}
	return short
}

func getRand(n int) string {
	s := ""
	flag := false
	for i := 0; i < n; i++ {
		des := rand.Intn(3)
		if des == 0 {
			s+= strconv.Itoa(rand.Intn(10))
		}else if des == 1&& !flag{
			s += string(symbols[rand.Intn(len(symbols))])
			flag = true
		}else{
			s += string(letters[rand.Intn(len(letters))])
		}
	}
	return s
}

func RetriveURL(path string)string{
	dbLock.RLock()
    defer dbLock.RUnlock()
	return dbmanager.FindOne(path)
    
}


