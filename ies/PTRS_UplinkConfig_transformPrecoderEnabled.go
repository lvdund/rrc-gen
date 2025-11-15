package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig_transformPrecoderEnabled struct {
	sampleDensity                 []int64                                                                   `lb:5,ub:5,e_lb:0,e_ub:0,madatory`
	timeDensityTransformPrecoding *PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding `optional`
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.timeDensityTransformPrecoding != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_sampleDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 5, Ub: 5}, false)
	for _, i := range ie.sampleDensity {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
		tmp_sampleDensity.Value = append(tmp_sampleDensity.Value, &tmp_ie)
	}
	if err = tmp_sampleDensity.Encode(w); err != nil {
		return utils.WrapError("Encode sampleDensity", err)
	}
	if ie.timeDensityTransformPrecoding != nil {
		if err = ie.timeDensityTransformPrecoding.Encode(w); err != nil {
			return utils.WrapError("Encode timeDensityTransformPrecoding", err)
		}
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled) Decode(r *uper.UperReader) error {
	var err error
	var timeDensityTransformPrecodingPresent bool
	if timeDensityTransformPrecodingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_sampleDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 5, Ub: 5}, false)
	fn_sampleDensity := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
		return &ie
	}
	if err = tmp_sampleDensity.Decode(r, fn_sampleDensity); err != nil {
		return utils.WrapError("Decode sampleDensity", err)
	}
	ie.sampleDensity = []int64{}
	for _, i := range tmp_sampleDensity.Value {
		ie.sampleDensity = append(ie.sampleDensity, int64(i.Value))
	}
	if timeDensityTransformPrecodingPresent {
		ie.timeDensityTransformPrecoding = new(PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding)
		if err = ie.timeDensityTransformPrecoding.Decode(r); err != nil {
			return utils.WrapError("Decode timeDensityTransformPrecoding", err)
		}
	}
	return nil
}
