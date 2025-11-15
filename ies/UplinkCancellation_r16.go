package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkCancellation_r16 struct {
	ci_RNTI_r16                        RNTI_Value                           `madatory`
	dci_PayloadSizeForCI_r16           int64                                `lb:0,ub:maxCI_DCI_PayloadSize_r16,madatory`
	ci_ConfigurationPerServingCell_r16 []CI_ConfigurationPerServingCell_r16 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *UplinkCancellation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ci_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ci_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.dci_PayloadSizeForCI_r16, &uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("WriteInteger dci_PayloadSizeForCI_r16", err)
	}
	tmp_ci_ConfigurationPerServingCell_r16 := utils.NewSequence[*CI_ConfigurationPerServingCell_r16]([]*CI_ConfigurationPerServingCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.ci_ConfigurationPerServingCell_r16 {
		tmp_ci_ConfigurationPerServingCell_r16.Value = append(tmp_ci_ConfigurationPerServingCell_r16.Value, &i)
	}
	if err = tmp_ci_ConfigurationPerServingCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ci_ConfigurationPerServingCell_r16", err)
	}
	return nil
}

func (ie *UplinkCancellation_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ci_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ci_RNTI_r16", err)
	}
	var tmp_int_dci_PayloadSizeForCI_r16 int64
	if tmp_int_dci_PayloadSizeForCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("ReadInteger dci_PayloadSizeForCI_r16", err)
	}
	ie.dci_PayloadSizeForCI_r16 = tmp_int_dci_PayloadSizeForCI_r16
	tmp_ci_ConfigurationPerServingCell_r16 := utils.NewSequence[*CI_ConfigurationPerServingCell_r16]([]*CI_ConfigurationPerServingCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_ci_ConfigurationPerServingCell_r16 := func() *CI_ConfigurationPerServingCell_r16 {
		return new(CI_ConfigurationPerServingCell_r16)
	}
	if err = tmp_ci_ConfigurationPerServingCell_r16.Decode(r, fn_ci_ConfigurationPerServingCell_r16); err != nil {
		return utils.WrapError("Decode ci_ConfigurationPerServingCell_r16", err)
	}
	ie.ci_ConfigurationPerServingCell_r16 = []CI_ConfigurationPerServingCell_r16{}
	for _, i := range tmp_ci_ConfigurationPerServingCell_r16.Value {
		ie.ci_ConfigurationPerServingCell_r16 = append(ie.ci_ConfigurationPerServingCell_r16, *i)
	}
	return nil
}
