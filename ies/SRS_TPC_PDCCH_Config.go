package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_TPC_PDCCH_Config struct {
	srs_CC_SetIndexlist []SRS_CC_SetIndex `lb:1,ub:4,optional`
}

func (ie *SRS_TPC_PDCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.srs_CC_SetIndexlist) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.srs_CC_SetIndexlist) > 0 {
		tmp_srs_CC_SetIndexlist := utils.NewSequence[*SRS_CC_SetIndex]([]*SRS_CC_SetIndex{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.srs_CC_SetIndexlist {
			tmp_srs_CC_SetIndexlist.Value = append(tmp_srs_CC_SetIndexlist.Value, &i)
		}
		if err = tmp_srs_CC_SetIndexlist.Encode(w); err != nil {
			return utils.WrapError("Encode srs_CC_SetIndexlist", err)
		}
	}
	return nil
}

func (ie *SRS_TPC_PDCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var srs_CC_SetIndexlistPresent bool
	if srs_CC_SetIndexlistPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_CC_SetIndexlistPresent {
		tmp_srs_CC_SetIndexlist := utils.NewSequence[*SRS_CC_SetIndex]([]*SRS_CC_SetIndex{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn_srs_CC_SetIndexlist := func() *SRS_CC_SetIndex {
			return new(SRS_CC_SetIndex)
		}
		if err = tmp_srs_CC_SetIndexlist.Decode(r, fn_srs_CC_SetIndexlist); err != nil {
			return utils.WrapError("Decode srs_CC_SetIndexlist", err)
		}
		ie.srs_CC_SetIndexlist = []SRS_CC_SetIndex{}
		for _, i := range tmp_srs_CC_SetIndexlist.Value {
			ie.srs_CC_SetIndexlist = append(ie.srs_CC_SetIndexlist, *i)
		}
	}
	return nil
}
