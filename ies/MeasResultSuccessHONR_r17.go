package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17 struct {
	measResult_r17 *MeasResultSuccessHONR_r17_measResult_r17 `optional`
}

func (ie *MeasResultSuccessHONR_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResult_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResult_r17 != nil {
		if err = ie.measResult_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measResult_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17) Decode(r *uper.UperReader) error {
	var err error
	var measResult_r17Present bool
	if measResult_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResult_r17Present {
		ie.measResult_r17 = new(MeasResultSuccessHONR_r17_measResult_r17)
		if err = ie.measResult_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measResult_r17", err)
		}
	}
	return nil
}
