package blockreader

import (
	"bufio"
	"bytes"
)

// closers are the pair matches for opening tags of blocks
var closers = map[byte]byte{
	'{': '}',
	'[': ']',
	'"': '"',
}

// NextBlock finds the next JSON block opening with open and finds the correct ending.
// You should set bufferSize to a number slightly larger than the expected result size for efficiency of allocation,
// but it grows automatically if necessary.
// It returns any error encountered reading input reader, including io.EOF if reached, because it should not reach EOF.
// It stops reading from reader as soon as it finds the right closer.
// It is linear in time, using recursion but only one buffer writer.
func NextBlock(reader *bufio.Reader, open byte, bufferSize int) ([]byte, error) {
	slice := make([]byte, 0, bufferSize)
	writer := bytes.NewBuffer(slice)

	// read and discard up to delimiter open
	_, err := reader.ReadBytes(open)
	if err != nil {
		return nil, err
	}

	_ = writer.WriteByte(open)
	toBlockCloser(reader, closers[open], writer)
	return writer.Bytes(), nil
}

// toBlockCloser expects to be inside the block that it must find the exit. This is for more efficient recursion.
func toBlockCloser(reader *bufio.Reader, closer byte, writer *bytes.Buffer) error {
	for {
		b, err := copyByte(reader, writer)
		if err != nil {
			return err
		}

		switch b {
		case '\\':

			// if we are inside a string and must scape the next character
			// copy next already to not process rules with it.

			if closer == '"' {
				_, err = copyByte(reader, writer)
				if err != nil {
					return err
				}
			}

		case closer:

			// if we found a close not escaped stop and return
			return nil


		case '{', '[', '"':

			// if

			err = toBlockCloser(reader, closers[b], writer)
			if err != nil {
				return err
			}
		}

	}
}

func copyByte(reader *bufio.Reader, writer *bytes.Buffer) (byte, error) {
	b, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	writer.WriteByte(b)
	return b, nil
}
