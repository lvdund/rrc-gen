package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_schedulingCellInfo_other struct {
	schedulingCellId     ServCellIndex `madatory`
	cif_InSchedulingCell int64         `lb:1,ub:7,madatory`
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_other) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.schedulingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingCellId", err)
	}
	if err = w.WriteInteger(ie.cif_InSchedulingCell, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger cif_InSchedulingCell", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_other) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.schedulingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode schedulingCellId", err)
	}
	var tmp_int_cif_InSchedulingCell int64
	if tmp_int_cif_InSchedulingCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger cif_InSchedulingCell", err)
	}
	ie.cif_InSchedulingCell = tmp_int_cif_InSchedulingCell
	return nil
}
