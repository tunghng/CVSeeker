package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
)

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func KindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// Struct2JSON convert interface to json string
func Struct2JSON(o interface{}) string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(b)
}

// StringInSlice - StringInSlice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringInMapKey - StringInMapKey
func StringInMapKey(a string, m map[string]interface{}) bool {
	for k := range m {
		if k == a {
			return true
		}
	}
	return false
}

// StringInMapValue - StringInMapValue
func StringInMapValue(a string, m map[string]string) bool {
	for _, v := range m {
		if v == a {
			return true
		}
	}
	return false
}

func RemoveAllUnCharacters(source string) (string, error) {
	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(source, ""), nil
}

func RemoveAllUnCharactersV2(source string) string {
	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(source, "")
}

// StringInSlice StringInSlice
func Int64InSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// BoolToInt64 BoolToInt64
func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// ContainsStructFieldValueContainsStructFieldValue
func ContainsStructFieldValue(slice interface{}, fieldName string, fieldValueToCheck interface{}) bool {
	rangeOnMe := reflect.ValueOf(slice)

	for i := 0; i < rangeOnMe.Len(); i++ {
		s := rangeOnMe.Index(i)
		f := s.FieldByName(fieldName)
		if f.IsValid() {
			if f.Interface() == fieldValueToCheck {
				return true
			}
		}
	}
	return false
}

// GetEnv : get env from .env file
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ExitF prints error and exit program
func ExitF(format string, a ...interface{}) {
	log.Fatalf(format, a...)
	os.Exit(1)
}

const (
	empty = ""
	tab   = "\t"
)

// PrettyJSON prints object as json
func PrettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

// CalculateSkipOffset ..
func CalculateSkipOffset(page, totalRecords, pageSize int64) (offset int64, take int64, totalPages int64) {
	totalPages = totalRecords / pageSize

	if totalRecords%pageSize > 0 {
		totalPages++
	}

	if page > totalPages {
		offset = 0
		take = 0
	} else {
		offset = (page - 1) * pageSize
		take = pageSize
	}
	return offset, take, totalPages
}

func StructToMapStringInterface(data interface{}) (map[string]interface{}, error) {
	s := structs.New(data)
	return s.Map(), nil
}

func ConvertNumberToAlpha(num int) string {
	res := ""
	if num == 0 {
		res = "A"
	} else if num == 1 {
		res = "B"
	} else if num == 2 {
		res = "C"
	} else if num == 3 {
		res = "D"
	} else if num == 4 {
		res = "E"
	} else if num == 5 {
		res = "F"
	} else if num == 6 {
		res = "G"
	} else {
		res = ""
	}
	return res
}

