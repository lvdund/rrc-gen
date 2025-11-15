package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_v1610 struct {
	supportedCSI_RS_ResourceListAlt_r16 []int64 `lb:1,ub:maxNrofCSI_RS_Resources,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookParameters_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.supportedCSI_RS_ResourceListAlt_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.supportedCSI_RS_ResourceListAlt_r16) > 0 {
		tmp_supportedCSI_RS_ResourceListAlt_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
		for _, i := range ie.supportedCSI_RS_ResourceListAlt_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_supportedCSI_RS_ResourceListAlt_r16.Value = append(tmp_supportedCSI_RS_ResourceListAlt_r16.Value, &tmp_ie)
		}
		if err = tmp_supportedCSI_RS_ResourceListAlt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedCSI_RS_ResourceListAlt_r16", err)
		}
	}
	return nil
}

func (ie *CodebookParameters_v1610) Decode(r *uper.UperReader) error {
	var err error
	var supportedCSI_RS_ResourceListAlt_r16Present bool
	if supportedCSI_RS_ResourceListAlt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedCSI_RS_ResourceListAlt_r16Present {
		tmp_supportedCSI_RS_ResourceListAlt_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
		fn_supportedCSI_RS_ResourceListAlt_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_supportedCSI_RS_ResourceListAlt_r16.Decode(r, fn_supportedCSI_RS_ResourceListAlt_r16); err != nil {
			return utils.WrapError("Decode supportedCSI_RS_ResourceListAlt_r16", err)
		}
		ie.supportedCSI_RS_ResourceListAlt_r16 = []int64{}
		for _, i := range tmp_supportedCSI_RS_ResourceListAlt_r16.Value {
			ie.supportedCSI_RS_ResourceListAlt_r16 = append(ie.supportedCSI_RS_ResourceListAlt_r16, int64(i.Value))
		}
	}
	return nil
}
