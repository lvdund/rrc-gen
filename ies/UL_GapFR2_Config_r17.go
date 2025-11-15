package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_GapFR2_Config_r17 struct {
	gapOffset_r17              int64                         `lb:0,ub:159,madatory`
	ugl_r17                    UL_GapFR2_Config_r17_ugl_r17  `madatory`
	ugrp_r17                   UL_GapFR2_Config_r17_ugrp_r17 `madatory`
	refFR2_ServCellAsyncCA_r17 *ServCellIndex                `optional`
}

func (ie *UL_GapFR2_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.refFR2_ServCellAsyncCA_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.gapOffset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger gapOffset_r17", err)
	}
	if err = ie.ugl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ugl_r17", err)
	}
	if err = ie.ugrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ugrp_r17", err)
	}
	if ie.refFR2_ServCellAsyncCA_r17 != nil {
		if err = ie.refFR2_ServCellAsyncCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode refFR2_ServCellAsyncCA_r17", err)
		}
	}
	return nil
}

func (ie *UL_GapFR2_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var refFR2_ServCellAsyncCA_r17Present bool
	if refFR2_ServCellAsyncCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_gapOffset_r17 int64
	if tmp_int_gapOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger gapOffset_r17", err)
	}
	ie.gapOffset_r17 = tmp_int_gapOffset_r17
	if err = ie.ugl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ugl_r17", err)
	}
	if err = ie.ugrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ugrp_r17", err)
	}
	if refFR2_ServCellAsyncCA_r17Present {
		ie.refFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
		if err = ie.refFR2_ServCellAsyncCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode refFR2_ServCellAsyncCA_r17", err)
		}
	}
	return nil
}
