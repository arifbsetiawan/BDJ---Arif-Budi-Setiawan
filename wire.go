// +build wireinject

package main

import (
	"jsc-join-api/app/hospital"
	"jsc-join-api/app/join"
	"jsc-join-api/app/subdistrict"

	"github.com/google/wire"
)

func JoinModule() join.JoinController {
	wire.Build(
		hospital.ProvideHospitalService,
		subdistrict.ProvideSubdistrictService,
		join.ProvideJoinService,
		join.ProvideJoinController,
	)

	return join.JoinController{}
}
