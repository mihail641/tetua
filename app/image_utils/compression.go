package image_utils

import (
	"mime/multipart"
)

// RunImageCompression отправляем файл для сжатия в канал и дожидаемя сжатого файла
func RunImageCompression(file multipart.File) (multipart.File, error) {
	var image imageCompression
	// канал для передачи результатов сжатия картинки
	image.resultChan = make(chan *result)
	go func() {
		job <- &imageCompression{imageFile: file, resultChan: image.resultChan}
	}()

	for {
		select {
		case result := <-image.resultChan:
			if result.results != nil {
				return result.results, nil
			}
			if result.err != nil {
				return nil, result.err
			}

		}
	}
}
