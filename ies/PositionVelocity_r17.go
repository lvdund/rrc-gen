package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PositionVelocity_r17 struct {
	positionX_r17  PositionStateVector_r17 `madatory`
	positionY_r17  PositionStateVector_r17 `madatory`
	positionZ_r17  PositionStateVector_r17 `madatory`
	velocityVX_r17 VelocityStateVector_r17 `madatory`
	velocityVY_r17 VelocityStateVector_r17 `madatory`
	velocityVZ_r17 VelocityStateVector_r17 `madatory`
}

func (ie *PositionVelocity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.positionX_r17.Encode(w); err != nil {
		return utils.WrapError("Encode positionX_r17", err)
	}
	if err = ie.positionY_r17.Encode(w); err != nil {
		return utils.WrapError("Encode positionY_r17", err)
	}
	if err = ie.positionZ_r17.Encode(w); err != nil {
		return utils.WrapError("Encode positionZ_r17", err)
	}
	if err = ie.velocityVX_r17.Encode(w); err != nil {
		return utils.WrapError("Encode velocityVX_r17", err)
	}
	if err = ie.velocityVY_r17.Encode(w); err != nil {
		return utils.WrapError("Encode velocityVY_r17", err)
	}
	if err = ie.velocityVZ_r17.Encode(w); err != nil {
		return utils.WrapError("Encode velocityVZ_r17", err)
	}
	return nil
}

func (ie *PositionVelocity_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.positionX_r17.Decode(r); err != nil {
		return utils.WrapError("Decode positionX_r17", err)
	}
	if err = ie.positionY_r17.Decode(r); err != nil {
		return utils.WrapError("Decode positionY_r17", err)
	}
	if err = ie.positionZ_r17.Decode(r); err != nil {
		return utils.WrapError("Decode positionZ_r17", err)
	}
	if err = ie.velocityVX_r17.Decode(r); err != nil {
		return utils.WrapError("Decode velocityVX_r17", err)
	}
	if err = ie.velocityVY_r17.Decode(r); err != nil {
		return utils.WrapError("Decode velocityVY_r17", err)
	}
	if err = ie.velocityVZ_r17.Decode(r); err != nil {
		return utils.WrapError("Decode velocityVZ_r17", err)
	}
	return nil
}
