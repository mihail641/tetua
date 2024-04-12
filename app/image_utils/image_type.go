package image_utils

import "mime/multipart"

const BMP FileExtension = "bmp"
const JPEG FileExtension = "jpeg"
const JPG FileExtension = "jpg"
const PNG FileExtension = "png"

type FileExtension string

const FileType = "image"

var job chan *imageCompression

type imageCompression struct {
	imageFile  multipart.File
	resultChan chan *result
}
type result struct {
	results multipart.File
	err     error
}
