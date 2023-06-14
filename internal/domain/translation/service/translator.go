package service

import "github.com/sesigl/go-project-ddd-template/internal/domain/translation/entity"

type Translator interface {
	Translate(translation entity.Translation) (entity.Translation, error)
}
