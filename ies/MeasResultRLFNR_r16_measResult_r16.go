package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16 struct {
	cellResults_r16    *MeasResultRLFNR_r16_measResult_r16_cellResults_r16    `optional`
	rsIndexResults_r16 *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16 `lb:64,ub:64,optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellResults_r16 != nil, ie.rsIndexResults_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellResults_r16 != nil {
		if err = ie.cellResults_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellResults_r16", err)
		}
	}
	if ie.rsIndexResults_r16 != nil {
		if err = ie.rsIndexResults_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsIndexResults_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var cellResults_r16Present bool
	if cellResults_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rsIndexResults_r16Present bool
	if rsIndexResults_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cellResults_r16Present {
		ie.cellResults_r16 = new(MeasResultRLFNR_r16_measResult_r16_cellResults_r16)
		if err = ie.cellResults_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cellResults_r16", err)
		}
	}
	if rsIndexResults_r16Present {
		ie.rsIndexResults_r16 = new(MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16)
		if err = ie.rsIndexResults_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsIndexResults_r16", err)
		}
	}
	return nil
}
