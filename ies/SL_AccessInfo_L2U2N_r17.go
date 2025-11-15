package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_AccessInfo_L2U2N_r17 struct {
	cellAccessRelatedInfo_r17 CellAccessRelatedInfo  `madatory`
	sl_ServingCellInfo_r17    SL_ServingCellInfo_r17 `madatory`
}

func (ie *SL_AccessInfo_L2U2N_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellAccessRelatedInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cellAccessRelatedInfo_r17", err)
	}
	if err = ie.sl_ServingCellInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ServingCellInfo_r17", err)
	}
	return nil
}

func (ie *SL_AccessInfo_L2U2N_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellAccessRelatedInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cellAccessRelatedInfo_r17", err)
	}
	if err = ie.sl_ServingCellInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ServingCellInfo_r17", err)
	}
	return nil
}
