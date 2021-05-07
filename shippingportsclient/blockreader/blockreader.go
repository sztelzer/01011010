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
func NextBlock(reader *bufio.Reader, open byte, bufferSize int, counter int) ([]byte, int, error) {
	slice := make([]byte, 0, bufferSize)
	writer := bytes.NewBuffer(slice)

	// read and discard up to delimiter open
	discard, err := reader.ReadBytes(open)
	if err != nil {
		return nil, counter, err
	}
	counter += len(discard)

	_ = writer.WriteByte(open)
	counter, err = toBlockCloser(reader, closers[open], writer, counter)
	return writer.Bytes(), counter, err
}

// toBlockCloser expects to be inside the block that it must find the exit. This is for more efficient recursion.
func toBlockCloser(reader *bufio.Reader, closer byte, writer *bytes.Buffer, counter int) (int, error) {
	var b byte
	var err error

	for {
		b, counter, err = copyByte(reader, writer, counter)
		if err != nil {
			return counter, err
		}

		switch b {
		case '\\':
			// if we are inside a string and must scape the next character
			// copy next already to not process rules with it.
			if closer == '"' {
				_, counter, err = copyByte(reader, writer, counter)
				if err != nil {
					return counter, err
				}
			}
			
		case closer:
			// if we found a close not escaped stop and return
			return counter, nil
			
		case '{', '[', '"':
			// if we find something opening, find exit again
			counter, err = toBlockCloser(reader, closers[b], writer, counter)
			if err != nil {
				return counter, err
			}
		}
	}
}

func copyByte(reader *bufio.Reader, writer *bytes.Buffer, counter int) (byte, int, error) {
	b, err := reader.ReadByte()
	if err != nil {
		return 0, counter, err
	}

	counter += 1

	writer.WriteByte(b)
	return b, counter, nil
}
