package Utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
)

func CheckEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func GetMd5Value(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetFileMd5Value(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}

	md5Value := fmt.Sprintf("%x", md5hash.Sum(nil))
	return md5Value, nil
}

/**
 * model和dto之间的相互转换
 * 注：
 * 1、不支持嵌套结构体，例如DeviceInfo.ExtraInfo.Dnum
 * 2、不支持类型转换
 * @param interface{} data 源数据结构体
 * @param interface{} result 转换后的结构体
 * @param []string src 需要固定转换的源字段，与dst一一对应
 * @param []string dst 需要固定转换的目标字段，与src一一对应
 * @return *Error
 */
func ConvertBetweenModelAndDto(data interface{}, result interface{}, src, dst []string) error {
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	resultType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(result)
	if resultType.Kind() == reflect.Ptr && dataType.Kind() == reflect.Ptr {
		// 传入的data、result是指针，需要.Elem()取得指针指向的value
		dataType = dataType.Elem()
		dataValue = dataValue.Elem()
		resultType = resultType.Elem()
		resultValue = resultValue.Elem()
	} else {
		return errors.New("invalid params")
	}

	//todo 复杂度优化（当前n*n*n）
	for i := 0; i < dataType.NumField(); i++ {
		for j := 0; j < resultType.NumField(); j++ {
			for k := 0;k < len(src); k++ {
				//todo 类型转换
				if dataType.Field(i).Tag.Get("bson")==resultType.Field(j).Tag.Get("json") ||
					dataType.Field(i).Tag.Get("json")==resultType.Field(j).Tag.Get("bson") {
					resultValue.Field(j).Set(dataValue.Field(i))
				}else if (dataType.Field(i).Tag.Get("bson")==src[k] && resultType.Field(j).Tag.Get("json")==dst[k]) ||
					(dataType.Field(i).Tag.Get("json")==src[k] && resultType.Field(j).Tag.Get("bson")==dst[k]) {
					resultValue.Field(j).Set(dataValue.Field(i))
				}
			}
		}
	}

	return nil
}