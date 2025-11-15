package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_MultiBandInfo struct {
	freqBandIndicatorNR *FreqBandIndicatorNR `optional`
	nr_NS_PmaxList      *NR_NS_PmaxList      `optional`
}

func (ie *NR_MultiBandInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.freqBandIndicatorNR != nil, ie.nr_NS_PmaxList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.freqBandIndicatorNR != nil {
		if err = ie.freqBandIndicatorNR.Encode(w); err != nil {
			return utils.WrapError("Encode freqBandIndicatorNR", err)
		}
	}
	if ie.nr_NS_PmaxList != nil {
		if err = ie.nr_NS_PmaxList.Encode(w); err != nil {
			return utils.WrapError("Encode nr_NS_PmaxList", err)
		}
	}
	return nil
}

func (ie *NR_MultiBandInfo) Decode(r *uper.UperReader) error {
	var err error
	var freqBandIndicatorNRPresent bool
	if freqBandIndicatorNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nr_NS_PmaxListPresent bool
	if nr_NS_PmaxListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if freqBandIndicatorNRPresent {
		ie.freqBandIndicatorNR = new(FreqBandIndicatorNR)
		if err = ie.freqBandIndicatorNR.Decode(r); err != nil {
			return utils.WrapError("Decode freqBandIndicatorNR", err)
		}
	}
	if nr_NS_PmaxListPresent {
		ie.nr_NS_PmaxList = new(NR_NS_PmaxList)
		if err = ie.nr_NS_PmaxList.Decode(r); err != nil {
			return utils.WrapError("Decode nr_NS_PmaxList", err)
		}
	}
	return nil
}
