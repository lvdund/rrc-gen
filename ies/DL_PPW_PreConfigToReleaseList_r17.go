package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_PPW_PreConfigToReleaseList_r17 struct {
	Value []DL_PPW_ID_r17 `lb:1,ub:maxNrofPPW_Config_r17,madatory`
}

func (ie *DL_PPW_PreConfigToReleaseList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*DL_PPW_ID_r17]([]*DL_PPW_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPPW_Config_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode DL_PPW_PreConfigToReleaseList_r17", err)
	}
	return nil
}

func (ie *DL_PPW_PreConfigToReleaseList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*DL_PPW_ID_r17]([]*DL_PPW_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPPW_Config_r17}, false)
	fn := func() *DL_PPW_ID_r17 {
		return new(DL_PPW_ID_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode DL_PPW_PreConfigToReleaseList_r17", err)
	}
	ie.Value = []DL_PPW_ID_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
