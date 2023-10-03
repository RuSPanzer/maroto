package core

import (
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64, config *config.Config, prop *props.Cell)

	// Features
	AddText(text string, cell Cell, prop props.Text)
	CalculateTextHeight(text string, width float64, prop props.Text) float64
	AddSignature(text string, cell Cell, prop props.Text)
	AddMatrixCode(code string, cell Cell, prop props.Rect)
	AddQrCode(code string, cell Cell, rect props.Rect)
	AddBarCode(code string, cell Cell, prop props.Barcode)
	AddImage(value string, cell Cell, prop props.Rect, extension extension.Type)
	CalculateImageHeight(value string, prop props.Rect, extension extension.Type) float64

	// General
	GetDimensions() (width float64, height float64)
	GetMargins() (left float64, top float64, right float64, bottom float64)
	GenerateFile(file string) error
	GenerateBytes() ([]byte, error)
	SetCache(cache cache.Cache)
}
