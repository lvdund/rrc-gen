package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig_transformPrecoderDisabled struct {
	frequencyDensity      []int64                                                            `lb:2,ub:2,e_lb:0,e_ub:0,optional`
	timeDensity           []int64                                                            `lb:3,ub:3,e_lb:0,e_ub:0,optional`
	maxNrofPorts          PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts           `madatory`
	resourceElementOffset *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset `optional`
	ptrs_Power            PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power             `madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.frequencyDensity) > 0, len(ie.timeDensity) > 0, ie.resourceElementOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.frequencyDensity) > 0 {
		tmp_frequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.frequencyDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_frequencyDensity.Value = append(tmp_frequencyDensity.Value, &tmp_ie)
		}
		if err = tmp_frequencyDensity.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyDensity", err)
		}
	}
	if len(ie.timeDensity) > 0 {
		tmp_timeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.timeDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_timeDensity.Value = append(tmp_timeDensity.Value, &tmp_ie)
		}
		if err = tmp_timeDensity.Encode(w); err != nil {
			return utils.WrapError("Encode timeDensity", err)
		}
	}
	if err = ie.maxNrofPorts.Encode(w); err != nil {
		return utils.WrapError("Encode maxNrofPorts", err)
	}
	if ie.resourceElementOffset != nil {
		if err = ie.resourceElementOffset.Encode(w); err != nil {
			return utils.WrapError("Encode resourceElementOffset", err)
		}
	}
	if err = ie.ptrs_Power.Encode(w); err != nil {
		return utils.WrapError("Encode ptrs_Power", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled) Decode(r *uper.UperReader) error {
	var err error
	var frequencyDensityPresent bool
	if frequencyDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDensityPresent bool
	if timeDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceElementOffsetPresent bool
	if resourceElementOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyDensityPresent {
		tmp_frequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_frequencyDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_frequencyDensity.Decode(r, fn_frequencyDensity); err != nil {
			return utils.WrapError("Decode frequencyDensity", err)
		}
		ie.frequencyDensity = []int64{}
		for _, i := range tmp_frequencyDensity.Value {
			ie.frequencyDensity = append(ie.frequencyDensity, int64(i.Value))
		}
	}
	if timeDensityPresent {
		tmp_timeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_timeDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_timeDensity.Decode(r, fn_timeDensity); err != nil {
			return utils.WrapError("Decode timeDensity", err)
		}
		ie.timeDensity = []int64{}
		for _, i := range tmp_timeDensity.Value {
			ie.timeDensity = append(ie.timeDensity, int64(i.Value))
		}
	}
	if err = ie.maxNrofPorts.Decode(r); err != nil {
		return utils.WrapError("Decode maxNrofPorts", err)
	}
	if resourceElementOffsetPresent {
		ie.resourceElementOffset = new(PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset)
		if err = ie.resourceElementOffset.Decode(r); err != nil {
			return utils.WrapError("Decode resourceElementOffset", err)
		}
	}
	if err = ie.ptrs_Power.Decode(r); err != nil {
		return utils.WrapError("Decode ptrs_Power", err)
	}
	return nil
}
