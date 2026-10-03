package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	buf := make([]byte, 4096)
	var tail []byte

	for {
		// проверка отмены контекста в начале итерации
		if err := ctx.Err(); err != nil {
			return false, err
		}

		k, err := r.Read(buf)

		if k > 0 {
			data := make([]byte, len(tail)+k)
			copy(data, tail)
			copy(data[len(tail):], buf[:k])

			// bytes.Contains(data, seq) вернёт true и для пустого seq
			if bytes.Contains(data, seq) {
				return true, nil
			}

			keep := len(seq) - 1
			if keep > len(data) {
				keep = len(data)
			}
			if keep < 0 {
				keep = 0 // защита от seq == []byte{}
			}
			tail = make([]byte, keep)
			copy(tail, data[len(data)-keep:])
		}

		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}
	}
}
