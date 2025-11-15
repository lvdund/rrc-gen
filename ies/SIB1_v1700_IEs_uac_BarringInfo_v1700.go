package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1700_IEs_uac_BarringInfo_v1700 struct {
	uac_BarringInfoSetList_v1700 UAC_BarringInfoSetList_v1700 `madatory`
}

func (ie *SIB1_v1700_IEs_uac_BarringInfo_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.uac_BarringInfoSetList_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode uac_BarringInfoSetList_v1700", err)
	}
	return nil
}

func (ie *SIB1_v1700_IEs_uac_BarringInfo_v1700) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.uac_BarringInfoSetList_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode uac_BarringInfoSetList_v1700", err)
	}
	return nil
}
