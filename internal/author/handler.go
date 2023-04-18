package author

import (
	"NewApiProd/internal/apperror"
	"NewApiProd/internal/author/service"
	"NewApiProd/internal/handlers"
	"NewApiProd/pkg/api/filter"
	"NewApiProd/pkg/api/pagination"
	"NewApiProd/pkg/api/sort"
	"NewApiProd/pkg/logging"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

const (
	authorsURL    = "/authors"
	authorURL     = "/authors/:uuid"
	defaultLimit  = 0
	defaultPtoken = 0
)

type handler struct {
	logger  *logging.Logger
	service *service.Service
}

func NewHandler(service *service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{

		service: service,
		logger:  logger,
	}
}
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, pagination.Middleware(filter.Middleware(sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC), defaultLimit), defaultPtoken))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	o := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	name := r.URL.Query().Get("name")

	if name != "" {
		err2 := o.AddField("name", filter.OperatorLike, name, filter.DataTypeStr)
		if err2 != nil {
			return err2
		}
	}

	age := r.URL.Query().Get("age")

	if age != "" {
		operator := filter.OperatorEq
		value := age
		if strings.Contains(age, ":") {
			splited := strings.Split(age, ":")
			operator = splited[0]
			value = splited[1]
			fmt.Println(splited, operator, value)
		}
		err2 := o.AddField("age", operator, value, filter.DataTypeInt)
		if err2 != nil {
			return err2
		}

	}

	isAlive := r.URL.Query().Get("is_alive")
	if isAlive != "" {
		_, err := strconv.ParseBool(isAlive)
		if err != nil {
			validationErr := apperror.BadRequestError("filter params validation failed", "bool value wrong parameter")
			validationErr.WithParams(map[string]string{
				"is_alive": "this field should be boolean: true or false",
			})

			return validationErr

		}
		err2 := o.AddField("is_alive", filter.OperatorEq, isAlive, filter.DataTypeBool)
		if err2 != nil {
			return err2
		}
	}

	createdAt := r.URL.Query().Get("created_at")
	fmt.Println("1@", createdAt)
	if createdAt != "" {
		var operator string

		if strings.Contains(createdAt, ":") {
			//range
			operator = filter.OperatorBetween
		} else {
			//single
			operator = filter.OperatorEq
		}

		err2 := o.AddField("created_at", operator, createdAt, filter.DataTypeDate)
		if err2 != nil {
			return err2
		}
	}

	var sortOptions sort.Options
	var pOptions pagination.POptions

	if options, ok := r.Context().Value(sort.OptionsContextKey).(sort.Options); ok {
		sortOptions = options
	}
	o2 := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	if poptions, ok := r.Context().Value(pagination.OptionsContextKey).(pagination.POptions); ok {
		pOptions = poptions
	}

	fmt.Println(pOptions, o2, sortOptions)
	all, err := h.service.GetAll(r.Context(), pOptions, o2, sortOptions)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}
