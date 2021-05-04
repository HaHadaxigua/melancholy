package service

import (
	"container/list"
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"sync"
)

/**
这个文件夹中存放一些不是那么重要的方法。一些逻辑无关的方法
*/

// buildFileSearchItem 构建文件搜索的返回体,需要将文件夹和文件作为统一的格式返回
func buildFileSearchResult(folders []*model.Folder, files []*model.File) *msg.RspFileSearchResult {
	lenFiles, lenFolders := len(files), len(folders)
	var res msg.RspFileSearchResult
	list := make([]*msg.RspFileSearchItem, lenFiles+lenFolders)
	for i := 0; i < lenFiles; i++ {
		list[i] = files[i].ToFileSearchItem()
	}
	cur := lenFiles
	for i := 0; i < lenFolders; i++ {
		list[cur] = folders[i].ToFileSearchItem()
		cur++
	}
	res.List = list
	res.Total = len(list)
	return &res
}

// 栈数据结构
type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	l := &sync.RWMutex{}
	return &Stack{list, l}
}

func (stack Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack Stack) Len() int {
	return stack.list.Len()
}

func (stack Stack) Empty() bool {
	return stack.list.Len() == 0
}