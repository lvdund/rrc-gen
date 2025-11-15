package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17 struct {
	cellBarredRedCap1Rx_r17 RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17_cellBarredRedCap1Rx_r17 `madatory`
	cellBarredRedCap2Rx_r17 RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17_cellBarredRedCap2Rx_r17 `madatory`
}

func (ie *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellBarredRedCap1Rx_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cellBarredRedCap1Rx_r17", err)
	}
	if err = ie.cellBarredRedCap2Rx_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cellBarredRedCap2Rx_r17", err)
	}
	return nil
}

func (ie *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellBarredRedCap1Rx_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cellBarredRedCap1Rx_r17", err)
	}
	if err = ie.cellBarredRedCap2Rx_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cellBarredRedCap2Rx_r17", err)
	}
	return nil
}
