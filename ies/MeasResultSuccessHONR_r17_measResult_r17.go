package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17_measResult_r17 struct {
	cellResults_r17    *MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17    `optional`
	rsIndexResults_r17 *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17 `optional`
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellResults_r17 != nil, ie.rsIndexResults_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellResults_r17 != nil {
		if err = ie.cellResults_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cellResults_r17", err)
		}
	}
	if ie.rsIndexResults_r17 != nil {
		if err = ie.rsIndexResults_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rsIndexResults_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17) Decode(r *uper.UperReader) error {
	var err error
	var cellResults_r17Present bool
	if cellResults_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rsIndexResults_r17Present bool
	if rsIndexResults_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cellResults_r17Present {
		ie.cellResults_r17 = new(MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17)
		if err = ie.cellResults_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cellResults_r17", err)
		}
	}
	if rsIndexResults_r17Present {
		ie.rsIndexResults_r17 = new(MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17)
		if err = ie.rsIndexResults_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rsIndexResults_r17", err)
		}
	}
	return nil
}
