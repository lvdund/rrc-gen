package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DLInformationTransfer_v1700_IEs_ta_PDC_r17_Enum_activate   uper.Enumerated = 0
	DLInformationTransfer_v1700_IEs_ta_PDC_r17_Enum_deactivate uper.Enumerated = 1
)

type DLInformationTransfer_v1700_IEs_ta_PDC_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DLInformationTransfer_v1700_IEs_ta_PDC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DLInformationTransfer_v1700_IEs_ta_PDC_r17", err)
	}
	return nil
}

func (ie *DLInformationTransfer_v1700_IEs_ta_PDC_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DLInformationTransfer_v1700_IEs_ta_PDC_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
