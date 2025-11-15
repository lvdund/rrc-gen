package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	QCL_Info_qcl_Type_Enum_typeA uper.Enumerated = 0
	QCL_Info_qcl_Type_Enum_typeB uper.Enumerated = 1
	QCL_Info_qcl_Type_Enum_typeC uper.Enumerated = 2
	QCL_Info_qcl_Type_Enum_typeD uper.Enumerated = 3
)

type QCL_Info_qcl_Type struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *QCL_Info_qcl_Type) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode QCL_Info_qcl_Type", err)
	}
	return nil
}

func (ie *QCL_Info_qcl_Type) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode QCL_Info_qcl_Type", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
