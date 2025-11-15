package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCellActivationRS_Config_r17 struct {
	scellActivationRS_Id_r17 SCellActivationRS_ConfigId_r17 `madatory`
	resourceSet_r17          NZP_CSI_RS_ResourceSetId       `madatory`
	gapBetweenBursts_r17     *int64                         `lb:2,ub:31,optional`
	qcl_Info_r17             TCI_StateId                    `madatory`
}

func (ie *SCellActivationRS_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gapBetweenBursts_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.scellActivationRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode scellActivationRS_Id_r17", err)
	}
	if err = ie.resourceSet_r17.Encode(w); err != nil {
		return utils.WrapError("Encode resourceSet_r17", err)
	}
	if ie.gapBetweenBursts_r17 != nil {
		if err = w.WriteInteger(*ie.gapBetweenBursts_r17, &uper.Constraint{Lb: 2, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode gapBetweenBursts_r17", err)
		}
	}
	if err = ie.qcl_Info_r17.Encode(w); err != nil {
		return utils.WrapError("Encode qcl_Info_r17", err)
	}
	return nil
}

func (ie *SCellActivationRS_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var gapBetweenBursts_r17Present bool
	if gapBetweenBursts_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.scellActivationRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode scellActivationRS_Id_r17", err)
	}
	if err = ie.resourceSet_r17.Decode(r); err != nil {
		return utils.WrapError("Decode resourceSet_r17", err)
	}
	if gapBetweenBursts_r17Present {
		var tmp_int_gapBetweenBursts_r17 int64
		if tmp_int_gapBetweenBursts_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode gapBetweenBursts_r17", err)
		}
		ie.gapBetweenBursts_r17 = &tmp_int_gapBetweenBursts_r17
	}
	if err = ie.qcl_Info_r17.Decode(r); err != nil {
		return utils.WrapError("Decode qcl_Info_r17", err)
	}
	return nil
}
