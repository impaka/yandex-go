package main

import (
	"io"
)

// Копирует n байт из r в w.
// Если количество байт, доступных для чтения, меньше n,
// функция должна копировать все данные.
// Если возникает ошибка — возвращать её.
func Copy(r io.Reader, w io.Writer, n uint) error {
	var many int
	buf := make([]byte, int(n))
	minN := int(n)

	for minN > 0 {
		read, err := r.Read(buf)
		if read > 0 {
			_, err := w.Write(buf[:read])
			if err != nil {
				return err
			}
			many = min(read, minN)
			minN -= many
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

	}

	return nil
}

// func main() {
// 	r := strings.NewReader("hello world")
// 	var w bytes.Buffer

// 	err := Copy(r, &w, 5)
// 	fmt.Printf("err = %v, w = %q\n", err, w.String())
// 	// err = <nil>, w = "hello"
// }
