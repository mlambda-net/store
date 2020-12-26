package command

import (
  "bytes"
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "io"
  "net/http"
)

// ImageAdd godoc
// @Summary Add the image
// @Produce json
// @Param data formData true "add image"
// @Success 200 {object} message.ImageUrl
// @Failure 500 {string} string "Internal error"
// @Router /brand [post]
func (c *control) ImageAdd(w http.ResponseWriter, r *http.Request) {
  id := r.FormValue("id")
  file, meta, err := r.FormFile("file")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    buf := bytes.NewBuffer(nil)
    if _, err := io.Copy(buf, file); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    } else {
      rsp, err := c.image.Request(&message.ImageAdd{
        Id:      id,
        File:    meta.Filename,
        Content: buf.Bytes(),
      }).Unwrap()
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
      } else {
        _ = json.NewEncoder(w).Encode(rsp)
      }
    }
    err = file.Close()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

