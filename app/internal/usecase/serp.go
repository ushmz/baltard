//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"math/rand"
	"time"

	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
)

type Serp interface {
	FetchSerp(taskId, offset int) (*[]model.SearchPage, error)
	FetchSerpWithIcon(taskId, offset, top int) (*[]model.SerpWithIcon, error)
	FetchSerpWithRatio(taskId, offset, top int) (*[]model.SerpWithRatio, error)
}

type SerpImpl struct {
	repository repo.SerpRepository
}

func NewSerpUsecase(serpRepository repo.SerpRepository) Serp {
	return &SerpImpl{repository: serpRepository}
}

func (s *SerpImpl) FetchSerp(taskId, offset int) (*[]model.SearchPage, error) {
	return s.repository.FetchSerpByTaskID(taskId, offset)
}

func (s *SerpImpl) FetchSerpWithIcon(taskId, offset, top int) (*[]model.SerpWithIcon, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithIcon{}

	swi, err := s.repository.FetchSerpWithIconByTaskID(taskId, offset, top)
	if err != nil {
		return &serp, err
	}

	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithIcon{}

	for _, v := range *swi {
		if _, ok := serpMap[v.PageId]; !ok {
			serpMap[v.PageId] = model.SerpWithIcon{
				PageId:  v.PageId,
				Title:   v.Title,
				Url:     v.Url,
				Snippet: v.Snippet,
				Leaks:   []model.SimilarwebPage{},
			}
		}

		if v.SimilarwebId != 0 {
			tempSerp := serpMap[v.PageId]
			tempSerp.Leaks = append(tempSerp.Leaks, model.SimilarwebPage{
				Id:       v.SimilarwebId,
				Title:    v.SimilarwebTitle,
				Url:      v.SimilarwebUrl,
				Icon:     v.SimilarwebIcon,
				Category: v.SimilarwebCategory,
			})
			serpMap[v.PageId] = tempSerp
		}
	}

	for _, v := range serpMap {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(v.Leaks), func(i, j int) { v.Leaks[i], v.Leaks[j] = v.Leaks[j], v.Leaks[i] })

		serp = append(serp, v)
	}

	return &serp, nil
}

func (s *SerpImpl) FetchSerpWithRatio(taskId, offset, top int) (*[]model.SerpWithRatio, error) {
	// serp : Return struct of this method
	serp := []model.SerpWithRatio{}

	swr, err := s.repository.FetchSerpWithRatioByTaskID(taskId, offset, top)
	if err != nil {
		return &serp, err
	}

	// serpMap : Map object to format SQL result to return struct.
	serpMap := map[int]model.SerpWithRatio{}

	for _, v := range *swr {
		if _, ok := serpMap[v.PageId]; !ok {
			serpMap[v.PageId] = model.SerpWithRatio{
				PageId:       v.PageId,
				Title:        v.Title,
				Url:          v.Url,
				Snippet:      v.Snippet,
				Total:        v.SimilarwebCount,
				Distribution: []model.CategoryCount{},
			}
		}

		if v.Category != "" {
			if v.CategoryRank <= top {
				tempSerp := serpMap[v.PageId]
				tempSerp.Distribution = append(tempSerp.Distribution, model.CategoryCount{
					Category:   v.Category,
					Count:      v.CategoryCount,
					Percentage: v.CategoryDistribution,
				})
				serpMap[v.PageId] = tempSerp
			}
		}
	}

	for _, v := range serpMap {
		serp = append(serp, v)
	}

	return &serp, nil
}
