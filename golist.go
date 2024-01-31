package golist

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"sync"
	"time"
)

type List struct {
	mu    sync.Mutex
	items []string
}

func NewList(initialItems []string) *List {
	l := new(List)
	l.items = append(l.items, initialItems...)
	return l
}

func NewFileList(filename string) (*List, error) {
	l := new(List)

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		i := scanner.Text()
		l.items = append(l.items, i)
	}

	return l, nil
}

func (l *List) Count() int {
	return len(l.items)
}

func (l *List) Randomize() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(l.items), func(i, j int) { l.items[i], l.items[j] = l.items[j], l.items[i] })

	return nil
}

func (l *List) Take() (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.items) > 0 {
		i := l.items[0]
		l.items = append(l.items[:0], l.items[1:]...)

		return i, nil
	}

	return "", errors.New("list is empty")
}

func (l *List) Add(item string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.items = append(l.items, item)
	return nil
}
