package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/jacklin/gores-test/unitl"
	"github.com/zhenghaoz/gorse/client"
	"math/big"
	"strconv"
	"time"
)

var cstZone = time.FixedZone("Asia/shanghai", 8*3600)
var (
	globaLabel     = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	globalCate     = []string{"1", "2", "3", "4", "5", "6"}
	globalFeedback = []string{"star", "like", "download", "share"}
)

func main() {
	// 创建中国时区（东八区，UTC+8）
	//for i := 0; i < 100; i++ {
	//	Doitems()
	//
	//}
	//for i := 0; i < 10; i++ {
	//	Dousers(i)
	//}
	DoFeedback()
}
func Doitems() {
	count := 10
	// Create a client
	gorse := client.NewGorseClient("http://192.168.2.211:8087", "VeiVsM2DR7aMdiNTBlGR")
	ctx := context.TODO()
	var items []client.Item
	for i := 0; i < count; i++ {
		// timestamp
		timestamp := time.Unix(time.Now().Unix(), 0).In(cstZone).Format(time.RFC3339)
		item := client.Item{
			ItemId:     unitl.RandString(10),
			IsHidden:   false,
			Labels:     unitl.RandStringSlice(globaLabel),
			Categories: unitl.RandStringSlice(globalCate),
			Timestamp:  timestamp,
			Comment:    unitl.GenerateRandomSentence(),
		}
		items = append(items, item)
	}
	// Get recommendation.
	fmt.Printf("Items:%v\n", items)
	rowAffected, err := gorse.InsertItems(ctx, items)
	if err != nil {
		fmt.Printf("gorse.InsterItems Error:%s\n", err)
	}
	fmt.Printf("InsterItems:%v\n", rowAffected.RowAffected)
}
func Dousers(startid int) {
	// Create a client
	gorse := client.NewGorseClient("http://192.168.2.211:8087", "VeiVsM2DR7aMdiNTBlGR")
	ctx := context.TODO()
	count := 10
	i := startid * count
	var users []client.User
	for ; i < (startid*count + count); i++ {
		user := client.User{
			UserId:    strconv.Itoa(i),
			Labels:    unitl.RandStringSlice(globaLabel),
			Subscribe: unitl.RandStringSlice(globalCate),
			Comment:   unitl.GenerateRandomSentence(),
		}
		users = append(users, user)
	}
	fmt.Printf("users:%v\n", users)
	// Insert user.
	rowAffected, err := gorse.InsertUsers(ctx, users)
	if err != nil {
		fmt.Printf("gorse.InsertUsers Error%s\n", err)
	}
	fmt.Printf("InsertUsers:%v\n", rowAffected.RowAffected)
}
func DoFeedback() {
	// Create a client
	gorse := client.NewGorseClient("http://192.168.2.211:8087", "VeiVsM2DR7aMdiNTBlGR")
	ctx := context.TODO()
	count := 10
	timestamp := time.Unix(1713864861, 0).In(cstZone).Format(time.RFC3339)

	var fbs []client.Feedback
	for i := 0; i < count; i++ {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(10)))
		val = val.Add(val, big.NewInt(1))
		itemID := getItem("next").Items[int(val.Int64())-1].ItemId
		feedbackType := unitl.RandStringSlice(globalFeedback)
		if len(feedbackType) >= 1 {
			fb := client.Feedback{
				FeedbackType: feedbackType[0],
				UserId:       strconv.Itoa(int(val.Int64())),
				Timestamp:    timestamp,
				ItemId:       itemID,
			}

			fbs = append(fbs, fb)
		}
	}
	fmt.Printf("fbs:%v\n", fbs)
	var rowAffected, err = gorse.InsertFeedback(ctx, fbs)
	if err != nil {
		fmt.Printf("InserFeedback error:%v\n", err)
	}
	fmt.Printf("InserFeedback:%v\n", rowAffected.RowAffected)

}
func getItem(cur string) client.Items {
	// Create a client
	gorse := client.NewGorseClient("http://192.168.2.211:8087", "VeiVsM2DR7aMdiNTBlGR")
	ctx := context.TODO()
	items, err := gorse.GetItems(ctx, cur, 10)
	if err != nil {
		fmt.Printf("getItem error:%v\n", err.Error())
		panic(err)
	}
	return items

}
