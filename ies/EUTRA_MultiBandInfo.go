package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MultiBandInfo struct {
	eutra_FreqBandIndicator FreqBandIndicatorEUTRA `madatory`
	eutra_NS_PmaxList       *EUTRA_NS_PmaxList     `optional`
}

func (ie *EUTRA_MultiBandInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.eutra_NS_PmaxList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.eutra_FreqBandIndicator.Encode(w); err != nil {
		return utils.WrapError("Encode eutra_FreqBandIndicator", err)
	}
	if ie.eutra_NS_PmaxList != nil {
		if err = ie.eutra_NS_PmaxList.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_NS_PmaxList", err)
		}
	}
	return nil
}

func (ie *EUTRA_MultiBandInfo) Decode(r *uper.UperReader) error {
	var err error
	var eutra_NS_PmaxListPresent bool
	if eutra_NS_PmaxListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eutra_FreqBandIndicator.Decode(r); err != nil {
		return utils.WrapError("Decode eutra_FreqBandIndicator", err)
	}
	if eutra_NS_PmaxListPresent {
		ie.eutra_NS_PmaxList = new(EUTRA_NS_PmaxList)
		if err = ie.eutra_NS_PmaxList.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_NS_PmaxList", err)
		}
	}
	return nil
}
