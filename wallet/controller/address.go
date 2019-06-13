package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-host-wallet/wallet/config"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-host-wallet/wallet/key"

	"github.com/iotexproject/iotex-host-wallet/wallet/dao"
	"github.com/labstack/echo"
)

// AddressRoute address router
func AddressRoute(e *echo.Echo) {
	e.POST("/address", getAddress)
}

// GetAddressRequest get address request vo
type GetAddressRequest struct {
	UserID string `json:"userID" form:"userID"`
	Sign   string `json:"sign" form:"sign"`
}

// GetAddressResponse get address response vo
type GetAddressResponse struct {
	Address   string `json:"userID"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

func getAddress(c echo.Context) error {
	request := new(GetAddressRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if request.UserID == "" {
		return errors.New("userID is empty")
	}
	err := config.Verify(request.UserID, request.Sign, config.C.Container.ServicePublicKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "sign error")
	}

	addr, err := dao.AddressFindByUserID(request.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if addr == nil {
		index, err := dao.ConfigNewIndex()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		pkb, err := key.MK.ChildKey(index)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		pk, err := crypto.BytesToPrivateKey(pkb[1:])
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		addrs, err := address.FromBytes(pk.PublicKey().Hash())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		addr = &dao.Address{
			UserID:  request.UserID,
			Index:   index,
			Address: addrs.String(),
		}
		err = addr.Save()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	resp := &GetAddressResponse{
		Address:   addr.Address,
		Timestamp: time.Now().Unix(),
	}
	sign, err := config.Sign(fmt.Sprintf("%s%d", resp.Address, resp.Timestamp), config.C.Container.WalletPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp.Sign = sign
	return c.JSON(http.StatusOK, resp)
}
