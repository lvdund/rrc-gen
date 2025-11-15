package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DownlinkPreemption struct {
	int_RNTI                        RNTI_Value                          `madatory`
	timeFrequencySet                DownlinkPreemption_timeFrequencySet `madatory`
	dci_PayloadSize                 int64                               `lb:0,ub:maxINT_DCI_PayloadSize,madatory`
	int_ConfigurationPerServingCell []INT_ConfigurationPerServingCell   `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *DownlinkPreemption) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.int_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode int_RNTI", err)
	}
	if err = ie.timeFrequencySet.Encode(w); err != nil {
		return utils.WrapError("Encode timeFrequencySet", err)
	}
	if err = w.WriteInteger(ie.dci_PayloadSize, &uper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("WriteInteger dci_PayloadSize", err)
	}
	tmp_int_ConfigurationPerServingCell := utils.NewSequence[*INT_ConfigurationPerServingCell]([]*INT_ConfigurationPerServingCell{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.int_ConfigurationPerServingCell {
		tmp_int_ConfigurationPerServingCell.Value = append(tmp_int_ConfigurationPerServingCell.Value, &i)
	}
	if err = tmp_int_ConfigurationPerServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode int_ConfigurationPerServingCell", err)
	}
	return nil
}

func (ie *DownlinkPreemption) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.int_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode int_RNTI", err)
	}
	if err = ie.timeFrequencySet.Decode(r); err != nil {
		return utils.WrapError("Decode timeFrequencySet", err)
	}
	var tmp_int_dci_PayloadSize int64
	if tmp_int_dci_PayloadSize, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("ReadInteger dci_PayloadSize", err)
	}
	ie.dci_PayloadSize = tmp_int_dci_PayloadSize
	tmp_int_ConfigurationPerServingCell := utils.NewSequence[*INT_ConfigurationPerServingCell]([]*INT_ConfigurationPerServingCell{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_int_ConfigurationPerServingCell := func() *INT_ConfigurationPerServingCell {
		return new(INT_ConfigurationPerServingCell)
	}
	if err = tmp_int_ConfigurationPerServingCell.Decode(r, fn_int_ConfigurationPerServingCell); err != nil {
		return utils.WrapError("Decode int_ConfigurationPerServingCell", err)
	}
	ie.int_ConfigurationPerServingCell = []INT_ConfigurationPerServingCell{}
	for _, i := range tmp_int_ConfigurationPerServingCell.Value {
		ie.int_ConfigurationPerServingCell = append(ie.int_ConfigurationPerServingCell, *i)
	}
	return nil
}
