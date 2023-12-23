package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIDIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add the product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchse")
)

func AddProductToCart() gin.HandlerFunc {

}

func RemoveCartItem() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func InstantBuyer() gin.HandlerFunc {

}