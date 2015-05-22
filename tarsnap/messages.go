package tarsnap

import (
	//"encoding/binary"
	"reflect"
	"strconv"
)

/**
 * Packet format:
 * position length
 * 0        1      packet type (encrypted)
 * 1        4      data length, big-endian (encrypted)
 * 5        32     SHA256(data) (encrypted)
 * 37       32     HMAC(ciphertext bytes 0--36) (not encrypted)
 * 69       N      packet data (encrypted)
 */

type Packet struct {
	// packet type (encrypted)
	Type byte `length:"1"`

	// data length, big-endian (encrypted)
	Length int `length:"4"`

	// SHA256(data) (encrypted)
	SHA256 []byte `length:"32"`

	// HMAC(ciphertext bytes 0--36) (not encrypted)
	HMAC []byte `length:"32"`

	// packet data (encrypted)
	Data []byte `length:"-1"`
}

func Marshal(msg interface{}) []byte {

	return nil
}

func UnMarshal(data []byte, out interface{}) (err error) {

	v := reflect.ValueOf(out).Elem()

	if len(data) == 0 {

		return
	}

	for i := 0; i < v.NumField(); i++ {

		field := v.Field(i)
		t := field.Type()
		typeField := v.Type().Field(i)
		tag := typeField.Tag

		switch t.Kind() {

		default:

			return

		case reflect.Int:

			if len(data) == 0 {

				return
			}

			field.SetInt(int64(data[0]))
			data = data[1:]

		case reflect.Uint8:

			if len(data) == 0 {

				return
			}

			field.SetUint(uint64(data[0]))
			data = data[1:]

		case reflect.Slice:

			if len(data) == 0 {

				return
			}

			length, err := strconv.Atoi(tag.Get("length"))
			if err != nil {

				return err
			}

			if length > len(data) {

				continue
			}

			slice, out, err := parseSlice(data, length)
			if err != nil {

				return err
			}

			field.SetBytes(slice)
			data = out
		}
	}

	if len(data) != 0 {

		return
	}

	return
}

func parseSlice(data []byte, length int) (slice []byte, out []byte, err error) {

	if length == -1 {

		length = len(data)
	}

	var b []byte
	for i, v := range data {

		if i >= length {

			break
		}
		b = append(b, v)
	}
	out = data[length:]

	return
}
