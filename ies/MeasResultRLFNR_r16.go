package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16 struct {
	measResult_r16 *MeasResultRLFNR_r16_measResult_r16 `lb:64,ub:64,optional`
}

func (ie *MeasResultRLFNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResult_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResult_r16 != nil {
		if err = ie.measResult_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResult_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResult_r16Present bool
	if measResult_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResult_r16Present {
		ie.measResult_r16 = new(MeasResultRLFNR_r16_measResult_r16)
		if err = ie.measResult_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResult_r16", err)
		}
	}
	return nil
}
