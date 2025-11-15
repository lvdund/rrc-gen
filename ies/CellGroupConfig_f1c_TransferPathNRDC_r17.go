package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_mcg  uper.Enumerated = 0
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_scg  uper.Enumerated = 1
	CellGroupConfig_f1c_TransferPathNRDC_r17_Enum_both uper.Enumerated = 2
)

type CellGroupConfig_f1c_TransferPathNRDC_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CellGroupConfig_f1c_TransferPathNRDC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CellGroupConfig_f1c_TransferPathNRDC_r17", err)
	}
	return nil
}

func (ie *CellGroupConfig_f1c_TransferPathNRDC_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CellGroupConfig_f1c_TransferPathNRDC_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
