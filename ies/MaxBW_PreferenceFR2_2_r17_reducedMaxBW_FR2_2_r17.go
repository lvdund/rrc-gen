package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17 struct {
	reducedBW_FR2_2_DL_r17 *ReducedAggregatedBandwidth_r17 `optional`
	reducedBW_FR2_2_UL_r17 *ReducedAggregatedBandwidth_r17 `optional`
}

func (ie *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedBW_FR2_2_DL_r17 != nil, ie.reducedBW_FR2_2_UL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedBW_FR2_2_DL_r17 != nil {
		if err = ie.reducedBW_FR2_2_DL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reducedBW_FR2_2_DL_r17", err)
		}
	}
	if ie.reducedBW_FR2_2_UL_r17 != nil {
		if err = ie.reducedBW_FR2_2_UL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reducedBW_FR2_2_UL_r17", err)
		}
	}
	return nil
}

func (ie *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var reducedBW_FR2_2_DL_r17Present bool
	if reducedBW_FR2_2_DL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedBW_FR2_2_UL_r17Present bool
	if reducedBW_FR2_2_UL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedBW_FR2_2_DL_r17Present {
		ie.reducedBW_FR2_2_DL_r17 = new(ReducedAggregatedBandwidth_r17)
		if err = ie.reducedBW_FR2_2_DL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reducedBW_FR2_2_DL_r17", err)
		}
	}
	if reducedBW_FR2_2_UL_r17Present {
		ie.reducedBW_FR2_2_UL_r17 = new(ReducedAggregatedBandwidth_r17)
		if err = ie.reducedBW_FR2_2_UL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reducedBW_FR2_2_UL_r17", err)
		}
	}
	return nil
}
