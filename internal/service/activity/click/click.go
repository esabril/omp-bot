package click

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/activity"
)

type ClickService interface {
	Describe(clickID uint64) (*activity.Click, error)
	List(cursor uint64, limit uint64) ([]*activity.Click, error)
	Create(m activity.Click) (uint64, error)
	Update(clickID uint64, click activity.Click) error
	Remove(clickID uint64) (bool, error)
}

type DummyClickService struct {
	clicks []*activity.Click
}

func NewDummyClickService() *DummyClickService {
	return &DummyClickService{
		clicks: []*activity.Click{{Title: "my1"}, {Title: "my2"}, {Title: "my3"}}, // TODO: remove example data
	}
}

func (s *DummyClickService) Describe(clickID uint64) (*activity.Click, error) {
	if clickID >= uint64(len(s.clicks)) {
		return nil, errors.New(s.getOutOfRangeErrorString("can't get item", clickID))
	}

	return s.clicks[clickID], nil
}

func (s *DummyClickService) List(cursor uint64, limit uint64) ([]*activity.Click, error) {
	if cursor == 0 && limit == 0 {
		return s.clicks, nil
	}

	l := uint64(len(s.clicks))

	if cursor >= l {
		return nil, errors.New("out of click's list range")
	}

	if cursor+limit >= l {
		return s.clicks[cursor:], nil
	}

	return s.clicks[cursor : limit+cursor], nil
}

func (s *DummyClickService) Create(m activity.Click) (uint64, error) {
	s.clicks = append(s.clicks, &m)

	return uint64(len(s.clicks) - 1), nil
}

func (s *DummyClickService) Update(clickID uint64, click activity.Click) error {
	length := uint64(len(s.clicks))

	if clickID >= length {
		return errors.New(s.getOutOfRangeErrorString("can't update item", clickID))
	}

	s.clicks[clickID] = &click

	return nil
}

func (s *DummyClickService) Remove(clickID uint64) (bool, error) {
	length := uint64(len(s.clicks))

	if clickID >= length {
		return false, errors.New(s.getOutOfRangeErrorString("can't remove item", clickID))
	}

	s.clicks = append(s.clicks[0:clickID], s.clicks[clickID+1:]...)

	return true, nil
}

func (s *DummyClickService) getOutOfRangeErrorString(action string, clickID uint64) string {
	length := len(s.clicks)

	if length == 0 {
		return fmt.Sprintf("%s %d: list is empty", action, clickID)
	}

	if length == 1 {
		return fmt.Sprintf("%s %d: now we have only one item with index 0", action, clickID)
	}

	return fmt.Sprintf("%s %d: now we have only %d items — from 0 to %d", action, clickID, length, length-1)
}
