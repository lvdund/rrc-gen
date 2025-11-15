package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_HARQ_ACK_EnhType3_r17 struct {
	pdsch_HARQ_ACK_EnhType3Index_r17 PDSCH_HARQ_ACK_EnhType3Index_r17                            `madatory`
	applicable_r17                   []int64                                                     `lb:1,ub:maxNrofServingCells,e_lb:0,e_ub:1,madatory`
	pdsch_HARQ_ACK_EnhType3NDI_r17   *PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3NDI_r17 `optional`
	pdsch_HARQ_ACK_EnhType3CBG_r17   *PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3CBG_r17 `optional`
}

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdsch_HARQ_ACK_EnhType3NDI_r17 != nil, ie.pdsch_HARQ_ACK_EnhType3CBG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pdsch_HARQ_ACK_EnhType3Index_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3Index_r17", err)
	}
	tmp_applicable_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.applicable_r17 {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 1}, false)
		tmp_applicable_r17.Value = append(tmp_applicable_r17.Value, &tmp_ie)
	}
	if err = tmp_applicable_r17.Encode(w); err != nil {
		return utils.WrapError("Encode applicable_r17", err)
	}
	if ie.pdsch_HARQ_ACK_EnhType3NDI_r17 != nil {
		if err = ie.pdsch_HARQ_ACK_EnhType3NDI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3NDI_r17", err)
		}
	}
	if ie.pdsch_HARQ_ACK_EnhType3CBG_r17 != nil {
		if err = ie.pdsch_HARQ_ACK_EnhType3CBG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3CBG_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdsch_HARQ_ACK_EnhType3NDI_r17Present bool
	if pdsch_HARQ_ACK_EnhType3NDI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_HARQ_ACK_EnhType3CBG_r17Present bool
	if pdsch_HARQ_ACK_EnhType3CBG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pdsch_HARQ_ACK_EnhType3Index_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3Index_r17", err)
	}
	tmp_applicable_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_applicable_r17 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 1}, false)
		return &ie
	}
	if err = tmp_applicable_r17.Decode(r, fn_applicable_r17); err != nil {
		return utils.WrapError("Decode applicable_r17", err)
	}
	ie.applicable_r17 = []int64{}
	for _, i := range tmp_applicable_r17.Value {
		ie.applicable_r17 = append(ie.applicable_r17, int64(i.Value))
	}
	if pdsch_HARQ_ACK_EnhType3NDI_r17Present {
		ie.pdsch_HARQ_ACK_EnhType3NDI_r17 = new(PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3NDI_r17)
		if err = ie.pdsch_HARQ_ACK_EnhType3NDI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3NDI_r17", err)
		}
	}
	if pdsch_HARQ_ACK_EnhType3CBG_r17Present {
		ie.pdsch_HARQ_ACK_EnhType3CBG_r17 = new(PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3CBG_r17)
		if err = ie.pdsch_HARQ_ACK_EnhType3CBG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3CBG_r17", err)
		}
	}
	return nil
}
