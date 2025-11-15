package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MobilityStateParameters struct {
	t_Evaluation       MobilityStateParameters_t_Evaluation `madatory`
	t_HystNormal       MobilityStateParameters_t_HystNormal `madatory`
	n_CellChangeMedium int64                                `lb:1,ub:16,madatory`
	n_CellChangeHigh   int64                                `lb:1,ub:16,madatory`
}

func (ie *MobilityStateParameters) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.t_Evaluation.Encode(w); err != nil {
		return utils.WrapError("Encode t_Evaluation", err)
	}
	if err = ie.t_HystNormal.Encode(w); err != nil {
		return utils.WrapError("Encode t_HystNormal", err)
	}
	if err = w.WriteInteger(ie.n_CellChangeMedium, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger n_CellChangeMedium", err)
	}
	if err = w.WriteInteger(ie.n_CellChangeHigh, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger n_CellChangeHigh", err)
	}
	return nil
}

func (ie *MobilityStateParameters) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.t_Evaluation.Decode(r); err != nil {
		return utils.WrapError("Decode t_Evaluation", err)
	}
	if err = ie.t_HystNormal.Decode(r); err != nil {
		return utils.WrapError("Decode t_HystNormal", err)
	}
	var tmp_int_n_CellChangeMedium int64
	if tmp_int_n_CellChangeMedium, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger n_CellChangeMedium", err)
	}
	ie.n_CellChangeMedium = tmp_int_n_CellChangeMedium
	var tmp_int_n_CellChangeHigh int64
	if tmp_int_n_CellChangeHigh, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger n_CellChangeHigh", err)
	}
	ie.n_CellChangeHigh = tmp_int_n_CellChangeHigh
	return nil
}
