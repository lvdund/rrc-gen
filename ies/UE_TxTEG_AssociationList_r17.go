package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_TxTEG_AssociationList_r17 struct {
	Value []UE_TxTEG_Association_r17 `lb:1,ub:maxNrOfTxTEGReport_r17,madatory`
}

func (ie *UE_TxTEG_AssociationList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UE_TxTEG_Association_r17]([]*UE_TxTEG_Association_r17{}, uper.Constraint{Lb: 1, Ub: maxNrOfTxTEGReport_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UE_TxTEG_AssociationList_r17", err)
	}
	return nil
}

func (ie *UE_TxTEG_AssociationList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UE_TxTEG_Association_r17]([]*UE_TxTEG_Association_r17{}, uper.Constraint{Lb: 1, Ub: maxNrOfTxTEGReport_r17}, false)
	fn := func() *UE_TxTEG_Association_r17 {
		return new(UE_TxTEG_Association_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UE_TxTEG_AssociationList_r17", err)
	}
	ie.Value = []UE_TxTEG_Association_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
