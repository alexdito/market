package response

import (
	"application/internal/app/src/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sort"
)

type Pair struct {
	Key           uuid.UUID
	Value         NavigationItemResponse
	ChildPairList PairList
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value.Sort < p[j].Value.Sort
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type SortedNavigation struct {
	ID           uuid.UUID            `json:"id"`
	Name         string               `json:"name"`
	Code         string               `json:"code"`
	Description  string               `json:"description"`
	Sort         int                  `json:"sort"`
	ChildrenList SortedNavigationList `json:"children_list"`
}

type SortedNavigationList []SortedNavigation

type NavigationListStruct struct {
	C           *gin.Context
	Navigations []entity.Navigation
}

type NavigationItemStruct struct {
	C *gin.Context
	entity.Navigation
}

type NavigationItemResponse struct {
	ID           uuid.UUID                            `json:"id"`
	Name         string                               `json:"name"`
	Code         string                               `json:"code"`
	Description  string                               `json:"description"`
	ParentUuid   uuid.UUID                            `json:"parent_uuid"`
	Sort         int                                  `json:"sort"`
	ChildrenList map[uuid.UUID]NavigationItemResponse `json:"children_list"`
}

func (s *NavigationListStruct) Response() SortedNavigationList {
	var response = make(map[uuid.UUID]NavigationItemResponse)

	for _, navigation := range s.Navigations {
		response[navigation.ID] = NavigationItemResponse{
			ID:           navigation.ID,
			Name:         navigation.Name,
			Code:         navigation.Code,
			Description:  navigation.Description,
			ParentUuid:   navigation.ParentUuid,
			Sort:         navigation.Sort,
			ChildrenList: make(map[uuid.UUID]NavigationItemResponse),
		}
	}

	sortedNavigationList := fillSortedNavigationList(
		sortBySort(
			buildTree(&response, uuid.UUID{}),
		),
	)

	return sortedNavigationList
}

func buildTree(employees *map[uuid.UUID]NavigationItemResponse, parentUuid uuid.UUID) map[uuid.UUID]NavigationItemResponse {
	tree := map[uuid.UUID]NavigationItemResponse{}

	for _, v := range *employees {
		if v.ParentUuid == parentUuid {
			children := buildTree(employees, v.ID)

			if len(children) > 0 {
				v.ChildrenList = children
			}

			tree[v.ID] = v
		}
	}

	return tree
}

func fillSortedNavigationList(list PairList) SortedNavigationList {
	sortedList := SortedNavigationList{}

	for _, v := range list {
		childrenList := SortedNavigationList{}

		if v.ChildPairList != nil {
			childrenList = fillSortedNavigationList(v.ChildPairList)
		}

		sortedList = append(sortedList, SortedNavigation{
			ID:           v.Value.ID,
			Name:         v.Value.Name,
			Code:         v.Value.Code,
			Description:  v.Value.Description,
			Sort:         v.Value.Sort,
			ChildrenList: childrenList,
		})
	}

	return sortedList
}

func sortBySort(navigations map[uuid.UUID]NavigationItemResponse) PairList {
	pairList := make(PairList, len(navigations))
	childrenList := make(PairList, 0)

	pairListIdx := 0
	for k, v := range navigations {
		childrenList = nil

		if len(v.ChildrenList) > 0 {
			childrenList = sortBySort(v.ChildrenList)
		}

		pairList[pairListIdx] = Pair{k, v, childrenList}
		pairListIdx++
	}

	sort.Sort(pairList)

	return pairList
}
