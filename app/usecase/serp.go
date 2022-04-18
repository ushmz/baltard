//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"ratri/domain/model"
	repo "ratri/domain/repository"
	"sort"
)

// SerpUsecase : Abstract operations that SERP usecase should have.
type SerpUsecase interface {
	FetchSerp(taskID, offset int) (*[]model.SearchPage, error)
	FetchSerpWithIcon(taskID, offset, top int) (*[]model.SerpWithIcon, error)
	FetchSerpWithRatio(taskID, offset, top int) (*[]model.SerpWithRatio, error)
}

// SerpImpl : Struct for serp usecase
type SerpImpl struct {
	lpRepo   repo.LinkedPageRepository
	serpRepo repo.SerpRepository
}

// NewSerpUsecase : Return new serp usecase struct
func NewSerpUsecase(serpRepository repo.SerpRepository, linkedPageRepository repo.LinkedPageRepository) SerpUsecase {
	return &SerpImpl{lpRepo: linkedPageRepository, serpRepo: serpRepository}
}

// FetchSerp : Get search results by taskID
func (s *SerpImpl) FetchSerp(taskID, offset int) (*[]model.SearchPage, error) {
	return s.serpRepo.FetchSerpByTaskID(taskID, offset)
}

// FetchSerpWithIcon : Get search results for Icon UI by taskID
func (s *SerpImpl) FetchSerpWithIcon(taskID, offset, top int) (*[]model.SerpWithIcon, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithIcon{}

	srp, err := s.serpRepo.FetchSerpByTaskID(taskID, offset)
	if err != nil {
		return &serp, err
	}

	pageIds := []int{}
	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithIcon{}

	for _, v := range *srp {
		pageIds = append(pageIds, v.PageID)
		serpMap[v.PageID] = model.SerpWithIcon{
			PageID:  v.PageID,
			Title:   v.Title,
			URL:     v.URL,
			Snippet: v.Snippet,
			Linked:  []model.LinkedPage{},
		}
	}

	linked, err := s.lpRepo.GetBySearchPageIDs(pageIds, taskID, top)
	if err != nil {
		return &serp, err
	}

	for _, v := range *linked {
		tempSerp := serpMap[v.PageID]
		tempSerp.Linked = append(tempSerp.Linked, model.LinkedPage{
			ID:       v.ID,
			Title:    v.Title,
			URL:      v.URL,
			Icon:     v.Icon,
			Category: v.Category,
		})
		serpMap[v.PageID] = tempSerp
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

// FetchSerpWithRatio : Get search results for Ratio UI by taskID
func (s *SerpImpl) FetchSerpWithRatio(taskID, offset, top int) (*[]model.SerpWithRatio, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithRatio{}

	srp, err := s.serpRepo.FetchSerpByTaskID(taskID, offset)
	if err != nil {
		return &serp, err
	}

	pageIds := []int{}
	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithRatio{}

	for _, v := range *srp {
		pageIds = append(pageIds, v.PageID)
		serpMap[v.PageID] = model.SerpWithRatio{
			PageID:       v.PageID,
			Title:        v.Title,
			URL:          v.URL,
			Snippet:      v.Snippet,
			Total:        0,
			Distribution: []model.CategoryCount{},
		}
	}

	linked, err := s.lpRepo.GetRatioBySearchPageIDs(pageIds, taskID)
	if err != nil {
		return &serp, err
	}

	for _, v := range *linked {
		tempSerp := serpMap[v.PageID]
		tempSerp.Total += v.CategoryCount
		if len(tempSerp.Distribution) < top {
			tempSerp.Distribution = append(tempSerp.Distribution, model.CategoryCount{
				Category: v.Category,
				Count:    v.CategoryCount,
			})
		}
		serpMap[v.PageID] = tempSerp
	}

	sort.Ints(pageIds)
	for _, v := range pageIds {
		serp = append(serp, serpMap[v])
	}

	return &serp, nil
}