func RemoveSpecialCharacter(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "?", "")
	str = strings.ReplaceAll(str, "\\", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ":", "")
	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, "à", "a")
	str = strings.ReplaceAll(str, "á", "a")
	str = strings.ReplaceAll(str, "ạ", "a")
	str = strings.ReplaceAll(str, "ả", "a")
	str = strings.ReplaceAll(str, "ã", "a")
	str = strings.ReplaceAll(str, "â", "a")
	str = strings.ReplaceAll(str, "ầ", "a")
	str = strings.ReplaceAll(str, "ấ", "a")
	str = strings.ReplaceAll(str, "ậ", "a")
	str = strings.ReplaceAll(str, "ẩ", "a")
	str = strings.ReplaceAll(str, "ẫ", "a")
	str = strings.ReplaceAll(str, "ă", "a")
	str = strings.ReplaceAll(str, "ằ", "a")
	str = strings.ReplaceAll(str, "ắ", "a")
	str = strings.ReplaceAll(str, "ặ", "a")
	str = strings.ReplaceAll(str, "ẳ", "a")
	str = strings.ReplaceAll(str, "ẵ", "a")
	str = strings.ReplaceAll(str, "ễ", "e")
	str = strings.ReplaceAll(str, "ể", "e")
	str = strings.ReplaceAll(str, "ệ", "e")
	str = strings.ReplaceAll(str, "ế", "e")
	str = strings.ReplaceAll(str, "ề", "e")
	str = strings.ReplaceAll(str, "ê", "e")
	str = strings.ReplaceAll(str, "ẽ", "e")
	str = strings.ReplaceAll(str, "ẻ", "e")
	str = strings.ReplaceAll(str, "ẹ", "e")
	str = strings.ReplaceAll(str, "é", "e")
	str = strings.ReplaceAll(str, "è", "e")
	str = strings.ReplaceAll(str, "ì", "i")
	str = strings.ReplaceAll(str, "í", "i")
	str = strings.ReplaceAll(str, "ị", "i")
	str = strings.ReplaceAll(str, "ỉ", "i")
	str = strings.ReplaceAll(str, "ĩ", "i")
	str = strings.ReplaceAll(str, "ò", "o")
	str = strings.ReplaceAll(str, "ó", "o")
	str = strings.ReplaceAll(str, "ọ", "o")
	str = strings.ReplaceAll(str, "ỏ", "o")
	str = strings.ReplaceAll(str, "õ", "o")
	str = strings.ReplaceAll(str, "ô", "o")
	str = strings.ReplaceAll(str, "ồ", "o")
	str = strings.ReplaceAll(str, "ố", "o")
	str = strings.ReplaceAll(str, "ộ", "o")
	str = strings.ReplaceAll(str, "ổ", "o")
	str = strings.ReplaceAll(str, "ỗ", "o")
	str = strings.ReplaceAll(str, "ơ", "o")
	str = strings.ReplaceAll(str, "ờ", "o")
	str = strings.ReplaceAll(str, "ớ", "o")
	str = strings.ReplaceAll(str, "ợ", "o")
	str = strings.ReplaceAll(str, "ở", "o")
	str = strings.ReplaceAll(str, "ỡ", "o")
	str = strings.ReplaceAll(str, "ù", "u")
	str = strings.ReplaceAll(str, "ú", "u")
	str = strings.ReplaceAll(str, "ụ", "u")
	str = strings.ReplaceAll(str, "ủ", "u")
	str = strings.ReplaceAll(str, "ũ", "u")
	str = strings.ReplaceAll(str, "ư", "u")
	str = strings.ReplaceAll(str, "ừ", "u")
	str = strings.ReplaceAll(str, "ứ", "u")
	str = strings.ReplaceAll(str, "ự", "u")
	str = strings.ReplaceAll(str, "ử", "u")
	str = strings.ReplaceAll(str, "ữ", "u")
	str = strings.ReplaceAll(str, "ỳ", "y")
	str = strings.ReplaceAll(str, "ý", "y")
	str = strings.ReplaceAll(str, "ỵ", "y")
	str = strings.ReplaceAll(str, "ỷ", "y")
	str = strings.ReplaceAll(str, "ỹ", "y")
	str = strings.ReplaceAll(str, "đ", "d")
	str = strings.ReplaceAll(str, "\u0300", "")
	str = strings.ReplaceAll(str, "\u0301", "")
	str = strings.ReplaceAll(str, "\u0303", "")
	str = strings.ReplaceAll(str, "\u0309", "")
	str = strings.ReplaceAll(str, "\u0323", "")
	str = strings.ReplaceAll(str, "\u031B", "")
	str = strings.ReplaceAll(str, "\u0306", "")
	str = strings.ReplaceAll(str, "\u02C6", "")
	return str
}

// Tạo ra một chuỗi ngẫu nhiên từ các ký tự cho trước.
func randomString(n int64, prefix string) string {
	var letters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return prefix + string(s)
}

func GenerateRandomStrings(n int64, x int64, prefix string) []string {
	rand.Seed(time.Now().UnixNano())
	if x > 1000 {
		x = 1000
	}
	result := make([]string, 0, x)
	lookup := make(map[string]bool)

	for int64(len(result)) < x {
		str := randomString(n, prefix)
		if _, exist := lookup[str]; !exist {
			result = append(result, str)
			lookup[str] = true
		}
	}

	return result
}

func ExtractSpreadsheetID(url string) (string, error) {
	// Mẫu biểu thức chính quy cho việc trích xuất ID của Google Spreadsheet từ URL
	re := regexp.MustCompile(`https://docs\.google\.com/spreadsheets/d/([a-zA-Z0-9-_]+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) < 2 {
		return "", fmt.Errorf("Không tìm thấy ID trong URL")
	}

	return matches[1], nil
}

func ConvertToUnixTimestamp(timeStr string) (int64, error) {
	hours := 0
	minutes := 0
	seconds := 0
	var err error

	parts := regexp.MustCompile(`,\s*`).Split(timeStr, -1)
	for _, part := range parts {
		if strings.Contains(part, "hour") {
			hours, err = parseTime(part)
			if err != nil {
				return 0, err
			}
		} else if strings.Contains(part, "minute") {
			minutes, err = parseTime(part)
			if err != nil {
				return 0, err
			}
		} else if strings.Contains(part, "second") {
			seconds, err = parseTime(part)
			if err != nil {
				return 0, err
			}
		}
	}

	totalSeconds := hours*3600 + minutes*60 + seconds
	return int64(totalSeconds), nil
}

func ConvertToUnixTimestampV2(timeStr string) int64 {
	// Phân tách giờ, phút và giây từ chuỗi
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		fmt.Println("Invalid time format:", timeStr)
		return 0
	}

	// Chuyển đổi giờ, phút và giây thành số nguyên
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Invalid hours:", err)
		return 0
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid minutes:", err)
		return 0

	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Invalid seconds:", err)
		return 0
	}

	// Tính toán tổng số giây
	totalSeconds := hours*3600 + minutes*60 + seconds
	return int64(totalSeconds)
}

func parseTime(s string) (int, error) {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		return strconv.Atoi(matches[1])
	}
	return 0, fmt.Errorf("could not parse time from string")
}
