package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters struct {
	singleUL_Transmission         *MRDC_Parameters_singleUL_Transmission         `optional`
	dynamicPowerSharingENDC       *MRDC_Parameters_dynamicPowerSharingENDC       `optional`
	tdm_Pattern                   *MRDC_Parameters_tdm_Pattern                   `optional`
	ul_SharingEUTRA_NR            *MRDC_Parameters_ul_SharingEUTRA_NR            `optional`
	ul_SwitchingTimeEUTRA_NR      *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR      `optional`
	simultaneousRxTxInterBandENDC *MRDC_Parameters_simultaneousRxTxInterBandENDC `optional`
	asyncIntraBandENDC            *MRDC_Parameters_asyncIntraBandENDC            `optional`
	dualPA_Architecture           *MRDC_Parameters_dualPA_Architecture           `optional,ext-1`
	intraBandENDC_Support         *MRDC_Parameters_intraBandENDC_Support         `optional,ext-1`
	ul_TimingAlignmentEUTRA_NR    *MRDC_Parameters_ul_TimingAlignmentEUTRA_NR    `optional,ext-1`
}

func (ie *MRDC_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dualPA_Architecture != nil || ie.intraBandENDC_Support != nil || ie.ul_TimingAlignmentEUTRA_NR != nil
	preambleBits := []bool{hasExtensions, ie.singleUL_Transmission != nil, ie.dynamicPowerSharingENDC != nil, ie.tdm_Pattern != nil, ie.ul_SharingEUTRA_NR != nil, ie.ul_SwitchingTimeEUTRA_NR != nil, ie.simultaneousRxTxInterBandENDC != nil, ie.asyncIntraBandENDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.singleUL_Transmission != nil {
		if err = ie.singleUL_Transmission.Encode(w); err != nil {
			return utils.WrapError("Encode singleUL_Transmission", err)
		}
	}
	if ie.dynamicPowerSharingENDC != nil {
		if err = ie.dynamicPowerSharingENDC.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPowerSharingENDC", err)
		}
	}
	if ie.tdm_Pattern != nil {
		if err = ie.tdm_Pattern.Encode(w); err != nil {
			return utils.WrapError("Encode tdm_Pattern", err)
		}
	}
	if ie.ul_SharingEUTRA_NR != nil {
		if err = ie.ul_SharingEUTRA_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ul_SharingEUTRA_NR", err)
		}
	}
	if ie.ul_SwitchingTimeEUTRA_NR != nil {
		if err = ie.ul_SwitchingTimeEUTRA_NR.Encode(w); err != nil {
			return utils.WrapError("Encode ul_SwitchingTimeEUTRA_NR", err)
		}
	}
	if ie.simultaneousRxTxInterBandENDC != nil {
		if err = ie.simultaneousRxTxInterBandENDC.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxInterBandENDC", err)
		}
	}
	if ie.asyncIntraBandENDC != nil {
		if err = ie.asyncIntraBandENDC.Encode(w); err != nil {
			return utils.WrapError("Encode asyncIntraBandENDC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dualPA_Architecture != nil || ie.intraBandENDC_Support != nil || ie.ul_TimingAlignmentEUTRA_NR != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MRDC_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dualPA_Architecture != nil, ie.intraBandENDC_Support != nil, ie.ul_TimingAlignmentEUTRA_NR != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dualPA_Architecture optional
			if ie.dualPA_Architecture != nil {
				if err = ie.dualPA_Architecture.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dualPA_Architecture", err)
				}
			}
			// encode intraBandENDC_Support optional
			if ie.intraBandENDC_Support != nil {
				if err = ie.intraBandENDC_Support.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraBandENDC_Support", err)
				}
			}
			// encode ul_TimingAlignmentEUTRA_NR optional
			if ie.ul_TimingAlignmentEUTRA_NR != nil {
				if err = ie.ul_TimingAlignmentEUTRA_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_TimingAlignmentEUTRA_NR", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MRDC_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var singleUL_TransmissionPresent bool
	if singleUL_TransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPowerSharingENDCPresent bool
	if dynamicPowerSharingENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdm_PatternPresent bool
	if tdm_PatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_SharingEUTRA_NRPresent bool
	if ul_SharingEUTRA_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_SwitchingTimeEUTRA_NRPresent bool
	if ul_SwitchingTimeEUTRA_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousRxTxInterBandENDCPresent bool
	if simultaneousRxTxInterBandENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var asyncIntraBandENDCPresent bool
	if asyncIntraBandENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if singleUL_TransmissionPresent {
		ie.singleUL_Transmission = new(MRDC_Parameters_singleUL_Transmission)
		if err = ie.singleUL_Transmission.Decode(r); err != nil {
			return utils.WrapError("Decode singleUL_Transmission", err)
		}
	}
	if dynamicPowerSharingENDCPresent {
		ie.dynamicPowerSharingENDC = new(MRDC_Parameters_dynamicPowerSharingENDC)
		if err = ie.dynamicPowerSharingENDC.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicPowerSharingENDC", err)
		}
	}
	if tdm_PatternPresent {
		ie.tdm_Pattern = new(MRDC_Parameters_tdm_Pattern)
		if err = ie.tdm_Pattern.Decode(r); err != nil {
			return utils.WrapError("Decode tdm_Pattern", err)
		}
	}
	if ul_SharingEUTRA_NRPresent {
		ie.ul_SharingEUTRA_NR = new(MRDC_Parameters_ul_SharingEUTRA_NR)
		if err = ie.ul_SharingEUTRA_NR.Decode(r); err != nil {
			return utils.WrapError("Decode ul_SharingEUTRA_NR", err)
		}
	}
	if ul_SwitchingTimeEUTRA_NRPresent {
		ie.ul_SwitchingTimeEUTRA_NR = new(MRDC_Parameters_ul_SwitchingTimeEUTRA_NR)
		if err = ie.ul_SwitchingTimeEUTRA_NR.Decode(r); err != nil {
			return utils.WrapError("Decode ul_SwitchingTimeEUTRA_NR", err)
		}
	}
	if simultaneousRxTxInterBandENDCPresent {
		ie.simultaneousRxTxInterBandENDC = new(MRDC_Parameters_simultaneousRxTxInterBandENDC)
		if err = ie.simultaneousRxTxInterBandENDC.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxInterBandENDC", err)
		}
	}
	if asyncIntraBandENDCPresent {
		ie.asyncIntraBandENDC = new(MRDC_Parameters_asyncIntraBandENDC)
		if err = ie.asyncIntraBandENDC.Decode(r); err != nil {
			return utils.WrapError("Decode asyncIntraBandENDC", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			dualPA_ArchitecturePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraBandENDC_SupportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_TimingAlignmentEUTRA_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dualPA_Architecture optional
			if dualPA_ArchitecturePresent {
				ie.dualPA_Architecture = new(MRDC_Parameters_dualPA_Architecture)
				if err = ie.dualPA_Architecture.Decode(extReader); err != nil {
					return utils.WrapError("Decode dualPA_Architecture", err)
				}
			}
			// decode intraBandENDC_Support optional
			if intraBandENDC_SupportPresent {
				ie.intraBandENDC_Support = new(MRDC_Parameters_intraBandENDC_Support)
				if err = ie.intraBandENDC_Support.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraBandENDC_Support", err)
				}
			}
			// decode ul_TimingAlignmentEUTRA_NR optional
			if ul_TimingAlignmentEUTRA_NRPresent {
				ie.ul_TimingAlignmentEUTRA_NR = new(MRDC_Parameters_ul_TimingAlignmentEUTRA_NR)
				if err = ie.ul_TimingAlignmentEUTRA_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_TimingAlignmentEUTRA_NR", err)
				}
			}
		}
	}
	return nil
}
