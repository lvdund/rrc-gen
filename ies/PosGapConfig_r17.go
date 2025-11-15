package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosGapConfig_r17 struct {
	measPosPreConfigGapId_r17 MeasPosPreConfigGapId_r17    `madatory`
	gapOffset_r17             int64                        `lb:0,ub:159,madatory`
	mgl_r17                   PosGapConfig_r17_mgl_r17     `madatory`
	mgrp_r17                  PosGapConfig_r17_mgrp_r17    `madatory`
	mgta_r17                  PosGapConfig_r17_mgta_r17    `madatory`
	gapType_r17               PosGapConfig_r17_gapType_r17 `madatory`
}

func (ie *PosGapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.measPosPreConfigGapId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode measPosPreConfigGapId_r17", err)
	}
	if err = w.WriteInteger(ie.gapOffset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger gapOffset_r17", err)
	}
	if err = ie.mgl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgl_r17", err)
	}
	if err = ie.mgrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgrp_r17", err)
	}
	if err = ie.mgta_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgta_r17", err)
	}
	if err = ie.gapType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode gapType_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.measPosPreConfigGapId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode measPosPreConfigGapId_r17", err)
	}
	var tmp_int_gapOffset_r17 int64
	if tmp_int_gapOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger gapOffset_r17", err)
	}
	ie.gapOffset_r17 = tmp_int_gapOffset_r17
	if err = ie.mgl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgl_r17", err)
	}
	if err = ie.mgrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgrp_r17", err)
	}
	if err = ie.mgta_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgta_r17", err)
	}
	if err = ie.gapType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode gapType_r17", err)
	}
	return nil
}
