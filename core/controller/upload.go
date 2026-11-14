package controller

import (
	"mime/multipart"
	"path"
	"regexp"
)

var validImageExt = regexp.MustCompile(`^[.](jpg|png|jpeg)$`)

var validImageMime = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/jpg":  true,
}

// validateImageFile memeriksa ekstensi dan MIME type file gambar.
// Mengembalikan true jika keduanya valid.
func validateImageFile(fh *multipart.FileHeader) bool {
	ext := path.Ext(fh.Filename)
	if !validImageExt.MatchString(ext) {
		return false
	}
	mime := fh.Header.Get("Content-Type")
	return validImageMime[mime]
}
