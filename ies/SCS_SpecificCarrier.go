package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCS_SpecificCarrier struct {
	offsetToCarrier         int64             `lb:0,ub:2199,madatory`
	subcarrierSpacing       SubcarrierSpacing `madatory`
	carrierBandwidth        int64             `lb:1,ub:maxNrofPhysicalResourceBlocks,madatory`
	txDirectCurrentLocation *int64            `lb:0,ub:4095,optional,ext-1`
}

func (ie *SCS_SpecificCarrier) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.txDirectCurrentLocation != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.offsetToCarrier, &uper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("WriteInteger offsetToCarrier", err)
	}
	if err = ie.subcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing", err)
	}
	if err = w.WriteInteger(ie.carrierBandwidth, &uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("WriteInteger carrierBandwidth", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.txDirectCurrentLocation != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SCS_SpecificCarrier", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.txDirectCurrentLocation != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode txDirectCurrentLocation optional
			if ie.txDirectCurrentLocation != nil {
				if err = extWriter.WriteInteger(*ie.txDirectCurrentLocation, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
					return utils.WrapError("Encode txDirectCurrentLocation", err)
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

func (ie *SCS_SpecificCarrier) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_offsetToCarrier int64
	if tmp_int_offsetToCarrier, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("ReadInteger offsetToCarrier", err)
	}
	ie.offsetToCarrier = tmp_int_offsetToCarrier
	if err = ie.subcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing", err)
	}
	var tmp_int_carrierBandwidth int64
	if tmp_int_carrierBandwidth, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("ReadInteger carrierBandwidth", err)
	}
	ie.carrierBandwidth = tmp_int_carrierBandwidth

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

			txDirectCurrentLocationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode txDirectCurrentLocation optional
			if txDirectCurrentLocationPresent {
				var tmp_int_txDirectCurrentLocation int64
				if tmp_int_txDirectCurrentLocation, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
					return utils.WrapError("Decode txDirectCurrentLocation", err)
				}
				ie.txDirectCurrentLocation = &tmp_int_txDirectCurrentLocation
			}
		}
	}
	return nil
}
