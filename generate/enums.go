package generate

import (
	"strings"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type (
	EnumMember struct {
		Name    string
		GoType  string
		Value   string
		Comment string
	}

	Enum struct {
		Name    string
		GoType  string
		Comment string
		Members []EnumMember
	}
)

func IsEnumType(tp spec.DefineStruct) bool {
	if len(tp.Members) == 0 {
		return false
	}
	isEnum := true
	memberType := tp.Members[0].Type.Name()
	for _, member := range tp.Members {
		// 判断是否都是同一个开头
		if !strings.HasPrefix(member.Name, tp.RawName) {
			isEnum = false
			break
		}
		// 判断成员类型是否一致
		if memberType != member.Type.Name() {
			isEnum = false
			break
		}
	}
	return isEnum
}

func isSetValTag(member spec.Member) (bool, string) {
	for _, tag := range member.Tags() {
		if tag.Key == "val" {
			return true, tag.Name
		}
	}
	return false, ""
}

func ToEnumType(tp spec.DefineStruct) Enum {
	// tp 是外围结构
	items := make([]EnumMember, 0)
	var enumGoType string
	for _, member := range tp.Members {
		// 枚举的go的类型
		enumGoType = member.Type.Name()
		comment := strings.TrimPrefix(member.Comment, "// ")
		comment = strings.TrimPrefix(comment, "//")
		item := EnumMember{
			Name:    member.Name,
			GoType:  tp.RawName,
			Value:   "",
			Comment: comment,
		}

		switch enumGoType {
		case "int32", "int", "int64":
			//  如果是int类型，则从tag中获取val的值
			for _, tag := range member.Tags() {
				if tag.Key == "val" {
					item.Value = tag.Name
					break
				}
			}
		case "string":
			// 如果设置了val标签，则使用val的值，否则使用去掉前缀后的值
			ok, val := isSetValTag(member)
			if !ok {
				// 去掉前缀，然后转下划线
				val = strings.TrimPrefix(member.Name, tp.RawName)
				val = lo.SnakeCase(val)
			}
			item.Value = val
		}

		items = append(items, item)
	}
	return Enum{
		Name:    tp.RawName,
		GoType:  enumGoType,
		Comment: "",
		Members: items,
	}
}
