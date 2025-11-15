package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig struct {
	schedulingCellInfo          CrossCarrierSchedulingConfig_schedulingCellInfo           `lb:1,ub:7,madatory`
	carrierIndicatorSize_r16    *CrossCarrierSchedulingConfig_carrierIndicatorSize_r16    `lb:0,ub:3,optional,ext-1`
	enableDefaultBeamForCCS_r16 *CrossCarrierSchedulingConfig_enableDefaultBeamForCCS_r16 `optional,ext-1`
	ccs_BlindDetectionSplit_r17 *CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17 `optional,ext-2`
}

func (ie *CrossCarrierSchedulingConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.carrierIndicatorSize_r16 != nil || ie.enableDefaultBeamForCCS_r16 != nil || ie.ccs_BlindDetectionSplit_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.schedulingCellInfo.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingCellInfo", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.carrierIndicatorSize_r16 != nil || ie.enableDefaultBeamForCCS_r16 != nil, ie.ccs_BlindDetectionSplit_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CrossCarrierSchedulingConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.carrierIndicatorSize_r16 != nil, ie.enableDefaultBeamForCCS_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode carrierIndicatorSize_r16 optional
			if ie.carrierIndicatorSize_r16 != nil {
				if err = ie.carrierIndicatorSize_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode carrierIndicatorSize_r16", err)
				}
			}
			// encode enableDefaultBeamForCCS_r16 optional
			if ie.enableDefaultBeamForCCS_r16 != nil {
				if err = ie.enableDefaultBeamForCCS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableDefaultBeamForCCS_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.ccs_BlindDetectionSplit_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ccs_BlindDetectionSplit_r17 optional
			if ie.ccs_BlindDetectionSplit_r17 != nil {
				if err = ie.ccs_BlindDetectionSplit_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ccs_BlindDetectionSplit_r17", err)
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

func (ie *CrossCarrierSchedulingConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.schedulingCellInfo.Decode(r); err != nil {
		return utils.WrapError("Decode schedulingCellInfo", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			carrierIndicatorSize_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableDefaultBeamForCCS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode carrierIndicatorSize_r16 optional
			if carrierIndicatorSize_r16Present {
				ie.carrierIndicatorSize_r16 = new(CrossCarrierSchedulingConfig_carrierIndicatorSize_r16)
				if err = ie.carrierIndicatorSize_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode carrierIndicatorSize_r16", err)
				}
			}
			// decode enableDefaultBeamForCCS_r16 optional
			if enableDefaultBeamForCCS_r16Present {
				ie.enableDefaultBeamForCCS_r16 = new(CrossCarrierSchedulingConfig_enableDefaultBeamForCCS_r16)
				if err = ie.enableDefaultBeamForCCS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableDefaultBeamForCCS_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ccs_BlindDetectionSplit_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ccs_BlindDetectionSplit_r17 optional
			if ccs_BlindDetectionSplit_r17Present {
				ie.ccs_BlindDetectionSplit_r17 = new(CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17)
				if err = ie.ccs_BlindDetectionSplit_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ccs_BlindDetectionSplit_r17", err)
				}
			}
		}
	}
	return nil
}
