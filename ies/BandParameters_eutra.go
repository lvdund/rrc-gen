package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_eutra struct {
	bandEUTRA                 FreqBandIndicatorEUTRA  `madatory`
	ca_BandwidthClassDL_EUTRA *CA_BandwidthClassEUTRA `optional`
	ca_BandwidthClassUL_EUTRA *CA_BandwidthClassEUTRA `optional`
}

func (ie *BandParameters_eutra) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_BandwidthClassDL_EUTRA != nil, ie.ca_BandwidthClassUL_EUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bandEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode bandEUTRA", err)
	}
	if ie.ca_BandwidthClassDL_EUTRA != nil {
		if err = ie.ca_BandwidthClassDL_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ca_BandwidthClassDL_EUTRA", err)
		}
	}
	if ie.ca_BandwidthClassUL_EUTRA != nil {
		if err = ie.ca_BandwidthClassUL_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ca_BandwidthClassUL_EUTRA", err)
		}
	}
	return nil
}

func (ie *BandParameters_eutra) Decode(r *uper.UperReader) error {
	var err error
	var ca_BandwidthClassDL_EUTRAPresent bool
	if ca_BandwidthClassDL_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_BandwidthClassUL_EUTRAPresent bool
	if ca_BandwidthClassUL_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bandEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode bandEUTRA", err)
	}
	if ca_BandwidthClassDL_EUTRAPresent {
		ie.ca_BandwidthClassDL_EUTRA = new(CA_BandwidthClassEUTRA)
		if err = ie.ca_BandwidthClassDL_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ca_BandwidthClassDL_EUTRA", err)
		}
	}
	if ca_BandwidthClassUL_EUTRAPresent {
		ie.ca_BandwidthClassUL_EUTRA = new(CA_BandwidthClassEUTRA)
		if err = ie.ca_BandwidthClassUL_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ca_BandwidthClassUL_EUTRA", err)
		}
	}
	return nil
}
