package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_TimersAndConstants struct {
	t300 UE_TimersAndConstants_t300 `madatory`
	t301 UE_TimersAndConstants_t301 `madatory`
	t310 UE_TimersAndConstants_t310 `madatory`
	n310 UE_TimersAndConstants_n310 `madatory`
	t311 UE_TimersAndConstants_t311 `madatory`
	n311 UE_TimersAndConstants_n311 `madatory`
	t319 UE_TimersAndConstants_t319 `madatory`
}

func (ie *UE_TimersAndConstants) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.t300.Encode(w); err != nil {
		return utils.WrapError("Encode t300", err)
	}
	if err = ie.t301.Encode(w); err != nil {
		return utils.WrapError("Encode t301", err)
	}
	if err = ie.t310.Encode(w); err != nil {
		return utils.WrapError("Encode t310", err)
	}
	if err = ie.n310.Encode(w); err != nil {
		return utils.WrapError("Encode n310", err)
	}
	if err = ie.t311.Encode(w); err != nil {
		return utils.WrapError("Encode t311", err)
	}
	if err = ie.n311.Encode(w); err != nil {
		return utils.WrapError("Encode n311", err)
	}
	if err = ie.t319.Encode(w); err != nil {
		return utils.WrapError("Encode t319", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.t300.Decode(r); err != nil {
		return utils.WrapError("Decode t300", err)
	}
	if err = ie.t301.Decode(r); err != nil {
		return utils.WrapError("Decode t301", err)
	}
	if err = ie.t310.Decode(r); err != nil {
		return utils.WrapError("Decode t310", err)
	}
	if err = ie.n310.Decode(r); err != nil {
		return utils.WrapError("Decode n310", err)
	}
	if err = ie.t311.Decode(r); err != nil {
		return utils.WrapError("Decode t311", err)
	}
	if err = ie.n311.Decode(r); err != nil {
		return utils.WrapError("Decode n311", err)
	}
	if err = ie.t319.Decode(r); err != nil {
		return utils.WrapError("Decode t319", err)
	}
	return nil
}
