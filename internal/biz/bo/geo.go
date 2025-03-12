package bo

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// ListContinentsParams 列出大洲的参数
type ListContinentsParams struct {
	PageSize  int32
	PageToken string
	// 解析后的偏移量，内部使用
	Offset int
}

// ListCountriesParams 列出国家的参数
type ListCountriesParams struct {
	ContinentID uint
	PageSize    int32
	PageToken   string
	// 解析后的偏移量，内部使用
	Offset int
}

// ListStatesParams 列出州/省的参数
type ListStatesParams struct {
	CountryID uint
	PageSize  int32
	PageToken string
	// 解析后的偏移量，内部使用
	Offset int
}

// ListCitiesParams 列出城市的参数
type ListCitiesParams struct {
	StateID   uint
	PageSize  int32
	PageToken string
	// 解析后的偏移量，内部使用
	Offset int
}

// PaginationResult 分页结果
type PaginationResult struct {
	// 下一页的令牌，如果为空表示没有更多数据
	NextPageToken string
	// 是否有更多数据
	HasMore bool
}

// 生成下一页的令牌
func GenerateNextPageToken(offset int, hasMore bool) string {
	if !hasMore {
		return ""
	}
	token := fmt.Sprintf("offset=%d", offset)
	return base64.StdEncoding.EncodeToString([]byte(token))
}

// 解析分页令牌
func ParsePageToken(token string) (int, error) {
	if token == "" {
		return 0, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	parts := strings.Split(string(decoded), "=")
	if len(parts) != 2 || parts[0] != "offset" {
		return 0, fmt.Errorf("invalid page token format")
	}

	offset, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return offset, nil
}
