package file

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

const delimeter = '\n'

type LineStorage struct {
	fileName       string
	base64Encoding *base64.Encoding
}

func NewLineStorage(fileName string) *LineStorage {
	return &LineStorage{
		fileName:       fileName,
		base64Encoding: base64.StdEncoding,
	}
}

func (s *LineStorage) Insert(line string) (retErr error) {
	f, err := os.OpenFile(s.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	defer func() {
		retErr = errors.Join(retErr, f.Close())
	}()

	var (
		base64line = s.base64Encoding.EncodeToString([]byte(line))
		lineRecord = base64line + string(delimeter)
	)

	if _, err := f.Write([]byte(lineRecord)); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (s *LineStorage) Read() (ls []string, retErr error) {
	f, err := os.OpenFile(s.fileName, os.O_RDONLY, 0)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("open file: %w", err)
	}

	if os.IsNotExist(err) {
		return nil, nil
	}

	defer func() {
		retErr = errors.Join(retErr, f.Close())
	}()

	b := bufio.NewReader(f)

	var lines []string

	for {
		lineRecord, err := b.ReadString(delimeter)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("read file: %w", err)
		}

		if errors.Is(err, io.EOF) {
			break
		}

		line, err := s.base64Encoding.DecodeString(string(lineRecord))
		if err != nil {
			return nil, fmt.Errorf("decode line: %w", err)
		}

		lines = append(lines, string(line))
	}

	return lines, err
}
