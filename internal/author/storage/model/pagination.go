package model

import (
	"NewApiProd/internal/author/storage"
	"NewApiProd/pkg/api/pagination"
)


type pOptions struct {
	ptoken int
}

func NewPaginationOptions(options pagination.POptions) storage.PaginationOptions {
	return &pOptions{
		ptoken: options.Ptoken,
	}
}


func (p *pOptions) GetPtoken() int {
	return p.ptoken 
}
