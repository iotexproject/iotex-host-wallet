package controller

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-host-wallet/wallet/config"
	"github.com/iotexproject/iotex-host-wallet/wallet/dao"
	"github.com/iotexproject/iotex-host-wallet/wallet/key"
	"github.com/labstack/echo"
)

// SignerRoute signer router
func SignerRoute(e *echo.Echo) {
	e.POST("/sign", sign)
}

// SignRequest sign request
type SignRequest struct {
	UserID string `json:"userID" form:"userID"`
	Data   string `json:"data" form:"data"`
	Sign   string `json:"sign" form:"sign"`
}

// SignResponse sign response
type SignResponse struct {
	Data      string `json:"data"`
	PublicKey string `json:"publicKey"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

func sign(c echo.Context) error {
	request := new(SignRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := config.Verify(request.UserID+request.Data, request.Sign, config.C.Container.ServicePublicKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "sign error")
	}
	addr, err := dao.AddressFindByUserID(request.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if addr == nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("userID %s not found", request.UserID))
	}
	pkb, err := key.MK.ChildKey(addr.Index)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pk, err := crypto.BytesToPrivateKey(pkb[1:])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	decoded, err := hex.DecodeString(request.Data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "decode data error")
	}
	data, err := pk.Sign(decoded[:])
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "sign data error")
	}
	encodedData := hex.EncodeToString(data)
	resp := &SignResponse{
		Data:      encodedData,
		PublicKey: pk.PublicKey().HexString(),
		Timestamp: time.Now().Unix(),
	}
	sign, err := config.Sign(fmt.Sprintf("%s%s%d", resp.Data, resp.PublicKey, resp.Timestamp), config.C.Container.WalletPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp.Sign = sign
	return c.JSON(http.StatusOK, resp)
}
