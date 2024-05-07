package unitl

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"strings"
)

func RandString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[val.Int64()]
	}
	return string(result)
}
func RandStringSlice(label []string) []string {
	// 随机获取切片中的5个不重复的元素
	pickedIndices := make(map[int]bool)
	var picks []string
	maxLen := mrand.Intn(len(label))
	for len(picks)-1 < maxLen {
		// 生成一个随机索引
		index := mrand.Intn(len(label))
		// 检查是否已经选过这个索引
		if _, exists := pickedIndices[index]; !exists {
			pickedIndices[index] = true
			picks = append(picks, label[index])
		}
	}
	return picks // 输出随机选取的元素
}

// GenerateRandomSentence 生成一个随机句子
func GenerateRandomSentence() string {
	var (
		// 定义一些单词列表
		nouns      = []string{"cat", "dog", "book", "computer", "house", "car"}
		verbs      = []string{"runs", "jumps", "reads", "drives", "sleeps", "flies"}
		adjectives = []string{"happy", "sad", "fast", "slow", "big", "small"}
		articles   = []string{"the", "a", "an"}
	)
	// 随机选择一个名词、动词、形容词和冠词
	nounIndex := mrand.Intn(len(nouns))
	verbIndex := mrand.Intn(len(verbs))
	adjectiveIndex := mrand.Intn(len(adjectives))
	articleIndex := mrand.Intn(len(articles))

	// 随机决定是否使用形容词
	useAdjective := mrand.Intn(2) == 0

	// 组合成句子
	var sentence []string
	sentence = append(sentence, articles[articleIndex])
	if useAdjective {
		sentence = append(sentence, adjectives[adjectiveIndex])
	}
	sentence = append(sentence, nouns[nounIndex])
	sentence = append(sentence, verbs[verbIndex])
	sentence = append(sentence, ".")

	// 使用strings.Join将单词组合成句子
	return strings.Join(sentence, " ")
}
