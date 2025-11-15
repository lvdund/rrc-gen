package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDAP_Parameters struct {
	as_ReflectiveQoS  *SDAP_Parameters_as_ReflectiveQoS  `optional`
	sdap_QOS_IAB_r16  *SDAP_Parameters_sdap_QOS_IAB_r16  `optional,ext-1`
	sdapHeaderIAB_r16 *SDAP_Parameters_sdapHeaderIAB_r16 `optional,ext-1`
}

func (ie *SDAP_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sdap_QOS_IAB_r16 != nil || ie.sdapHeaderIAB_r16 != nil
	preambleBits := []bool{hasExtensions, ie.as_ReflectiveQoS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.as_ReflectiveQoS != nil {
		if err = ie.as_ReflectiveQoS.Encode(w); err != nil {
			return utils.WrapError("Encode as_ReflectiveQoS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sdap_QOS_IAB_r16 != nil || ie.sdapHeaderIAB_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SDAP_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sdap_QOS_IAB_r16 != nil, ie.sdapHeaderIAB_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sdap_QOS_IAB_r16 optional
			if ie.sdap_QOS_IAB_r16 != nil {
				if err = ie.sdap_QOS_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sdap_QOS_IAB_r16", err)
				}
			}
			// encode sdapHeaderIAB_r16 optional
			if ie.sdapHeaderIAB_r16 != nil {
				if err = ie.sdapHeaderIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sdapHeaderIAB_r16", err)
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

func (ie *SDAP_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var as_ReflectiveQoSPresent bool
	if as_ReflectiveQoSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if as_ReflectiveQoSPresent {
		ie.as_ReflectiveQoS = new(SDAP_Parameters_as_ReflectiveQoS)
		if err = ie.as_ReflectiveQoS.Decode(r); err != nil {
			return utils.WrapError("Decode as_ReflectiveQoS", err)
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

			sdap_QOS_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sdapHeaderIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sdap_QOS_IAB_r16 optional
			if sdap_QOS_IAB_r16Present {
				ie.sdap_QOS_IAB_r16 = new(SDAP_Parameters_sdap_QOS_IAB_r16)
				if err = ie.sdap_QOS_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sdap_QOS_IAB_r16", err)
				}
			}
			// decode sdapHeaderIAB_r16 optional
			if sdapHeaderIAB_r16Present {
				ie.sdapHeaderIAB_r16 = new(SDAP_Parameters_sdapHeaderIAB_r16)
				if err = ie.sdapHeaderIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sdapHeaderIAB_r16", err)
				}
			}
		}
	}
	return nil
}
