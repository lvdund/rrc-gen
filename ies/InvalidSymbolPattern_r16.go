package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InvalidSymbolPattern_r16 struct {
	symbols_r16               InvalidSymbolPattern_r16_symbols_r16                `lb:14,ub:14,madatory`
	periodicityAndPattern_r16 *InvalidSymbolPattern_r16_periodicityAndPattern_r16 `lb:2,ub:2,optional`
}

func (ie *InvalidSymbolPattern_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndPattern_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.symbols_r16.Encode(w); err != nil {
		return utils.WrapError("Encode symbols_r16", err)
	}
	if ie.periodicityAndPattern_r16 != nil {
		if err = ie.periodicityAndPattern_r16.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndPattern_r16", err)
		}
	}
	return nil
}

func (ie *InvalidSymbolPattern_r16) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndPattern_r16Present bool
	if periodicityAndPattern_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.symbols_r16.Decode(r); err != nil {
		return utils.WrapError("Decode symbols_r16", err)
	}
	if periodicityAndPattern_r16Present {
		ie.periodicityAndPattern_r16 = new(InvalidSymbolPattern_r16_periodicityAndPattern_r16)
		if err = ie.periodicityAndPattern_r16.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndPattern_r16", err)
		}
	}
	return nil
}
