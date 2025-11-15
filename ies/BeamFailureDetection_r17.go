package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureDetection_r17 struct {
	failureDetectionSet1_r17 *BeamFailureDetectionSet_r17 `optional`
	failureDetectionSet2_r17 *BeamFailureDetectionSet_r17 `optional`
	additionalPCI_r17        *AdditionalPCIIndex_r17      `optional`
}

func (ie *BeamFailureDetection_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.failureDetectionSet1_r17 != nil, ie.failureDetectionSet2_r17 != nil, ie.additionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.failureDetectionSet1_r17 != nil {
		if err = ie.failureDetectionSet1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode failureDetectionSet1_r17", err)
		}
	}
	if ie.failureDetectionSet2_r17 != nil {
		if err = ie.failureDetectionSet2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode failureDetectionSet2_r17", err)
		}
	}
	if ie.additionalPCI_r17 != nil {
		if err = ie.additionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *BeamFailureDetection_r17) Decode(r *uper.UperReader) error {
	var err error
	var failureDetectionSet1_r17Present bool
	if failureDetectionSet1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var failureDetectionSet2_r17Present bool
	if failureDetectionSet2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalPCI_r17Present bool
	if additionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if failureDetectionSet1_r17Present {
		ie.failureDetectionSet1_r17 = new(BeamFailureDetectionSet_r17)
		if err = ie.failureDetectionSet1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode failureDetectionSet1_r17", err)
		}
	}
	if failureDetectionSet2_r17Present {
		ie.failureDetectionSet2_r17 = new(BeamFailureDetectionSet_r17)
		if err = ie.failureDetectionSet2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode failureDetectionSet2_r17", err)
		}
	}
	if additionalPCI_r17Present {
		ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.additionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPCI_r17", err)
		}
	}
	return nil
}
