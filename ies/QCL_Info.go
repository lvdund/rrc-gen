package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QCL_Info struct {
	cell            *ServCellIndex           `optional`
	bwp_Id          *BWP_Id                  `optional`
	referenceSignal QCL_Info_referenceSignal `madatory`
	qcl_Type        QCL_Info_qcl_Type        `madatory`
}

func (ie *QCL_Info) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cell != nil, ie.bwp_Id != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cell != nil {
		if err = ie.cell.Encode(w); err != nil {
			return utils.WrapError("Encode cell", err)
		}
	}
	if ie.bwp_Id != nil {
		if err = ie.bwp_Id.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_Id", err)
		}
	}
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	if err = ie.qcl_Type.Encode(w); err != nil {
		return utils.WrapError("Encode qcl_Type", err)
	}
	return nil
}

func (ie *QCL_Info) Decode(r *uper.UperReader) error {
	var err error
	var cellPresent bool
	if cellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_IdPresent bool
	if bwp_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cellPresent {
		ie.cell = new(ServCellIndex)
		if err = ie.cell.Decode(r); err != nil {
			return utils.WrapError("Decode cell", err)
		}
	}
	if bwp_IdPresent {
		ie.bwp_Id = new(BWP_Id)
		if err = ie.bwp_Id.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_Id", err)
		}
	}
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	if err = ie.qcl_Type.Decode(r); err != nil {
		return utils.WrapError("Decode qcl_Type", err)
	}
	return nil
}
