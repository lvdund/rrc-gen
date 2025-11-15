package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SelectedBandEntriesMN struct {
	Value []BandEntryIndex `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *SelectedBandEntriesMN) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandEntryIndex]([]*BandEntryIndex{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SelectedBandEntriesMN", err)
	}
	return nil
}

func (ie *SelectedBandEntriesMN) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandEntryIndex]([]*BandEntryIndex{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandEntryIndex {
		return new(BandEntryIndex)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SelectedBandEntriesMN", err)
	}
	ie.Value = []BandEntryIndex{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
