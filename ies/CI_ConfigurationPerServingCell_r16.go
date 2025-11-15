package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CI_ConfigurationPerServingCell_r16 struct {
	servingCellId                    ServCellIndex                                                        `madatory`
	positionInDCI_r16                int64                                                                `lb:0,ub:maxCI_DCI_PayloadSize_1_r16,madatory`
	positionInDCI_ForSUL_r16         *int64                                                               `lb:0,ub:maxCI_DCI_PayloadSize_1_r16,optional`
	ci_PayloadSize_r16               CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16                `madatory`
	timeFrequencyRegion_r16          *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16          `lb:0,ub:37949,optional`
	uplinkCancellationPriority_v1610 *CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610 `optional,ext`
}

func (ie *CI_ConfigurationPerServingCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.positionInDCI_ForSUL_r16 != nil, ie.timeFrequencyRegion_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode servingCellId", err)
	}
	if err = w.WriteInteger(ie.positionInDCI_r16, &uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
		return utils.WrapError("WriteInteger positionInDCI_r16", err)
	}
	if ie.positionInDCI_ForSUL_r16 != nil {
		if err = w.WriteInteger(*ie.positionInDCI_ForSUL_r16, &uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Encode positionInDCI_ForSUL_r16", err)
		}
	}
	if err = ie.ci_PayloadSize_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ci_PayloadSize_r16", err)
	}
	if ie.timeFrequencyRegion_r16 != nil {
		if err = ie.timeFrequencyRegion_r16.Encode(w); err != nil {
			return utils.WrapError("Encode timeFrequencyRegion_r16", err)
		}
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16) Decode(r *uper.UperReader) error {
	var err error
	var positionInDCI_ForSUL_r16Present bool
	if positionInDCI_ForSUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeFrequencyRegion_r16Present bool
	if timeFrequencyRegion_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode servingCellId", err)
	}
	var tmp_int_positionInDCI_r16 int64
	if tmp_int_positionInDCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
		return utils.WrapError("ReadInteger positionInDCI_r16", err)
	}
	ie.positionInDCI_r16 = tmp_int_positionInDCI_r16
	if positionInDCI_ForSUL_r16Present {
		var tmp_int_positionInDCI_ForSUL_r16 int64
		if tmp_int_positionInDCI_ForSUL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Decode positionInDCI_ForSUL_r16", err)
		}
		ie.positionInDCI_ForSUL_r16 = &tmp_int_positionInDCI_ForSUL_r16
	}
	if err = ie.ci_PayloadSize_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ci_PayloadSize_r16", err)
	}
	if timeFrequencyRegion_r16Present {
		ie.timeFrequencyRegion_r16 = new(CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16)
		if err = ie.timeFrequencyRegion_r16.Decode(r); err != nil {
			return utils.WrapError("Decode timeFrequencyRegion_r16", err)
		}
	}
	return nil
}
