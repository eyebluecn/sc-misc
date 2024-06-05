package util

import (
	"context"
	"sync"
)

//https://github.com/smirzaei/parallel

func ForEach[T any](arr []T, fn func(T)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	for _, item := range arr {
		go func(x T) {
			defer wg.Done()

			fn(x)
		}(item)
	}

	wg.Wait()
}

// ForEachEnhance 并行执行
func ForEachEnhance[T any](arr []T, fn func(T) error) error {
	wg := &sync.WaitGroup{}

	length := len(arr)
	wg.Add(length)

	errorCollector := make([]error, length)
	for i, item := range arr {
		go func(index int, x T) {
			defer wg.Done()

			err := fn(x)
			errorCollector[index] = err
		}(i, item)
	}

	wg.Wait()

	// 检查是否有任务失败
	for _, err := range errorCollector {
		if err != nil {
			return err
		}
	}
	return nil
}

func ForEachLimit[T any](arr []T, limit int, fn func(T)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	limiter := make(chan struct{}, limit)

	for _, item := range arr {
		limiter <- struct{}{}
		go func(x T) {
			defer wg.Done()

			fn(x)
			<-limiter
		}(item)
	}

	wg.Wait()
}

func Map[T1 any, T2 any](arr []T1, fn func(T1) T2) []T2 {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	output := make([]T2, len(arr))

	for i := range arr {
		go func(index int, x T1) {
			defer wg.Done()

			result := fn(x)
			output[index] = result

		}(i, arr[i])
	}

	wg.Wait()
	return output
}

// MapFast 快捷模式，方便一句话调用。能够搜集子任务的错误，会返回第一条错误。
func MapFast[T1 any, T2 any](ctx context.Context, arr []T1, fn func(context.Context, T1) (T2, error)) ([]T2, error) {

	length := len(arr)

	wg := &sync.WaitGroup{}
	wg.Add(length)

	resultCollector := make([]T2, length)
	errorCollector := make([]error, length)

	for i := range arr {
		go func(index int, x T1) {
			defer wg.Done()

			result, err := fn(ctx, x)
			if err != nil {
				errorCollector[index] = err
			} else {
				resultCollector[index] = result
			}
		}(i, arr[i])
	}

	wg.Wait()

	// 检查是否有任务失败
	for _, err := range errorCollector {
		if err != nil {
			return nil, err
		}
	}

	return resultCollector, nil
}

// MapEnhance 手动模式
func MapEnhance[T1 any, T2 any](arr []T1, fn func(T1) (T2, error)) ([]T2, error) {

	length := len(arr)

	wg := &sync.WaitGroup{}
	wg.Add(length)

	resultCollector := make([]T2, length)
	errorCollector := make([]error, length)

	for i := range arr {
		go func(index int, x T1) {
			defer wg.Done()

			result, err := fn(x)
			if err != nil {
				errorCollector[index] = err
			} else {
				resultCollector[index] = result
			}
		}(i, arr[i])
	}

	wg.Wait()

	// 检查是否有任务失败
	for _, err := range errorCollector {
		if err != nil {
			return nil, err
		}
	}

	return resultCollector, nil
}

func MapLimit[T1 any, T2 any](arr []T1, limit int, fn func(T1) T2) []T2 {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	output := make([]T2, len(arr))
	limiter := make(chan struct{}, limit)

	for i := range arr {
		limiter <- struct{}{}
		go func(index int, x T1) {
			defer wg.Done()

			result := fn(x)
			output[index] = result

			<-limiter
		}(i, arr[i])
	}

	wg.Wait()
	return output
}
