package msg

// 用于fileInclude的文件排序方法
type IFileSearchResultSort func(r *RspFileSearchResult, i, j int) bool

var RspFileSearchResultSort IFileSearchResultSort

// 动态构建排序顺序
func BuildRspFileSearchResultSort(sortedBy string, descending bool) {
	switch sortedBy {
	case "name":
		RspFileSearchResultSort = func(r *RspFileSearchResult, i, j int) bool {
			if descending {
				return r.List[i].Filename < r.List[j].Filename
			}
			return r.List[i].Filename > r.List[j].Filename
		}
	case "size":
		RspFileSearchResultSort = func(r *RspFileSearchResult, i, j int) bool {
			if descending {
				return r.List[i].Size < r.List[j].Size
			}
			return r.List[i].Size > r.List[j].Size
		}
	case "createdAt":
		RspFileSearchResultSort = func(r *RspFileSearchResult, i, j int) bool {
			if descending {
				return r.List[i].CreatedAt.Before(r.List[j].CreatedAt)
			}
			return r.List[i].CreatedAt.After(r.List[j].CreatedAt)
		}
	case "updatedAt":
		RspFileSearchResultSort = func(r *RspFileSearchResult, i, j int) bool {
			if descending {
				return r.List[i].UpdatedAt.Before(r.List[j].UpdatedAt)
			}
			return r.List[i].UpdatedAt.After(r.List[j].UpdatedAt)
		}
	default:
		RspFileSearchResultSort = func(r *RspFileSearchResult, i, j int) bool {
			return r.List[i].Filename < r.List[j].Filename
		}
	}
}
