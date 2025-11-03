package utils

import (
	"go-posts/dto"
	"go-posts/errs"
	"go-posts/models"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func ToPublicPost(u models.Post) dto.PublicPost {
	return dto.PublicPost{
		Title:    u.Title,
		Content:  u.Content,
		Category: u.Category,
		Status:   u.Status,
	}
}

func ToPublicPosts(posts []models.Post) []dto.PublicPost {
	result := make([]dto.PublicPost, 0, len(posts))
	for _, post := range posts {
		result = append(result, dto.PublicPost{
			Title:    post.Title,
			Content:  post.Content,
			Category: post.Category,
			Status:   post.Status,
		})
	}
	return result
}

func AssignedKeyModel(model any, data map[string]any) error {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errs.New("model must be pointer to struct", http.StatusBadRequest)
	}
	v = v.Elem()

	for key, value := range data {
		if value == nil {
			continue
		}

		val := reflect.ValueOf(value)

		// Dereference pointer
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				continue
			}
			val = val.Elem()
			value = val.Interface()
		}

		if val.Kind() == reflect.String && val.Len() == 0 {
			continue
		}

		if t, ok := value.(time.Time); ok && t.IsZero() {
			continue
		}

		if strings.EqualFold(key, "Price") {
			if num, ok := value.(float64); ok && num <= 0 {
				continue
			}
			if num, ok := value.(int); ok && num <= 0 {
				continue
			}
			if num, ok := value.(uint); ok && num == 0 {
				continue
			}
		}

		if (val.Kind() == reflect.Slice || val.Kind() == reflect.Map) && val.Len() == 0 {
			continue
		}

		field := v.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		if val.Type().AssignableTo(field.Type()) {
			field.Set(val)
		} else if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		}
	}

	return nil
}
