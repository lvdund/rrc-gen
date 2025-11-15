package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPattern_patternType_bitmaps struct {
	resourceBlocks         uper.BitString                                              `lb:275,ub:275,madatory`
	symbolsInResourceBlock RateMatchPattern_patternType_bitmaps_symbolsInResourceBlock `lb:14,ub:14,madatory`
	periodicityAndPattern  *RateMatchPattern_patternType_bitmaps_periodicityAndPattern `lb:2,ub:2,optional`
}

func (ie *RateMatchPattern_patternType_bitmaps) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndPattern != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.resourceBlocks.Bytes, uint(ie.resourceBlocks.NumBits), &uper.Constraint{Lb: 275, Ub: 275}, false); err != nil {
		return utils.WrapError("WriteBitString resourceBlocks", err)
	}
	if err = ie.symbolsInResourceBlock.Encode(w); err != nil {
		return utils.WrapError("Encode symbolsInResourceBlock", err)
	}
	if ie.periodicityAndPattern != nil {
		if err = ie.periodicityAndPattern.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndPattern", err)
		}
	}
	return nil
}

func (ie *RateMatchPattern_patternType_bitmaps) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndPatternPresent bool
	if periodicityAndPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_resourceBlocks []byte
	var n_resourceBlocks uint
	if tmp_bs_resourceBlocks, n_resourceBlocks, err = r.ReadBitString(&uper.Constraint{Lb: 275, Ub: 275}, false); err != nil {
		return utils.WrapError("ReadBitString resourceBlocks", err)
	}
	ie.resourceBlocks = uper.BitString{
		Bytes:   tmp_bs_resourceBlocks,
		NumBits: uint64(n_resourceBlocks),
	}
	if err = ie.symbolsInResourceBlock.Decode(r); err != nil {
		return utils.WrapError("Decode symbolsInResourceBlock", err)
	}
	if periodicityAndPatternPresent {
		ie.periodicityAndPattern = new(RateMatchPattern_patternType_bitmaps_periodicityAndPattern)
		if err = ie.periodicityAndPattern.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndPattern", err)
		}
	}
	return nil
}
