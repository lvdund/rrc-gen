package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_ActivatedConfig_r17 struct {
	pdcch_TCI_r17 []TCI_StateId  `lb:1,ub:5,madatory`
	pdsch_TCI_r17 uper.BitString `lb:1,ub:maxNrofTCI_States,madatory`
}

func (ie *TCI_ActivatedConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_pdcch_TCI_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	for _, i := range ie.pdcch_TCI_r17 {
		tmp_pdcch_TCI_r17.Value = append(tmp_pdcch_TCI_r17.Value, &i)
	}
	if err = tmp_pdcch_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_TCI_r17", err)
	}
	if err = w.WriteBitString(ie.pdsch_TCI_r17.Bytes, uint(ie.pdsch_TCI_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false); err != nil {
		return utils.WrapError("WriteBitString pdsch_TCI_r17", err)
	}
	return nil
}

func (ie *TCI_ActivatedConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_pdcch_TCI_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	fn_pdcch_TCI_r17 := func() *TCI_StateId {
		return new(TCI_StateId)
	}
	if err = tmp_pdcch_TCI_r17.Decode(r, fn_pdcch_TCI_r17); err != nil {
		return utils.WrapError("Decode pdcch_TCI_r17", err)
	}
	ie.pdcch_TCI_r17 = []TCI_StateId{}
	for _, i := range tmp_pdcch_TCI_r17.Value {
		ie.pdcch_TCI_r17 = append(ie.pdcch_TCI_r17, *i)
	}
	var tmp_bs_pdsch_TCI_r17 []byte
	var n_pdsch_TCI_r17 uint
	if tmp_bs_pdsch_TCI_r17, n_pdsch_TCI_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false); err != nil {
		return utils.WrapError("ReadBitString pdsch_TCI_r17", err)
	}
	ie.pdsch_TCI_r17 = uper.BitString{
		Bytes:   tmp_bs_pdsch_TCI_r17,
		NumBits: uint64(n_pdsch_TCI_r17),
	}
	return nil
}
