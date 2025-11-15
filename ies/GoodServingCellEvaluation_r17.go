package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GoodServingCellEvaluation_r17 struct {
	offset_r17 *GoodServingCellEvaluation_r17_offset_r17 `optional`
}

func (ie *GoodServingCellEvaluation_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.offset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.offset_r17 != nil {
		if err = ie.offset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode offset_r17", err)
		}
	}
	return nil
}

func (ie *GoodServingCellEvaluation_r17) Decode(r *uper.UperReader) error {
	var err error
	var offset_r17Present bool
	if offset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if offset_r17Present {
		ie.offset_r17 = new(GoodServingCellEvaluation_r17_offset_r17)
		if err = ie.offset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode offset_r17", err)
		}
	}
	return nil
}
