package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Str2StrPointer(value string) *string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return &value
}

func Str2Bool(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return v
}

func Str2StrInt64(value string, acceptZero bool) int64 {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}
	r, err := strconv.ParseInt(value, 10, 64)
	if err != nil || (r == 0 && !acceptZero) {
		return 0
	}
	return r
}

func Int642Int64(r int64, acceptZero bool) *int64 {
	if r == 0 && !acceptZero {
		return nil
	}
	return &r
}

func Str2Float(value string) float64 {
	value = strings.TrimSpace(value)
	if value == "" || value == "#N/A" {
		return 0
	}

	r, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return r
}

var re = regexp.MustCompile(`^(03|05|06|07|08|09|01[2|6|8|9])+([0-9]{8})$`)

func IsPhone(phone string) bool {
	if len(phone) < 8 || len(phone) > 15 {
		return false
	}
	return re.MatchString(phone)
}

func GetPaginationQuery(c *gin.Context, defaultPageSize, minPageSize, maxPageSize int64) (pageOffset, pageSize int64) {
	pageOffsetValue := Str2StrInt64(c.Query("offset"), false)
	if pageOffsetValue == 0 {
		pageOffset = int64(0)
	} else {
		pageOffset = pageOffsetValue
	}

	pageSizeValue := Str2StrInt64(c.Query("size"), false)
	if pageSizeValue == 0 {
		pageSize = defaultPageSize
	} else if pageSizeValue > maxPageSize {
		pageSize = maxPageSize
	} else if pageSizeValue < minPageSize {
		pageSize = minPageSize
	} else {
		pageSize = pageSizeValue
	}
	return pageOffset, pageSize
}

// Format "-created_at +id"
func ParseOrderBy(query string) *string {
	query = strings.TrimSpace(query)
	if len(query) == 0 {
		return nil
	}
	sorts := strings.Split(strings.ToLower(query), " ")
	sortQuery := ""
	for _, sort := range sorts {
		if len(sort) < 2 {
			continue
		}
		if sort[0] == '-' {
			sortQuery = fmt.Sprintf("%s %s %s,", sortQuery, sort[1:], "DESC")
		} else if sort[0] == '+' {
			sortQuery = fmt.Sprintf("%s %s %s,", sortQuery, sort[1:], "ASC")
		}
	}
	if len(sortQuery) < 2 {
		return nil
	}
	sortQuery = sortQuery[1 : len(sortQuery)-1]
	return &sortQuery
}

func StrPointer2Str(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func TimePointer2Str(t *time.Time, format string) string {
	if t == nil {
		return ""
	}

	return t.Format(format)
}

func ShuffleArray(array []map[string]interface{}) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})
}

func ShuffleArrayInt64(arr []int64) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func GetRandomItemsInt64(arr []int64, n int) []int64 {
	// Kiểm tra số lượng phần tử trong mảng
	if len(arr) <= n {
		// Nếu số lượng phần tử nhỏ hơn hoặc bằng 10, trả về toàn bộ mảng
		return arr
	}

	// Trộn lẫn thứ tự của mảng mỗi khi gọi hàm
	ShuffleArrayInt64(arr)

	// Trả về 10 phần tử đầu tiên
	return arr[:n]
}

func TrimSpaces(arr []string) []string {
	for i, str := range arr {
		arr[i] = strings.TrimSpace(str)
	}
	return arr
}
