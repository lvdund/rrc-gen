package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectRxTxDiff_r17 struct {
	dl_Ref_r17 *MeasObjectRxTxDiff_r17_dl_Ref_r17 `optional`
}

func (ie *MeasObjectRxTxDiff_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_Ref_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_Ref_r17 != nil {
		if err = ie.dl_Ref_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_Ref_r17", err)
		}
	}
	return nil
}

func (ie *MeasObjectRxTxDiff_r17) Decode(r *uper.UperReader) error {
	var err error
	var dl_Ref_r17Present bool
	if dl_Ref_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_Ref_r17Present {
		ie.dl_Ref_r17 = new(MeasObjectRxTxDiff_r17_dl_Ref_r17)
		if err = ie.dl_Ref_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_Ref_r17", err)
		}
	}
	return nil
}
