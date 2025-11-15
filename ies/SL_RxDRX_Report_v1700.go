package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RxDRX_Report_v1700 struct {
	sl_DRX_ConfigFromTx_r17 SL_DRX_ConfigUC_SemiStatic_r17 `madatory`
}

func (ie *SL_RxDRX_Report_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_DRX_ConfigFromTx_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DRX_ConfigFromTx_r17", err)
	}
	return nil
}

func (ie *SL_RxDRX_Report_v1700) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_DRX_ConfigFromTx_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DRX_ConfigFromTx_r17", err)
	}
	return nil
}
