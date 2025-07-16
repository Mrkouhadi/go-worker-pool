package main

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

var tasks = []Task{
	{
		ID:       0,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(137 * time.Millisecond)
			return "Task-0 done", nil
		},
	},
	{
		ID:       1,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(482 * time.Millisecond)
			return "Task-1 done", nil
		},
	},
	{
		ID:       2,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(215 * time.Millisecond)
			return "Task-2 done", nil
		},
	},
	{
		ID:       3,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-3 done", nil
		},
	},
	{
		ID:       4,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-4 done", nil
		},
	},
	{
		ID:       5,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-5 done", nil
		},
	},
	{
		ID:       6,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(92 * time.Millisecond)
			return "Task-6 done", nil
		},
	},
	{
		ID:       7,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(387 * time.Millisecond)
			return "Task-7 done", nil
		},
	},
	{
		ID:       8,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(264 * time.Millisecond)
			return "Task-8 done", nil
		},
	},
	{
		ID:       9,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(143 * time.Millisecond)
			return "Task-9 done", nil
		},
	},
	{
		ID:       10,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-10 done", nil
		},
	},
	{
		ID:       11,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(187 * time.Millisecond)
			return "Task-11 done", nil
		},
	},
	{
		ID:       12,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(423 * time.Millisecond)
			return "Task-12 done", nil
		},
	},
	{
		ID:       13,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(156 * time.Millisecond)
			return "Task-13 done", nil
		},
	},
	{
		ID:       14,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(389 * time.Millisecond)
			return "Task-14 done", nil
		},
	},
	{
		ID:       15,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(245 * time.Millisecond)
			return "Task-15 done", nil
		},
	},
	{
		ID:       16,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(432 * time.Millisecond)
			return "Task-16 done", nil
		},
	},
	{
		ID:       17,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-17 done", nil
		},
	},
	{
		ID:       18,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-18 done", nil
		},
	},
	{
		ID:       19,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-19 done", nil
		},
	},
	{
		ID:       20,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-20 done", nil
		},
	},
	{
		ID:       21,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(143 * time.Millisecond)
			return "Task-21 done", nil
		},
	},
	{
		ID:       22,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(376 * time.Millisecond)
			return "Task-22 done", nil
		},
	},
	{
		ID:       23,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(234 * time.Millisecond)
			return "Task-23 done", nil
		},
	},
	{
		ID:       24,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-24 done", nil
		},
	},
	{
		ID:       25,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(189 * time.Millisecond)
			return "Task-25 done", nil
		},
	},
	{
		ID:       26,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-26 done", nil
		},
	},
	{
		ID:       27,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-27 done", nil
		},
	},
	{
		ID:       28,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-28 done", nil
		},
	},
	{
		ID:       29,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-29 done", nil
		},
	},
	{
		ID:       30,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-30 done", nil
		},
	},
	{
		ID:       31,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(143 * time.Millisecond)
			return "Task-31 done", nil
		},
	},
	{
		ID:       32,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(376 * time.Millisecond)
			return "Task-32 done", nil
		},
	},
	{
		ID:       33,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(234 * time.Millisecond)
			return "Task-33 done", nil
		},
	},
	{
		ID:       34,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-34 done", nil
		},
	},
	{
		ID:       35,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(189 * time.Millisecond)
			return "Task-35 done", nil
		},
	},
	{
		ID:       36,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-36 done", nil
		},
	},
	{
		ID:       37,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-37 done", nil
		},
	},
	{
		ID:       38,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-38 done", nil
		},
	},
	{
		ID:       39,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-39 done", nil
		},
	},
	{
		ID:       40,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-40 done", nil
		},
	},
	{
		ID:       41,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(143 * time.Millisecond)
			return "Task-41 done", nil
		},
	},
	{
		ID:       42,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(376 * time.Millisecond)
			return "Task-42 done", nil
		},
	},
	{
		ID:       43,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(234 * time.Millisecond)
			return "Task-43 done", nil
		},
	},
	{
		ID:       44,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-44 done", nil
		},
	},
	{
		ID:       45,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(189 * time.Millisecond)
			return "Task-45 done", nil
		},
	},
	{
		ID:       46,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-46 done", nil
		},
	},
	{
		ID:       47,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-47 done", nil
		},
	},
	{
		ID:       48,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-48 done", nil
		},
	},
	{
		ID:       49,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-49 done", nil
		},
	},

	{
		ID:       50,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(198 * time.Millisecond)
			return "Task-50 done", nil
		},
	},
	{
		ID:       51,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(342 * time.Millisecond)
			return "Task-51 done", nil
		},
	},
	{
		ID:       52,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-52 done", nil
		},
	},
	{
		ID:       53,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(123 * time.Millisecond)
			return "Task-53 done", nil
		},
	},
	{
		ID:       54,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(287 * time.Millisecond)
			return "Task-54 done", nil
		},
	},
	{
		ID:       55,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-55 done", nil
		},
	},
	{
		ID:       56,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(156 * time.Millisecond)
			return "Task-56 done", nil
		},
	},
	{
		ID:       57,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-57 done", nil
		},
	},
	{
		ID:       58,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(234 * time.Millisecond)
			return "Task-58 done", nil
		},
	},
	{
		ID:       59,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-59 done", nil
		},
	},
	{
		ID:       60,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(178 * time.Millisecond)
			return "Task-60 done", nil
		},
	},
	{
		ID:       61,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-61 done", nil
		},
	},
	{
		ID:       62,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-62 done", nil
		},
	},
	{
		ID:       63,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(389 * time.Millisecond)
			return "Task-63 done", nil
		},
	},
	{
		ID:       64,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(145 * time.Millisecond)
			return "Task-64 done", nil
		},
	},
	{
		ID:       65,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(432 * time.Millisecond)
			return "Task-65 done", nil
		},
	},
	{
		ID:       66,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(276 * time.Millisecond)
			return "Task-66 done", nil
		},
	},
	{
		ID:       67,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(198 * time.Millisecond)
			return "Task-67 done", nil
		},
	},
	{
		ID:       68,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(354 * time.Millisecond)
			return "Task-68 done", nil
		},
	},
	{
		ID:       69,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-69 done", nil
		},
	},
	{
		ID:       70,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(167 * time.Millisecond)
			return "Task-70 done", nil
		},
	},
	{
		ID:       71,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(289 * time.Millisecond)
			return "Task-71 done", nil
		},
	},
	{
		ID:       72,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(376 * time.Millisecond)
			return "Task-72 done", nil
		},
	},
	{
		ID:       73,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(143 * time.Millisecond)
			return "Task-73 done", nil
		},
	},
	{
		ID:       74,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-74 done", nil
		},
	},
	{
		ID:       75,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(234 * time.Millisecond)
			return "Task-75 done", nil
		},
	},
	{
		ID:       76,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(321 * time.Millisecond)
			return "Task-76 done", nil
		},
	},
	{
		ID:       77,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(456 * time.Millisecond)
			return "Task-77 done", nil
		},
	},
	{
		ID:       78,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(189 * time.Millisecond)
			return "Task-78 done", nil
		},
	},
	{
		ID:       79,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(267 * time.Millisecond)
			return "Task-79 done", nil
		},
	},
	{
		ID:       80,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(412 * time.Millisecond)
			return "Task-80 done", nil
		},
	},
	{
		ID:       81,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(156 * time.Millisecond)
			return "Task-81 done", nil
		},
	},
	{
		ID:       82,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(378 * time.Millisecond)
			return "Task-82 done", nil
		},
	},
	{
		ID:       83,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(291 * time.Millisecond)
			return "Task-83 done", nil
		},
	},
	{
		ID:       84,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(467 * time.Millisecond)
			return "Task-84 done", nil
		},
	},
	{
		ID:       85,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(103 * time.Millisecond)
			return "Task-85 done", nil
		},
	},
	{
		ID:       86,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(324 * time.Millisecond)
			return "Task-86 done", nil
		},
	},
	{
		ID:       87,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(245 * time.Millisecond)
			return "Task-87 done", nil
		},
	},
	{
		ID:       88,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(189 * time.Millisecond)
			return "Task-88 done", nil
		},
	},
	{
		ID:       89,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(431 * time.Millisecond)
			return "Task-89 done", nil
		},
	},
	{
		ID:       90,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(276 * time.Millisecond)
			return "Task-90 done", nil
		},
	},
	{
		ID:       91,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(498 * time.Millisecond)
			return "Task-91 done", nil
		},
	},
	{
		ID:       92,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(112 * time.Millisecond)
			return "Task-92 done", nil
		},
	},
	{
		ID:       93,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(367 * time.Millisecond)
			return "Task-93 done", nil
		},
	},
	{
		ID:       94,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(204 * time.Millisecond)
			return "Task-94 done", nil
		},
	},
	{
		ID:       95,
		Priority: 3,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(489 * time.Millisecond)
			return "Task-95 done", nil
		},
	},
	{
		ID:       96,
		Priority: 1,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(157 * time.Millisecond)
			return "Task-96 done", nil
		},
	},
	{
		ID:       97,
		Priority: 4,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(332 * time.Millisecond)
			return "Task-97 done", nil
		},
	},
	{
		ID:       98,
		Priority: 2,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(278 * time.Millisecond)
			return "Task-98 done", nil
		},
	},
	{
		ID:       99,
		Priority: 0,
		Timeout:  5 * time.Second,
		Retries:  3,
		ResultCh: nil,
		Exec: func(ctx context.Context) (any, error) {
			if rand.Float32() < 0.3 {
				return nil, errors.New("random error")
			}
			time.Sleep(421 * time.Millisecond)
			return "Task-99 done", nil
		},
	},
}
