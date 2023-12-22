package services

import (
	routey "github.com/joseph-beck/routey/pkg/router"
)

type Service interface {
	Add() []routey.Route
}
