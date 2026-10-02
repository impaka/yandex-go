package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// должна найти в данных первое вхождение байт seq,
// которые доступны через Reader r.
// Если последовательность найдена, программа возвращает true, nil,
// иначе false, nil.
// Если возникает ошибка, функция должна возвращать false и ошибку.
func Contains(r io.Reader, seq []byte) (bool, error) {
	// Пустая последовательность содержится в чём угодно.
	if len(seq) == 0 {
		return true, nil
	}

	buf := make([]byte, 4096) // рабочее "окошко"
	tail := make([]byte, 0)   // хвост с прошлого чтения

	for {
		k, err := r.Read(buf)

		if k > 0 {
			// ЯВНО склеиваем хвост + прочитанное в НОВЫЙ массив.
			// Никаких append — чтобы не зависеть от capacity.
			data := make([]byte, len(tail)+k)
			copy(data, tail)
			copy(data[len(tail):], buf[:k])

			if bytes.Contains(data, seq) {
				return true, nil
			}

			// Обновляем хвост: последние len(seq)-1 байт из data.
			keep := len(seq) - 1
			if keep > len(data) {
				keep = len(data)
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

func main() {
	r := strings.NewReader("Hello, World!")
	found, err := Contains(r, []byte("World"))
	fmt.Println(found, err) // true <nil>

	r2 := strings.NewReader("Hello, World!")
	found2, err2 := Contains(r2, []byte("xyz"))
	fmt.Println(found2, err2) // false <nil>
}
