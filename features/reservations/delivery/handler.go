package delivery

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	reservationService reservations.ReservationServiceInterface_
}

// AddReservation implements reservations.ReservationDeliveryInterface_
func (reservationHandler *ReservationHandler) AddReservation(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	idParam, errParam := helpers.ExtractIDParam(c)
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errParam.Error()))
	}

	input := reservations.ReservationInsert{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.RESERVATION_ErrorBindReservationData))
	}

	reservationEntity, errMapping := insertToEntity(input)
	if errMapping != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errMapping.Error()))
	}

	errInsert := reservationHandler.reservationService.Create(userID, uint(idParam), reservationEntity)
	if errInsert != nil {
		return c.JSON(helpers.ErrorResponse(errInsert))
	}

	return c.JSON(http.StatusCreated, helpers.Response(consts.RESERVATION_InsertSuccess))
}

// CheckReservation implements reservations.ReservationDeliveryInterface_
func (*ReservationHandler) CheckReservation(c echo.Context) error {
	panic("unimplemented")
}

// GetAllReservation implements reservations.ReservationDeliveryInterface_
func (*ReservationHandler) GetAllReservation(c echo.Context) error {
	panic("unimplemented")
}

func New(reservationService reservations.ReservationServiceInterface_) reservations.ReservationDeliveryInterface_ {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}
