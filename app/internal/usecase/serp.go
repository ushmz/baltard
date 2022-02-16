//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
	"sort"
)

type SerpUsecase interface {
	FetchSerp(taskId, offset int) (*[]model.SearchPage, error)
	FetchSerpWithIcon(taskId, offset, top int) (*[]model.SerpWithIcon, error)
	FetchSerpWithRatio(taskId, offset, top int) (*[]model.SerpWithRatio, error)
}

type SerpImpl struct {
	lpRepo   repo.LinkedPageRepository
	serpRepo repo.SerpRepository
}

func NewSerpUsecase(serpRepository repo.SerpRepository, linkedPageRepository repo.LinkedPageRepository) SerpUsecase {
	return &SerpImpl{lpRepo: linkedPageRepository, serpRepo: serpRepository}
}

func (s *SerpImpl) FetchSerp(taskId, offset int) (*[]model.SearchPage, error) {
	return s.serpRepo.FetchSerpByTaskID(taskId, offset)
}

func (s *SerpImpl) FetchSerpWithIcon(taskId, offset, top int) (*[]model.SerpWithIcon, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithIcon{}

	srp, err := s.serpRepo.FetchSerpByTaskID(taskId, offset)
	if err != nil {
		return &serp, err
	}

	pageIds := []int{}
	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithIcon{}

	for _, v := range *srp {
		pageIds = append(pageIds, v.PageId)
		serpMap[v.PageId] = model.SerpWithIcon{
			PageId:  v.PageId,
			Title:   v.Title,
			Url:     v.Url,
			Snippet: v.Snippet,
			Linked:  []model.LinkedPage{},
		}
	}

	linked, err := s.lpRepo.GetBySearchPageIds(pageIds, taskId, top)
	if err != nil {
		return &serp, err
	}

	for _, v := range *linked {
		tempSerp := serpMap[v.PageId]
		tempSerp.Linked = append(tempSerp.Linked, model.LinkedPage{
			Id:       v.Id,
			Title:    v.Title,
			Url:      v.Url,
			Icon:     v.Icon,
			Category: v.Category,
		})
		serpMap[v.PageId] = tempSerp
	}

	// To fix the order of search result page, sort pageIds
	sort.Ints(pageIds)
	for _, v := range pageIds {
		// If you need to randomize `LinkedPage` order, use following code block.
		// -----------------------------------------
		// rand.Seed(time.Now().UnixNano())
		// rand.Shuffle(len(v.Linked), func(i, j int) { v.Linked[i], v.Linked[j] = v.Linked[j], v.Linked[i] })
		serp = append(serp, serpMap[v])
	}

	return &serp, nil
}

func (s *SerpImpl) FetchSerpWithRatio(taskId, offset, top int) (*[]model.SerpWithRatio, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithRatio{}

	srp, err := s.serpRepo.FetchSerpByTaskID(taskId, offset)
	if err != nil {
		return &serp, err
	}

	pageIds := []int{}
	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithRatio{}

	for _, v := range *srp {
		pageIds = append(pageIds, v.PageId)
		serpMap[v.PageId] = model.SerpWithRatio{
			PageId:       v.PageId,
			Title:        v.Title,
			Url:          v.Url,
			Snippet:      v.Snippet,
			Total:        0,
			Distribution: []model.CategoryCount{},
		}
	}

	linked, err := s.lpRepo.GetRatioBySearchPageIds(pageIds, taskId)
	if err != nil {
		return &serp, err
	}

	for _, v := range *linked {
		tempSerp := serpMap[v.PageId]
		tempSerp.Total += v.CategoryCount
		if len(tempSerp.Distribution) < top {
			tempSerp.Distribution = append(tempSerp.Distribution, model.CategoryCount{
				Category: v.Category,
				Count:    v.CategoryCount,
			})
		}
		serpMap[v.PageId] = tempSerp
	}

	sort.Ints(pageIds)
	for _, v := range pageIds {
		serp = append(serp, serpMap[v])
	}

	return &serp, nil
}
